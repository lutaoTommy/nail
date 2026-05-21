# Apple 登录回调页面 - 服务器端实现

## 📋 说明

由于UniApp的WebView有CSP限制，无法直接在WebView中完成Apple授权。我们采用以下方案：

1. **App端**：使用系统浏览器打开Apple授权页面
2. **用户授权**：在系统浏览器中完成Apple登录
3. **回调处理**：Apple POST请求到你的服务器回调接口
4. **返回App**：通过Deep Link将结果传回App

---

## 🔧 方案A：Node.js 后端实现（推荐）

### 1. 创建回调接口

```javascript
// server.js - Node.js + Express 示例
const express = require('express');
const jwt = require('jsonwebtoken');
const jwksClient = require('jwks-rsa');
const app = express();

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

// Apple JWKS 客户端
const client = jwksClient({
  jwksUri: 'https://appleid.apple.com/auth/keys'
});

function getKey(header, callback) {
  client.getSigningKey(header.kid, function(err, key) {
    const signingKey = key.publicKey || key.rsaPublicKey;
    callback(null, signingKey);
  });
}

// POST /auth/apple/callback
app.post('/auth/apple/callback', async (req, res) => {
  try {
    const { code, id_token, state, error } = req.body;
    
    console.log('[Apple Callback] Received:', { code: !!code, id_token: !!id_token, state, error });
    
    // 检查是否有错误
    if (error) {
      console.error('[Apple Callback] Error:', error);
      return res.send(`
        <html>
          <body>
            <script>
              alert('授权失败: ${error}');
              window.close();
            </script>
          </body>
        </html>
      `);
    }
    
    // 验证 id_token（可选，增强安全性）
    let appleUserInfo = {};
    if (id_token) {
      try {
        const decoded = await new Promise((resolve, reject) => {
          jwt.verify(id_token, getKey, {
            algorithms: ['RS256'],
            issuer: 'https://appleid.apple.com',
            audience: 'com.tintashift.service' // 你的 Service ID
          }, (err, decoded) => {
            if (err) reject(err);
            else resolve(decoded);
          });
        });
        
        appleUserInfo = {
          sub: decoded.sub, // Apple用户ID
          email: decoded.email,
          email_verified: decoded.email_verified
        };
        
        console.log('[Apple Callback] User info:', appleUserInfo);
      } catch (verifyErr) {
        console.error('[Apple Callback] Token verification failed:', verifyErr);
      }
    }
    
    // 构建 Deep Link URL
    // 格式: tintashift://apple-login?code=xxx&id_token=xxx&state=xxx
    const deepLinkParams = new URLSearchParams();
    if (code) deepLinkParams.set('code', code);
    if (id_token) deepLinkParams.set('id_token', id_token);
    if (state) deepLinkParams.set('state', state);
    
    const deepLinkUrl = `tintashift://apple-login?${deepLinkParams.toString()}`;
    
    console.log('[Apple Callback] Redirecting to:', deepLinkUrl);
    
    // 返回HTML页面，自动打开Deep Link
    res.send(`
      <!DOCTYPE html>
      <html>
      <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>授权成功</title>
        <style>
          body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            margin: 0;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
          }
          .container {
            text-align: center;
            padding: 40px;
          }
          h1 { font-size: 24px; margin-bottom: 10px; }
          p { font-size: 14px; opacity: 0.8; }
          .spinner {
            border: 3px solid rgba(255,255,255,0.3);
            border-top: 3px solid white;
            border-radius: 50%;
            width: 40px;
            height: 40px;
            animation: spin 1s linear infinite;
            margin: 20px auto;
          }
          @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
          }
        </style>
      </head>
      <body>
        <div class="container">
          <div class="spinner"></div>
          <h1>授权成功</h1>
          <p>正在返回App...</p>
        </div>
        <script>
          // 尝试打开 Deep Link
          window.location.href = '${deepLinkUrl}';
          
          // 如果5秒后还在页面，提示用户手动打开App
          setTimeout(() => {
            document.querySelector('h1').textContent = '请手动打开App';
            document.querySelector('p').textContent = '授权已完成，请返回App查看';
          }, 5000);
        </script>
      </body>
      </html>
    `);
    
  } catch (err) {
    console.error('[Apple Callback] Server error:', err);
    res.status(500).send('Server error');
  }
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`Apple callback server running on port ${PORT}`);
});
```

### 2. 安装依赖

```bash
npm init -y
npm install express jsonwebtoken jwks-rsa
```

### 3. 运行服务器

```bash
node server.js
```

---

## 🔧 方案B：Nginx + 静态HTML（简化版）

如果你没有Node.js环境，可以使用Nginx配置一个简单的回调页面：

### 1. 创建回调HTML页面

```html
<!-- /var/www/html/auth/apple/callback.html -->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Apple Login Callback</title>
  <style>
    body {
      font-family: -apple-system, BlinkMacSystemFont, sans-serif;
      display: flex;
      justify-content: center;
      align-items: center;
      min-height: 100vh;
      margin: 0;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: white;
    }
    .container { text-align: center; padding: 40px; }
    h1 { font-size: 24px; margin-bottom: 10px; }
    p { font-size: 14px; opacity: 0.8; }
  </style>
