package language

var languageMapZhcn map[string]string
/*初始化*/
func initLanguageZhcn() {
	languageMapZhcn = make(map[string]string)
	/*user*/
	languageMapZhcn["E_NO_IV"] = "请填写初始向量"
	languageMapZhcn["E_NO_CODE"] = "请填写授权码"
	languageMapZhcn["E_NO_DATA"] = "请填写数据"
	languageMapZhcn["E_NO_PHONE"] = "请输入手机号"
	languageMapZhcn["E_NO_PWD"] = "请输入密码"
	languageMapZhcn["E_INVALID_USER"] = "无此用户"
	languageMapZhcn["E_INVALID_PWD"] = "密码错误"
	languageMapZhcn["E_NO_NAME"] = "请输入名称"
	languageMapZhcn["E_USER_EXIST"] = "账号已存在"
	languageMapZhcn["E_NO_TOKEN"] = "请先登录"
	languageMapZhcn["E_NO_AUTH"] = "没有权限"
	languageMapZhcn["E_NO_AVATAR"] = "无此头像"
	languageMapZhcn["E_NO_EMAIL"] = "请填写邮箱"
	languageMapZhcn["E_NO_CERT"] = "请输入验证码"
	languageMapZhcn["E_INVALID_EMAIL"] = "邮箱格式错误"
	languageMapZhcn["E_DATA_TOO_SHORT"] = "加密数据过短"
	languageMapZhcn["E_INVALID_IV"] = "初始向量无效"
	languageMapZhcn["E_INVALID_SESSION_KEY"] = "会话密钥无效"
	languageMapZhcn["E_DECRYPT_FAIL"] = "解密失败，请确保授权码与手机号为同一次授权"
	languageMapZhcn["E_INVALID_PHONE"] = "手机号格式错误"
	languageMapZhcn["VERIFICATION_CODE"] = "验证码"
	languageMapZhcn["DO_NOT_TELL"] = "请勿告诉他人"
	languageMapZhcn["E_CERT_ERR"] = "验证码错误"
	languageMapZhcn["E_CERT_FIRST"] = "请先获取验证码"
	languageMapZhcn["E_VERIFICATION_QUICKLY"] = "获取验证码太频繁"
	languageMapZhcn["E_ACCOUNT_LOCKED"] = "尝试次数过多，请稍后再试"
	languageMapZhcn["E_VERIFICATION_LIMIT"] = "请求验证码过于频繁，请稍后再试"
	/*common*/
	languageMapZhcn["E_TOO_LONG"] = "字符长度超过限制"
	languageMapZhcn["E_INVALID_PARAM"] = "参数错误"
	languageMapZhcn["E_FILE_NOT_FOUND"] = "文件不存在"
	/*color*/
	languageMapZhcn["E_INVALID_COLOR"] = "无此颜色"
	languageMapZhcn["E_INVALID_FILE"] = "请上传文件"
	languageMapZhcn["E_INVALID_PICTURE"] = "请上传图片(JPG/PNG)"
	/*device*/
	languageMapZhcn["E_NO_ID"] = "请填写ID"
	languageMapZhcn["E_NO_MAC"] = "请填写MAC地址"
	languageMapZhcn["E_NO_INFO"] = "无此记录"
	languageMapZhcn["E_NO_DEVICE"] = "无此设备"
	languageMapZhcn["E_INVALID_MAC"] = "请填写正确格式的MAC地址"
	/*suggest*/
	languageMapZhcn["E_NO_CONTENT"] = "请填写内容"
	languageMapZhcn["E_NO_SUGGEST"] = "无此建议"
	/*word*/
	languageMapZhcn["E_NO_WORD"] = "无此敏感词"
	languageMapZhcn["E_WORD_EXIST"] = "敏感词已存在"
	/*follow*/
	languageMapZhcn["E_FOLLOWED"] = "已关注"
	languageMapZhcn["E_UNFOLLOWED"] = "已取关"
	languageMapZhcn["E_NOT_FOLLOWING"] = "未关注"
	languageMapZhcn["E_NO_SELF_FOLLOW"] = "不能关注自己"
	/*circle*/
	languageMapZhcn["E_NO_CIRCLE_POST"] = "无此动态"
	languageMapZhcn["E_SENSITIVEWORD"] = "您的发言中包含敏感信息,为维护友好的交流氛围,请调整相关表述后再发布"
	/*comment*/
	languageMapZhcn["E_NO_COMMENT"] = "无此评论"
	/*like*/
	languageMapZhcn["E_NO_LIKE"] = "无此点赞"
	/*collect*/
	languageMapZhcn["E_NO_COLLECT"] = "未收藏过"
	/*ota*/
	languageMapZhcn["E_NO_VERSION"] = "请填写版本"
}