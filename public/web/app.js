/* TINTAShift Circle Review Admin */
(() => {
  const TOKEN_KEY = 'circle_admin_token';
  const USER_KEY = 'circle_admin_user';
  const PAGE_SIZE = 10;

  const state = {
    view: 'review',
    page: 1,
    total: 0,
    loading: false,
    topics: [],
    topicLoading: false,
    suggestPage: 1,
    suggestTotal: 0,
    suggestLoading: false,
    wordPage: 1,
    wordTotal: 0,
    wordLoading: false,
    wordSearchTimer: null,
    pendingDelete: null, // { type: 'post' | 'topic' | 'suggest' | 'word', id, name }
  };

  const $ = (id) => document.getElementById(id);

  // ===== Token =====
  function getToken() {
    return localStorage.getItem(TOKEN_KEY);
  }
  function setToken(token) {
    localStorage.setItem(TOKEN_KEY, token);
  }
  function setUser(userId) {
    localStorage.setItem(USER_KEY, userId || '');
  }
  function clearAuth() {
    localStorage.removeItem(TOKEN_KEY);
    localStorage.removeItem(USER_KEY);
  }

  // ===== Toast =====
  let toastTimer = null;
  function showToast(msg, duration = 2200) {
    const toast = $('toast');
    toast.textContent = msg;
    toast.style.display = 'block';
    // 触发重排以播放过渡动画
    void toast.offsetWidth;
    toast.classList.add('show');
    clearTimeout(toastTimer);
    toastTimer = setTimeout(() => {
      toast.classList.remove('show');
      setTimeout(() => {
        toast.style.display = 'none';
      }, 260);
    }, duration);
  }

  // ===== API =====
  async function request(path, options = {}) {
    const headers = { ...(options.headers || {}) };
    const token = getToken();
    if (token) headers['token'] = token;
    if (options.body) headers['Content-Type'] = 'application/json';

    let res;
    try {
      res = await fetch(path, { ...options, headers });
    } catch (e) {
      throw new Error('网络异常，请稍后重试');
    }

    let data;
    try {
      data = await res.json();
    } catch (e) {
      throw new Error('服务响应异常');
    }

    if (data.result_code === 401) {
      clearAuth();
      showLogin();
      throw new Error(data.result_msg || '登录已失效，请重新登录');
    }
    if (data.result_code === 403) {
      throw new Error(data.result_msg || '没有权限，仅管理员可操作');
    }
    if (data.result_code !== 200) {
      throw new Error(data.result_msg || '请求失败');
    }
    return data;
  }

  // ===== Views =====
  function showLogin() {
    $('loginPage').style.display = 'flex';
    $('app').style.display = 'none';
  }

  function showApp() {
    $('loginPage').style.display = 'none';
    $('app').style.display = 'flex';
  }

  // ===== View Tabs =====
  function switchView(view) {
    state.view = view;
    document.querySelectorAll('.view-tab').forEach((tab) => {
      tab.classList.toggle('active', tab.dataset.view === view);
    });
    $('reviewView').style.display = view === 'review' ? 'block' : 'none';
    $('topicView').style.display = view === 'topic' ? 'block' : 'none';
    $('suggestView').style.display = view === 'suggest' ? 'block' : 'none';
    $('wordView').style.display = view === 'word' ? 'block' : 'none';
    if (view === 'topic') {
      loadTopics();
    } else if (view === 'suggest') {
      loadSuggestions(state.suggestPage);
    } else if (view === 'word') {
      loadWords(state.wordPage);
    }
  }

  document.querySelectorAll('.view-tab').forEach((tab) => {
    tab.addEventListener('click', () => switchView(tab.dataset.view));
  });

  // ===== Login =====
  const loginForm = $('loginForm');
  const loginError = $('loginError');

  loginForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    loginError.style.display = 'none';

    const phone = $('loginAccount').value.trim();
    const passwd = $('loginPassword').value;
    if (!phone || !passwd) {
      loginError.textContent = '请输入账号和密码';
      loginError.style.display = 'block';
      return;
    }

    const btn = loginForm.querySelector('button[type="submit"]');
    btn.disabled = true;
    btn.textContent = '登录中...';

    try {
      const data = await request('/user/login', {
        method: 'POST',
        body: JSON.stringify({ phone, passwd }),
      });
      if (data.user_id !== 'admin') {
        clearAuth();
        throw new Error('该账号不是管理员，无法进入审核后台');
      }
      setToken(data.token);
      setUser(data.user_id);
      $('loginPassword').value = '';
      showApp();
      switchView('review');
      showToast('登录成功');
      loadList(1);
    } catch (err) {
      loginError.textContent = err.message;
      loginError.style.display = 'block';
    } finally {
      btn.disabled = false;
      btn.textContent = '登录';
    }
  });

  // ===== Logout =====
  $('logoutBtn').addEventListener('click', () => {
    clearAuth();
    showLogin();
    showToast('已退出登录');
  });

  // ===== Filters =====
  $('filterForm').addEventListener('submit', (e) => {
    e.preventDefault();
    loadList(1);
  });

  $('resetBtn').addEventListener('click', () => {
    $('filterReview').value = '';
    $('filterKeyword').value = '';
    $('filterTopic').value = '';
    loadList(1);
  });

  $('refreshBtn').addEventListener('click', () => loadList(state.page));

  // ===== Topic Handlers =====
  $('topicForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const input = $('topicName');
    const name = input.value.trim();
    if (!name) {
      showToast('请输入话题名称');
      input.focus();
      return;
    }
    if (name.length > 50) {
      showToast('话题名称不能超过 50 个字符');
      return;
    }

    const btn = e.target.querySelector('button[type="submit"]');
    const originalText = btn.textContent;
    btn.disabled = true;
    btn.textContent = '添加中...';
    try {
      await request('/topic/add?name=' + encodeURIComponent(name), { method: 'POST' });
      input.value = '';
      $('topicSearch').value = '';
      showToast('话题添加成功');
      await loadTopics();
    } catch (err) {
      showToast(err.message);
    } finally {
      btn.disabled = false;
      btn.textContent = originalText;
    }
  });

  $('topicSearch').addEventListener('input', renderTopics);
  $('topicRefreshBtn').addEventListener('click', loadTopics);

  async function loadTopics() {
    if (state.topicLoading) return;
    state.topicLoading = true;
    $('topicLoading').style.display = 'block';
    $('topicEmpty').style.display = 'none';
    $('topicList').innerHTML = '';
    $('topicRefreshBtn').disabled = true;

    try {
      const data = await request('/topic/list');
      state.topics = Array.isArray(data.data) ? data.data : [];
      renderTopics();
    } catch (err) {
      state.topics = [];
      $('topicCount').textContent = '0 个话题';
      $('topicEmpty').textContent = err.message;
      $('topicEmpty').style.display = 'block';
    } finally {
      state.topicLoading = false;
      $('topicLoading').style.display = 'none';
      $('topicRefreshBtn').disabled = false;
    }
  }

  function renderTopics() {
    const list = $('topicList');
    list.innerHTML = '';
    const keyword = $('topicSearch').value.trim().toLowerCase();
    const topics = state.topics.filter((topic) => {
      if (!keyword) return true;
      return String(topic.name || '').toLowerCase().includes(keyword);
    });

    $('topicCount').textContent = `${state.topics.length} 个话题`;

    if (topics.length === 0) {
      $('topicEmpty').textContent = state.topics.length === 0 ? '暂无话题' : '没有匹配的话题';
      $('topicEmpty').style.display = 'block';
      return;
    }
    $('topicEmpty').style.display = 'none';

    topics.forEach((topic) => {
      const row = el('div', 'topic-item');

      const main = el('div', 'topic-main');
      main.appendChild(el('div', 'topic-name', topic.name || '-'));
      main.appendChild(
        el('div', 'topic-meta', `ID: ${topic.id || '-'} · 创建于 ${topic.create_time || '-'}`)
      );
      row.appendChild(main);

      const deleteBtn = el('button', 'btn-card btn-delete', '删除');
      deleteBtn.type = 'button';
      deleteBtn.addEventListener('click', () => {
        openConfirm({ type: 'topic', id: topic.id, name: topic.name });
      });
      row.appendChild(deleteBtn);

      list.appendChild(row);
    });
  }

  // ===== Suggest Handlers =====
  $('suggestRefreshBtn').addEventListener('click', () => loadSuggestions(state.suggestPage));

  async function loadSuggestions(page) {
    if (state.suggestLoading) return;
    state.suggestLoading = true;
    $('suggestLoading').style.display = 'block';
    $('suggestEmpty').style.display = 'none';
    $('suggestList').innerHTML = '';
    $('suggestPagination').innerHTML = '';
    $('suggestRefreshBtn').disabled = true;

    try {
      const params = new URLSearchParams();
      params.set('page', String(page));
      params.set('limit', String(PAGE_SIZE));
      const data = await request('/suggest/list?' + params.toString());
      const items = Array.isArray(data.data) ? data.data : [];
      state.suggestTotal = Number(data.total) || 0;

      if (items.length === 0 && page > 1 && state.suggestTotal > 0) {
        state.suggestLoading = false;
        $('suggestLoading').style.display = 'none';
        $('suggestRefreshBtn').disabled = false;
        return loadSuggestions(page - 1);
      }

      state.suggestPage = page;
      renderSuggestions(items);
      renderSuggestPagination();
      $('suggestCount').textContent = `${state.suggestTotal} 条建议`;
    } catch (err) {
      $('suggestEmpty').textContent = err.message;
      $('suggestEmpty').style.display = 'block';
    } finally {
      state.suggestLoading = false;
      $('suggestLoading').style.display = 'none';
      $('suggestRefreshBtn').disabled = false;
    }
  }

  function renderSuggestions(items) {
    const list = $('suggestList');
    list.innerHTML = '';
    if (items.length === 0) {
      $('suggestEmpty').textContent = '暂无建议';
      $('suggestEmpty').style.display = 'block';
      return;
    }
    $('suggestEmpty').style.display = 'none';

    items.forEach((item) => {
      const row = el('div', 'suggest-item');

      const head = el('div', 'suggest-head');
      head.appendChild(renderAvatar(item));

      const user = el('div', 'suggest-user');
      user.appendChild(el('div', 'topic-name', item.nickname || '匿名用户'));
      user.appendChild(
        el('div', 'topic-meta', [item.phone, item.time].filter(Boolean).join(' · ') || '-')
      );
      head.appendChild(user);

      const actions = el('div', 'actions');
      const detailBtn = el('button', 'btn-card', '查看');
      detailBtn.type = 'button';
      detailBtn.addEventListener('click', () => openSuggestDetail(item.id));
      actions.appendChild(detailBtn);

      const deleteBtn = el('button', 'btn-card btn-delete', '删除');
      deleteBtn.type = 'button';
      deleteBtn.addEventListener('click', () => {
        openConfirm({ type: 'suggest', id: item.id });
      });
      actions.appendChild(deleteBtn);
      head.appendChild(actions);

      row.appendChild(head);
      if (item.content) {
        row.appendChild(el('p', 'suggest-content', item.content));
      }
      list.appendChild(row);
    });
  }

  function renderSuggestPagination() {
    const wrap = $('suggestPagination');
    wrap.innerHTML = '';
    const totalPages = Math.max(1, Math.ceil(state.suggestTotal / PAGE_SIZE));

    const prev = el('button', 'page-btn', '上一页');
    prev.type = 'button';
    prev.disabled = state.suggestPage <= 1;
    prev.addEventListener('click', () => loadSuggestions(state.suggestPage - 1));

    const info = el('span', 'page-info', `第 ${state.suggestPage} / ${totalPages} 页`);

    const next = el('button', 'page-btn', '下一页');
    next.type = 'button';
    next.disabled = state.suggestPage >= totalPages;
    next.addEventListener('click', () => loadSuggestions(state.suggestPage + 1));

    wrap.appendChild(prev);
    wrap.appendChild(info);
    wrap.appendChild(next);
  }

  // ===== Suggest Detail =====
  const suggestDetailModal = $('suggestDetailModal');
  const suggestDetailBody = $('suggestDetailBody');

  async function openSuggestDetail(id) {
    suggestDetailBody.innerHTML = '<div class="state-box">加载中...</div>';
    suggestDetailModal.style.display = 'flex';
    try {
      const data = await request('/suggest/detail?id=' + encodeURIComponent(id));
      renderSuggestDetail(data.data || {});
    } catch (err) {
      suggestDetailBody.innerHTML = '';
      suggestDetailBody.appendChild(el('div', 'state-box', err.message));
    }
  }

  function renderSuggestDetail(detail) {
    suggestDetailBody.innerHTML = '';

    const user = el('div', 'suggest-detail-user');
    user.appendChild(renderAvatar(detail));
    const userInfo = el('div', 'suggest-user');
    userInfo.appendChild(el('div', 'topic-name', detail.nickname || '匿名用户'));
    userInfo.appendChild(
      el('div', 'topic-meta', [detail.phone, detail.time].filter(Boolean).join(' · ') || '-')
    );
    user.appendChild(userInfo);
    suggestDetailBody.appendChild(user);

    suggestDetailBody.appendChild(el('div', 'suggest-detail-content', detail.content || ''));

    const attachs = Array.isArray(detail.attachs) ? detail.attachs : [];
    if (attachs.length > 0) {
      const grid = el('div', 'media-grid suggest-detail-images');
      attachs.forEach((attach) => {
        const link = el('a', 'media-item');
        link.href = attach.url;
        link.target = '_blank';
        link.rel = 'noopener noreferrer';
        const img = el('img');
        img.src = attach.url;
        img.alt = '建议图片';
        img.loading = 'lazy';
        link.appendChild(img);
        grid.appendChild(link);
      });
      suggestDetailBody.appendChild(grid);
    }
  }

  function closeSuggestDetail() {
    suggestDetailModal.style.display = 'none';
    suggestDetailBody.innerHTML = '';
  }

  $('suggestDetailClose').addEventListener('click', closeSuggestDetail);
  $('suggestDetailFooterClose').addEventListener('click', closeSuggestDetail);
  suggestDetailModal.addEventListener('click', (e) => {
    if (e.target === suggestDetailModal) closeSuggestDetail();
  });

  // ===== Sensitive Word Handlers =====
  $('wordForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const input = $('wordName');
    const name = input.value.trim();
    if (!name) {
      showToast('请输入敏感词');
      input.focus();
      return;
    }

    const btn = e.target.querySelector('button[type="submit"]');
    const originalText = btn.textContent;
    btn.disabled = true;
    btn.textContent = '添加中...';
    try {
      await request('/word/add', {
        method: 'POST',
        body: JSON.stringify({ name }),
      });
      input.value = '';
      $('wordSearch').value = '';
      showToast('敏感词添加成功');
      await loadWords(1);
    } catch (err) {
      showToast(err.message);
    } finally {
      btn.disabled = false;
      btn.textContent = originalText;
    }
  });

  $('wordSearch').addEventListener('input', () => {
    clearTimeout(state.wordSearchTimer);
    state.wordSearchTimer = setTimeout(() => loadWords(1), 300);
  });

  $('wordRefreshBtn').addEventListener('click', () => loadWords(state.wordPage));

  async function loadWords(page) {
    if (state.wordLoading) return;
    state.wordLoading = true;
    $('wordLoading').style.display = 'block';
    $('wordEmpty').style.display = 'none';
    $('wordList').innerHTML = '';
    $('wordPagination').innerHTML = '';
    $('wordRefreshBtn').disabled = true;

    try {
      const params = new URLSearchParams();
      params.set('page', String(page));
      params.set('limit', String(PAGE_SIZE));
      const name = $('wordSearch').value.trim();
      if (name) params.set('name', name);

      const data = await request('/word/list?' + params.toString());
      const items = Array.isArray(data.data) ? data.data : [];
      state.wordTotal = Number(data.total) || 0;

      if (items.length === 0 && page > 1 && state.wordTotal > 0) {
        state.wordLoading = false;
        $('wordLoading').style.display = 'none';
        $('wordRefreshBtn').disabled = false;
        return loadWords(page - 1);
      }

      state.wordPage = page;
      renderWords(items);
      renderWordPagination();
      $('wordCount').textContent = `${state.wordTotal} 个敏感词`;
    } catch (err) {
      $('wordEmpty').textContent = err.message;
      $('wordEmpty').style.display = 'block';
    } finally {
      state.wordLoading = false;
      $('wordLoading').style.display = 'none';
      $('wordRefreshBtn').disabled = false;
    }
  }

  function renderWords(items) {
    const list = $('wordList');
    list.innerHTML = '';
    if (items.length === 0) {
      $('wordEmpty').textContent = $('wordSearch').value.trim() ? '没有匹配的敏感词' : '暂无敏感词';
      $('wordEmpty').style.display = 'block';
      return;
    }
    $('wordEmpty').style.display = 'none';

    items.forEach((word) => {
      const row = el('div', 'topic-item');

      const main = el('div', 'topic-main');
      main.appendChild(el('div', 'topic-name', word.name || '-'));
      main.appendChild(el('div', 'topic-meta', `ID: ${word.id || '-'}`));
      row.appendChild(main);

      const deleteBtn = el('button', 'btn-card btn-delete', '删除');
      deleteBtn.type = 'button';
      deleteBtn.addEventListener('click', () => {
        openConfirm({ type: 'word', id: word.id, name: word.name });
      });
      row.appendChild(deleteBtn);

      list.appendChild(row);
    });
  }

  function renderWordPagination() {
    const wrap = $('wordPagination');
    wrap.innerHTML = '';
    const totalPages = Math.max(1, Math.ceil(state.wordTotal / PAGE_SIZE));

    const prev = el('button', 'page-btn', '上一页');
    prev.type = 'button';
    prev.disabled = state.wordPage <= 1;
    prev.addEventListener('click', () => loadWords(state.wordPage - 1));

    const info = el('span', 'page-info', `第 ${state.wordPage} / ${totalPages} 页`);

    const next = el('button', 'page-btn', '下一页');
    next.type = 'button';
    next.disabled = state.wordPage >= totalPages;
    next.addEventListener('click', () => loadWords(state.wordPage + 1));

    wrap.appendChild(prev);
    wrap.appendChild(info);
    wrap.appendChild(next);
  }

  // ===== List =====
  function buildQuery(page) {
    const params = new URLSearchParams();
    params.set('page', String(page));
    params.set('limit', String(PAGE_SIZE));
    const keyword = $('filterKeyword').value.trim();
    const topic = $('filterTopic').value.trim();
    const review = $('filterReview').value;
    if (keyword) params.set('keyword', keyword);
    if (topic) params.set('topic', topic);
    if (review !== '') params.set('review_status', review);
    return params.toString();
  }

  async function loadList(page) {
    if (state.loading) return;
    state.loading = true;
    $('loadingState').style.display = 'block';
    $('emptyState').style.display = 'none';
    $('postList').innerHTML = '';
    $('pagination').innerHTML = '';
    $('refreshBtn').disabled = true;

    try {
      const data = await request('/circle/list?' + buildQuery(page));
      const items = Array.isArray(data.data) ? data.data : [];
      state.total = Number(data.total) || 0;

      if (items.length === 0 && page > 1 && state.total > 0) {
        state.loading = false;
        $('loadingState').style.display = 'none';
        $('refreshBtn').disabled = false;
        return loadList(page - 1);
      }

      state.page = page;
      renderList(items);
      renderPagination();
      $('listCount').textContent = `${state.total} 条记录`;
    } catch (err) {
      $('emptyState').textContent = err.message;
      $('emptyState').style.display = 'block';
    } finally {
      state.loading = false;
      $('loadingState').style.display = 'none';
      $('refreshBtn').disabled = false;
    }
  }

  function renderList(items) {
    const list = $('postList');
    list.innerHTML = '';
    if (items.length === 0) {
      $('emptyState').textContent = '暂无符合条件的动态';
      $('emptyState').style.display = 'block';
      return;
    }
    $('emptyState').style.display = 'none';
    items.forEach((item) => list.appendChild(renderCard(item)));
  }

  function el(tag, className, text) {
    const node = document.createElement(tag);
    if (className) node.className = className;
    if (text !== undefined && text !== null) node.textContent = text;
    return node;
  }

  function renderCard(item) {
    const card = el('article', 'post-card');

    // 头部：头像 + 用户信息 + 审核徽标
    const head = el('div', 'post-head');
    head.appendChild(renderAvatar(item));

    const meta = el('div', 'user-meta');
    meta.appendChild(el('div', 'user-name', item.nickname || item.user_id || '匿名用户'));
    meta.appendChild(el('div', 'user-sub', `ID: ${item.user_id || '-'}`));
    head.appendChild(meta);

    const reviewed = Number(item.review_status) === 1;
    head.appendChild(
      el('span', reviewed ? 'badge badge-reviewed' : 'badge badge-unreviewed', reviewed ? '已审核' : '未审核')
    );
    card.appendChild(head);

    // 话题
    if (item.topic) {
      card.appendChild(el('span', 'post-topic', `#${item.topic}`));
    }

    // 正文
    if (item.content) {
      card.appendChild(el('p', 'post-content', item.content));
    }

    // 图片 / 视频
    if (Array.isArray(item.images) && item.images.length > 0) {
      card.appendChild(renderMedia(item.images));
    }

    // 底部：统计 + 操作
    const footer = el('div', 'post-footer');
    const stats = el('div', 'stats');
    stats.appendChild(el('span', '', `点赞 ${Number(item.like_count) || 0}`));
    stats.appendChild(el('span', '', `评论 ${Number(item.comment_count) || 0}`));
    stats.appendChild(el('span', '', `收藏 ${Number(item.collect_count) || 0}`));
    footer.appendChild(stats);

    const actions = el('div', 'actions');
    const reviewBtn = el(
      'button',
      reviewed ? 'btn-card btn-unapprove' : 'btn-card btn-approve',
      reviewed ? '取消审核' : '标记已审核'
    );
    reviewBtn.type = 'button';
    reviewBtn.addEventListener('click', () => {
      updateReview(item, reviewed ? 0 : 1, reviewBtn);
    });
    actions.appendChild(reviewBtn);

    const deleteBtn = el('button', 'btn-card btn-delete', '删除');
    deleteBtn.type = 'button';
    deleteBtn.addEventListener('click', () => {
      openConfirm({ type: 'post', id: item.id });
    });
    actions.appendChild(deleteBtn);

    footer.appendChild(actions);
    card.appendChild(footer);

    // 审核信息
    if (reviewed && (item.review_time || item.reviewer)) {
      const info = [item.reviewer ? `审核人 ${item.reviewer}` : '', item.review_time || '']
        .filter(Boolean)
        .join(' · ');
      card.appendChild(el('div', 'review-info', info));
    }

    if (item.create_time) {
      card.appendChild(el('div', 'review-info', `发布时间 ${item.create_time}`));
    }

    return card;
  }

  function renderAvatar(item) {
    if (item.avatar) {
      const img = el('img', 'avatar');
      img.src = item.avatar;
      img.alt = item.nickname || 'avatar';
      img.loading = 'lazy';
      img.addEventListener('error', () => {
        img.replaceWith(renderAvatarFallback(item));
      });
      return img;
    }
    return renderAvatarFallback(item);
  }

  function renderAvatarFallback(item) {
    const name = item.nickname || item.user_id || '?';
    return el('div', 'avatar-fallback', name.slice(0, 1).toUpperCase());
  }

  function renderMedia(images) {
    const grid = el('div', 'media-grid');
    images.forEach((media) => {
      const isVideo =
        media.media_type === 'video' ||
        media.type === 'mp4' ||
        media.type === 'mov' ||
        /\.(mp4|mov)(\?.*)?$/i.test(media.image || '');

      if (isVideo) {
        const wrap = el('div', 'media-item video');
        const video = el('video');
        video.controls = true;
        video.preload = 'metadata';
        if (media.thumbnail) video.poster = media.thumbnail;
        const source = el('source');
        source.src = media.image;
        source.type = media.type === 'mov' ? 'video/quicktime' : 'video/mp4';
        video.appendChild(source);
        wrap.appendChild(video);
        grid.appendChild(wrap);
      } else {
        const link = el('a', 'media-item');
        link.href = media.image || media.thumbnail || '#';
        link.target = '_blank';
        link.rel = 'noopener noreferrer';
        const img = el('img');
        img.src = media.thumbnail || media.image;
        img.alt = '动态图片';
        img.loading = 'lazy';
        link.appendChild(img);
        grid.appendChild(link);
      }
    });
    return grid;
  }

  async function updateReview(item, reviewStatus, btn) {
    const originalText = btn.textContent;
    btn.disabled = true;
    btn.textContent = '处理中...';
    try {
      await request('/circle/review', {
        method: 'POST',
        body: JSON.stringify({ id: item.id, review_status: reviewStatus }),
      });
      showToast(reviewStatus === 1 ? '已标记为已审核' : '已取消审核');
      await loadList(state.page);
    } catch (err) {
      showToast(err.message);
      btn.disabled = false;
      btn.textContent = originalText;
    }
  }

  // ===== Pagination =====
  function renderPagination() {
    const wrap = $('pagination');
    wrap.innerHTML = '';
    const totalPages = Math.max(1, Math.ceil(state.total / PAGE_SIZE));

    const prev = el('button', 'page-btn', '上一页');
    prev.type = 'button';
    prev.disabled = state.page <= 1;
    prev.addEventListener('click', () => loadList(state.page - 1));

    const info = el('span', 'page-info', `第 ${state.page} / ${totalPages} 页`);

    const next = el('button', 'page-btn', '下一页');
    next.type = 'button';
    next.disabled = state.page >= totalPages;
    next.addEventListener('click', () => loadList(state.page + 1));

    wrap.appendChild(prev);
    wrap.appendChild(info);
    wrap.appendChild(next);
  }

  // ===== Confirm Modal =====
  const confirmModal = $('confirmModal');
  const confirmTitle = $('confirmTitle');
  const confirmBody = $('confirmBody');
  const confirmBtn = $('confirmBtn');

  function openConfirm(pending) {
    state.pendingDelete = pending;
    confirmBtn.disabled = false;
    confirmBtn.textContent = '确认删除';

    if (pending.type === 'topic') {
      confirmTitle.textContent = '删除话题';
      confirmBody.textContent = `确认删除话题「${pending.name || pending.id}」吗？删除后该话题不再出现在客户端话题列表，但已发布动态中的话题文字不会改变。`;
    } else if (pending.type === 'suggest') {
      confirmTitle.textContent = '删除建议';
      confirmBody.textContent = '确认删除这条建议吗？删除后不可恢复，关联的图片也会一并删除。';
    } else if (pending.type === 'word') {
      confirmTitle.textContent = '删除敏感词';
      confirmBody.textContent = `确认删除敏感词「${pending.name || pending.id}」吗？删除后新发布内容将不再拦截该词。`;
    } else {
      confirmTitle.textContent = '删除动态';
      confirmBody.textContent = '确认删除这条动态吗？删除后不可恢复，关联的图片、评论、点赞和收藏也会一并删除。';
    }
    confirmModal.style.display = 'flex';
  }

  function closeConfirm() {
    state.pendingDelete = null;
    confirmModal.style.display = 'none';
  }

  $('confirmClose').addEventListener('click', closeConfirm);
  $('confirmCancelBtn').addEventListener('click', closeConfirm);
  confirmModal.addEventListener('click', (e) => {
    if (e.target === confirmModal) closeConfirm();
  });

  confirmBtn.addEventListener('click', async () => {
    const pending = state.pendingDelete;
    if (!pending) return;
    confirmBtn.disabled = true;
    confirmBtn.textContent = '处理中...';
    try {
      if (pending.type === 'topic') {
        await request('/topic/remove', {
          method: 'POST',
          body: JSON.stringify({ id: pending.id }),
        });
        showToast('已删除话题');
        closeConfirm();
        await loadTopics();
      } else if (pending.type === 'suggest') {
        await request('/suggest/remove', {
          method: 'POST',
          body: JSON.stringify({ id: pending.id }),
        });
        showToast('已删除建议');
        closeConfirm();
        await loadSuggestions(state.suggestPage);
      } else if (pending.type === 'word') {
        await request('/word/remove', {
          method: 'POST',
          body: JSON.stringify({ id: pending.id }),
        });
        showToast('已删除敏感词');
        closeConfirm();
        await loadWords(state.wordPage);
      } else {
        await request('/circle/remove', {
          method: 'POST',
          body: JSON.stringify({ id: pending.id }),
        });
        showToast('已删除违规动态');
        closeConfirm();
        await loadList(state.page);
      }
    } catch (err) {
      showToast(err.message);
      confirmBtn.disabled = false;
      confirmBtn.textContent = '确认删除';
    }
  });

  // ===== Init =====
  if (getToken()) {
    showApp();
    switchView('review');
    loadList(1);
  } else {
    showLogin();
  }
})();