</head>
<body>
  <div class="container">
    <h1>授权处理中...</h1>
    <p>请稍候</p>
  </div>
  
  <script>
    // 获取URL参数
    const params = new URLSearchParams(window.location.search);
    const code = params.get('code');
    const id_token = params.get('id_token');
    const state = params.get('state');
    const error = params.get('error');
    
    if (error) {
      document.querySelector('h1').textContent = '授权失败';
      document.querySelector('p').textContent = error;
    } else {
      // 构建 Deep Link
      const deepLinkParams = new URLSearchParams();
      if (code) deepLinkParams.set('code', code);
      if (id_token) deepLinkParams.set('id_token', id_token);
      if (state) deepLinkParams.set('state', state);
      
      const deepLinkUrl = 'tintashift://apple-login?' + deepLinkParams.toString();
      
      // 打开 Deep Link
      window.location.href = deepLinkUrl;
      
      setTimeout(() => {
        document.querySelector('h1').textContent = '请手动打开App';
        document.querySelector('p').textContent = '授权已完成';
      }, 3000);
    }
  </script>
</body>
</html>
```

### 2. Nginx配置

```nginx
server {
    listen 443 ssl;
    server_name www.tintashift.top;
    
    # SSL证书配置
    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;
    
    # Apple回调接口
    location /auth/apple/callback {
        # Apple会以POST方式发送数据
        # 需要将其转换为GET重定向
        
        # 方法1：使用Lua脚本（需要openresty）
        # 方法2：直接返回HTML页面，让前端JavaScript处理
        
        default_type text/html;
        return 200 '
            <!DOCTYPE html>
            <html>
            <head><meta charset="UTF-8"><title>Processing...</title></head>
            <body>
                <script>
                    // 这个方案有限，因为POST数据无法直接获取
                    // 建议使用方案A（Node.js）或云函数
                    window.location.href = "tintashift://apple-login?status=callback_received";
                </script>
            </body>
            </html>
        ';
    }
}
```

---

## 🔧 方案C：使用云函数（最简单）

### 腾讯云云函数示例

```javascript
// index.js
exports.main = async (event, context) => {
  const { code, id_token, state, error } = event.body;
  
  if (error) {
    return {
      statusCode: 200,
      headers: { 'Content-Type': 'text/html' },
      body: `<html><body><script>alert('授权失败');window.close();</script></body></html>`
    };
  }
  
  // 构建 Deep Link
  const params = [];
  if (code) params.push(`code=${encodeURIComponent(code)}`);
  if (id_token) params.push(`id_token=${encodeURIComponent(id_token)}`);
  if (state) params.push(`state=${encodeURIComponent(state)}`);
  
  const deepLinkUrl = `tintashift://apple-login?${params.join('&')}`;
  
  return {
    statusCode: 200,
    headers: { 'Content-Type': 'text/html' },
    body: `
      <!DOCTYPE html>
      <html>
      <head>
        <meta charset="UTF-8">
        <meta http-equiv="refresh" content="0;url=${deepLinkUrl}">
      </head>
      <body>
        <script>window.location.href = '${deepLinkUrl}';</script>
        <p>Redirecting to App...</p>
      </body>
      </html>
    `
  };
};
```

---

## 📱 App端接收Deep Link

在 `App.vue` 中添加Deep Link监听：

```vue
<script setup lang="ts">
import { onLaunch, onShow } from '@dcloudio/uni-app'

onLaunch(() => {
  console.log('App launched')
})

onShow(() => {
  // #ifdef APP-PLUS
  // 监听Deep Link
  plus.globalEvent.addEventListener('newintent', (e) => {
    const url = plus.runtime.arguments
    console.log('[DeepLink] Received:', url)
    
    if (url && url.startsWith('tintashift://apple-login')) {
      handleAppleDeepLink(url)
    }
  })
  // #endif
})

function handleAppleDeepLink(url: string) {
  try {
    // 解析URL参数
    const queryString = url.split('?')[1]
    const params = new URLSearchParams(queryString)
    
    const code = params.get('code')
    const id_token = params.get('id_token')
    const state = params.get('state')
    
    console.log('[DeepLink] Parsed:', { code: !!code, id_token: !!id_token, state })
    
    // 保存授权结果到本地存储
    uni.setStorageSync('apple_auth_result', {
      code,
      id_token,
      state
    })
    
  } catch (err) {
    console.error('[DeepLink] Parse error:', err)
  }
}
</script>
```

---

## ✅ 完整流程测试

1. **部署后端**：选择上述任一方案部署回调接口
2. **配置Apple Developer**：
   - Return URLs: `https://www.tintashift.top/auth/apple/callback`
3. **更新App配置**：
   - `apple-login.vue` 中的 `redirectUri`
4. **重新编译App**
5. **测试流程**：
   - 点击Apple登录
   - 系统浏览器打开
   - 完成授权
   - 自动跳转回App
   - 登录成功

---

## 💡 注意事项

1. **HTTPS必需**：回调URL必须使用HTTPS
2. **URL Scheme**：确保已在manifest.json中配置 `tintashift`
3. **Service ID匹配**：clientId必须与Apple Developer配置一致
4. **测试环境**：建议先在开发环境测试完整流程
