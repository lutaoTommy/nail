package language

var languageMapZhhant map[string]string

/*初始化*/
func initLanguageZhhant() {
	languageMapZhhant = make(map[string]string)
	/*user*/
	languageMapZhhant["E_NO_IV"] = "請填寫初始向量"
	languageMapZhhant["E_NO_CODE"] = "請填寫授權碼"
	languageMapZhhant["E_NO_DATA"] = "請填寫數據"
	languageMapZhhant["E_NO_PHONE"] = "請輸入手機號"
	languageMapZhhant["E_NO_PWD"] = "請輸入密碼"
	languageMapZhhant["E_INVALID_USER"] = "無此用戶"
	languageMapZhhant["E_INVALID_PWD"] = "密碼錯誤"
	languageMapZhhant["E_NO_NAME"] = "請輸入名稱"
	languageMapZhhant["E_USER_EXIST"] = "帳號已存在"
	languageMapZhhant["E_NO_TOKEN"] = "請先登入"
	languageMapZhhant["E_NO_AUTH"] = "沒有權限"
	languageMapZhhant["E_NO_AVATAR"] = "無此頭像"
	languageMapZhhant["E_NO_EMAIL"] = "請填寫郵箱"
	languageMapZhhant["E_NO_CERT"] = "請輸入驗證碼"
	languageMapZhhant["E_INVALID_EMAIL"] = "郵箱格式錯誤"
	languageMapZhhant["E_DATA_TOO_SHORT"] = "加密數據過短"
	languageMapZhhant["E_INVALID_IV"] = "初始向量無效"
	languageMapZhhant["E_INVALID_SESSION_KEY"] = "會話密鑰無效"
	languageMapZhhant["E_DECRYPT_FAIL"] = "解密失敗，請確保授權碼與手機號為同一次授權"
	languageMapZhhant["E_INVALID_PHONE"] = "手機號格式錯誤"
	languageMapZhhant["VERIFICATION_CODE"] = "驗證碼"
	languageMapZhhant["DO_NOT_TELL"] = "請勿告訴他人"
	languageMapZhhant["E_CERT_ERR"] = "驗證碼錯誤"
	languageMapZhhant["E_CERT_EXPIRED"] = "驗證碼已過期，請重新獲取"
	languageMapZhhant["E_CERT_FIRST"] = "請先獲取驗證碼"
	languageMapZhhant["E_VERIFICATION_QUICKLY"] = "獲取驗證碼太頻繁"
	languageMapZhhant["E_ACCOUNT_LOCKED"] = "嘗試次數過多，請稍後再試"
	languageMapZhhant["E_ACCOUNT_LOCKED_RETRY"] = "嘗試次數過多，請 %d 秒後再試"
	languageMapZhhant["E_RETRIES_REMAINING"] = "，還可嘗試 %d 次"
	languageMapZhhant["E_VERIFICATION_LIMIT"] = "請求驗證碼過於頻繁，請稍後再試"
	/*common*/
	languageMapZhhant["E_TOO_LONG"] = "字符長度超過限制"
	languageMapZhhant["E_INVALID_PARAM"] = "參數錯誤"
	languageMapZhhant["E_FILE_NOT_FOUND"] = "文件不存在"
	/*color*/
	languageMapZhhant["E_INVALID_COLOR"] = "無此顏色"
	languageMapZhhant["E_INVALID_FILE"] = "請上傳文件"
	languageMapZhhant["E_INVALID_PICTURE"] = "請上傳圖片(JPG/PNG)"
	languageMapZhhant["E_AI_RECOMMEND_FAILED"] = "顏色推薦服務暫不可用，請稍後再試"
	/*device*/
	languageMapZhhant["E_NO_ID"] = "請填寫ID"
	languageMapZhhant["E_NO_MAC"] = "請填寫MAC地址"
	languageMapZhhant["E_NO_INFO"] = "無此記錄"
	languageMapZhhant["E_NO_DEVICE"] = "無此設備"
	languageMapZhhant["E_INVALID_MAC"] = "請填寫正確格式的MAC地址"
	/*suggest*/
	languageMapZhhant["E_NO_CONTENT"] = "請填寫內容"
	languageMapZhhant["E_NO_SUGGEST"] = "無此建議"
	/*word*/
	languageMapZhhant["E_NO_WORD"] = "無此敏感詞"
	languageMapZhhant["E_WORD_EXIST"] = "敏感詞已存在"
	/*follow*/
	languageMapZhhant["E_FOLLOWED"] = "已關注"
	languageMapZhhant["E_UNFOLLOWED"] = "已取消關注"
	languageMapZhhant["E_NOT_FOLLOWING"] = "未關注"
	languageMapZhhant["E_NO_SELF_FOLLOW"] = "不能關注自己"
	/*circle*/
	languageMapZhhant["E_NO_CIRCLE_POST"] = "無此動態"
	languageMapZhhant["E_INVALID_REVIEW_STATUS"] = "審核狀態不合法"
	languageMapZhhant["E_SENSITIVEWORD"] = "您的發言中包含敏感信息，為維護友好的交流氛圍，請調整相關表述後再發布"
	/*topic*/
	languageMapZhhant["E_NO_TOPIC"] = "無此話題"
	languageMapZhhant["E_TOPIC_EXIST"] = "話題已存在"
	/*activity*/
	languageMapZhhant["E_NO_ACTIVITY"] = "無此活動"
	languageMapZhhant["E_NO_TITLE"] = "請填寫標題"
	languageMapZhhant["E_NO_TIME"] = "請填寫時間"
	languageMapZhhant["E_ALREADY_JOINED"] = "已參與此活動"
	languageMapZhhant["E_NOT_JOINED"] = "未參與此活動"
	languageMapZhhant["E_ALREADY_VOTED"] = "已投過票"
	languageMapZhhant["E_NOT_VOTED"] = "未投過票"
	/*comment*/
	languageMapZhhant["E_NO_COMMENT"] = "無此評論"
	/*like*/
	languageMapZhhant["E_NO_LIKE"] = "無此點讚"
	/*collect*/
	languageMapZhhant["E_NO_COLLECT"] = "未收藏過"
	/*ota*/
	languageMapZhhant["E_NO_VERSION"] = "請填寫版本"
	/*voice*/
	languageMapZhhant["E_NO_VOICE_TEXT"] = "請輸入語音指令"
	languageMapZhhant["E_NO_COLOR_NAME"] = "未識別到顏色名稱"
	languageMapZhhant["V_UNKNOWN_COMMAND"] = "未識別的語音指令"
	languageMapZhhant["V_ACTION_HELP"] = "無法識別指令。請嘗試：\n· 換顏色：第一個換成午夜藍"
	/*checkin*/
	languageMapZhhant["E_ALREADY_CHECKIN"] = "今日已簽到"
	/*anti_fake*/
	languageMapZhhant["E_INVALID_PRODUCT_MODEL"] = "產品型號格式錯誤（需2位字母或數字）"
	languageMapZhhant["E_INVALID_BATCH"] = "生產批次格式錯誤（需6位字母或數字）"
	languageMapZhhant["E_INVALID_NUMBER"] = "編號格式錯誤（需2位字母或數字）"
	languageMapZhhant["E_INVALID_ANTI_FAKE_CODE"] = "防偽碼格式錯誤（需24位字母或數字）"
	languageMapZhhant["E_GENERATE_CODE_FAILED"] = "防偽碼生成失敗"
	languageMapZhhant["E_ANTI_FAKE_CODE_ACTIVATED"] = "防偽碼已激活"
	languageMapZhhant["E_NO_ANTI_FAKE_CODE"] = "防偽碼不存在"
	languageMapZhhant["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "該防偽碼已被其他用戶驗證"
	languageMapZhhant["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "防偽碼未激活，無法驗證"
	/*mail*/
	languageMapZhhant["MAIL_VERIFY_SUBJECT"] = "TintaShift 驗證碼"
	languageMapZhhant["MAIL_VERIFY_TITLE"] = "歡迎使用 TintaShift"
	languageMapZhhant["MAIL_VERIFY_DESC"] = "您好，您正在進行帳戶安全驗證，您的驗證碼是："
	languageMapZhhant["MAIL_VERIFY_SECURITY_TITLE"] = "安全提示："
	languageMapZhhant["MAIL_VERIFY_SECURITY_DESC"] = "TintaShift 的工作人員絕不會向您索要此驗證碼，請勿告知任何人。"
	languageMapZhhant["MAIL_VERIFY_IGNORE"] = "如果您未曾請求此驗證碼，請忽略此郵件。"
	languageMapZhhant["MAIL_COPYRIGHT"] = "© 2026 TintaShift. 保留所有權利。"
	/*social login*/
	languageMapZhhant["E_NO_GOOGLE_ACCESS_TOKEN"] = "請提供 Google access token"
	languageMapZhhant["E_GOOGLE_NOT_CONFIGURED"] = "伺服器未設定 Google 登入"
	languageMapZhhant["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Google access token 無效或已過期"
	languageMapZhhant["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "Google 郵箱未驗證，無法登入"
	languageMapZhhant["E_GOOGLE_EMAIL_BOUND_OTHER"] = "該郵箱已綁定其他 Google 帳號"
	languageMapZhhant["E_NO_APPLE_ID_TOKEN"] = "請提供 Apple identity token"
	languageMapZhhant["E_APPLE_NOT_CONFIGURED"] = "伺服器未設定 Apple 登入"
	languageMapZhhant["E_APPLE_ID_TOKEN_INVALID"] = "Apple identity token 無效或已過期"
	languageMapZhhant["E_APPLE_EMAIL_NOT_VERIFIED"] = "Apple 郵箱未驗證，無法登入"
	languageMapZhhant["E_APPLE_EMAIL_BOUND_OTHER"] = "該郵箱已綁定其他 Apple 帳號"
	languageMapZhhant["E_SOCIAL_NOT_CONFIGURED"] = "伺服器未設定社交登入簽名密鑰"
	languageMapZhhant["E_NO_PROVIDER"] = "請提供 provider"
	languageMapZhhant["E_SOCIAL_PROVIDER_INVALID"] = "不支援的社交登入 provider"
	languageMapZhhant["E_NO_OPENID"] = "請提供 openid"
	languageMapZhhant["E_NO_NONCE"] = "請提供 nonce"
	languageMapZhhant["E_NO_TIMESTAMP"] = "請提供 timestamp"
	languageMapZhhant["E_NO_SIGNATURE"] = "請提供 signature"
	languageMapZhhant["E_NO_DEVICE_ID"] = "請提供 device_id"
	languageMapZhhant["E_REQUEST_EXPIRED"] = "請求已過期，請重試"
	languageMapZhhant["E_CHALLENGE_INVALID"] = "challenge 無效"
	languageMapZhhant["E_CHALLENGE_EXPIRED"] = "challenge 已過期"
	languageMapZhhant["E_CHALLENGE_USED"] = "challenge 已被使用"
	languageMapZhhant["E_SIGNATURE_INVALID"] = "簽名校驗失敗"
	languageMapZhhant["E_SOCIAL_EMAIL_BOUND_OTHER"] = "該郵箱已綁定其他社交帳號"
	languageMapZhhant["E_INVALID_VIDEO"] = "請上傳正確的影片(MP4/MOV)"
	registerLocale("zh-Hant", languageMapZhhant)
}
