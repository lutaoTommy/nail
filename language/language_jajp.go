package language

var languageMapJa map[string]string

/*初始化*/
func initLanguageJajp() {
	languageMapJa = make(map[string]string)
	/*user*/
	languageMapJa["E_NO_IV"] = "初期化ベクトルを入力してください"
	languageMapJa["E_NO_CODE"] = "認証コードを入力してください"
	languageMapJa["E_NO_DATA"] = "データを入力してください"
	languageMapJa["E_NO_PHONE"] = "電話番号を入力してください"
	languageMapJa["E_NO_PWD"] = "パスワードを入力してください"
	languageMapJa["E_INVALID_USER"] = "ユーザーが存在しません"
	languageMapJa["E_INVALID_PWD"] = "パスワードが間違っています"
	languageMapJa["E_NO_NAME"] = "名前を入力してください"
	languageMapJa["E_USER_EXIST"] = "アカウントは既に存在します"
	languageMapJa["E_NO_TOKEN"] = "先にログインしてください"
	languageMapJa["E_NO_AUTH"] = "権限がありません"
	languageMapJa["E_NO_AVATAR"] = "アバターが存在しません"
	languageMapJa["E_NO_EMAIL"] = "メールアドレスを入力してください"
	languageMapJa["E_NO_CERT"] = "認証コードを入力してください"
	languageMapJa["E_INVALID_EMAIL"] = "メールアドレスの形式が正しくありません"
	languageMapJa["E_DATA_TOO_SHORT"] = "暗号化データが短すぎます"
	languageMapJa["E_INVALID_IV"] = "初期化ベクトルが無効です"
	languageMapJa["E_INVALID_SESSION_KEY"] = "セッションキーが無効です"
	languageMapJa["E_DECRYPT_FAIL"] = "復号に失敗しました。認証コードと電話番号が同じセッションのものであることを確認してください"
	languageMapJa["E_INVALID_PHONE"] = "電話番号の形式が正しくありません"
	languageMapJa["VERIFICATION_CODE"] = "認証コード"
	languageMapJa["DO_NOT_TELL"] = "他の人に教えないでください"
	languageMapJa["E_CERT_ERR"] = "認証コードが正しくありません"
	languageMapJa["E_CERT_EXPIRED"] = "認証コードの有効期限が切れました。新しく取得してください"
	languageMapJa["E_CERT_FIRST"] = "先に認証コードを取得してください"
	languageMapJa["E_VERIFICATION_QUICKLY"] = "認証コードの取得が頻繁すぎます。しばらくお待ちください"
	languageMapJa["E_ACCOUNT_LOCKED"] = "試行回数が多すぎます。しばらくしてから再度お試しください"
	languageMapJa["E_ACCOUNT_LOCKED_RETRY"] = "試行回数が多すぎます。%d秒後にもう一度お試しください"
	languageMapJa["E_RETRIES_REMAINING"] = "。残り %d 回試行できます"
	languageMapJa["E_VERIFICATION_LIMIT"] = "認証コードのリクエストが多すぎます。しばらくしてから再度お試しください"
	/*common*/
	languageMapJa["E_TOO_LONG"] = "文字数が上限を超えています"
	languageMapJa["E_INVALID_PARAM"] = "パラメータが無効です"
	languageMapJa["E_FILE_NOT_FOUND"] = "ファイルが見つかりません"
	/*color*/
	languageMapJa["E_INVALID_COLOR"] = "該当する色がありません"
	languageMapJa["E_INVALID_FILE"] = "ファイルをアップロードしてください"
	languageMapJa["E_INVALID_PICTURE"] = "画像をアップロードしてください(JPG/PNG)"
	languageMapJa["E_AI_RECOMMEND_FAILED"] = "色の推薦サービスは一時的に利用できません。しばらくしてから再度お試しください"
	/*device*/
	languageMapJa["E_NO_ID"] = "IDを入力してください"
	languageMapJa["E_NO_MAC"] = "MACアドレスを入力してください"
	languageMapJa["E_NO_INFO"] = "該当する記録がありません"
	languageMapJa["E_NO_DEVICE"] = "該当するデバイスがありません"
	languageMapJa["E_INVALID_MAC"] = "正しい形式のMACアドレスを入力してください"
	/*suggest*/
	languageMapJa["E_NO_CONTENT"] = "内容を入力してください"
	languageMapJa["E_NO_SUGGEST"] = "該当する提案がありません"
	/*word*/
	languageMapJa["E_NO_WORD"] = "該当する禁止ワードがありません"
	languageMapJa["E_WORD_EXIST"] = "禁止ワードは既に存在します"
	/*follow*/
	languageMapJa["E_FOLLOWED"] = "フォロー済み"
	languageMapJa["E_UNFOLLOWED"] = "フォロー解除済み"
	languageMapJa["E_NOT_FOLLOWING"] = "フォローしていません"
	languageMapJa["E_NO_SELF_FOLLOW"] = "自分自身はフォローできません"
	/*circle*/
	languageMapJa["E_NO_CIRCLE_POST"] = "投稿が存在しません"
	languageMapJa["E_INVALID_REVIEW_STATUS"] = "審査ステータスが無効です"
	languageMapJa["E_SENSITIVEWORD"] = "メッセージにセンシティブな情報が含まれています。友好的な交流環境を維持するため、投稿前に該当する表現を調整してください"
	/*topic*/
	languageMapJa["E_NO_TOPIC"] = "トピックが存在しません"
	languageMapJa["E_TOPIC_EXIST"] = "トピックは既に存在します"
	/*activity*/
	languageMapJa["E_NO_ACTIVITY"] = "イベントが存在しません"
	languageMapJa["E_NO_TITLE"] = "タイトルを入力してください"
	languageMapJa["E_NO_TIME"] = "時間を入力してください"
	languageMapJa["E_ALREADY_JOINED"] = "このイベントには既に参加しています"
	languageMapJa["E_NOT_JOINED"] = "このイベントには参加していません"
	languageMapJa["E_ALREADY_VOTED"] = "既に投票済みです"
	languageMapJa["E_NOT_VOTED"] = "投票していません"
	/*comment*/
	languageMapJa["E_NO_COMMENT"] = "コメントが存在しません"
	/*like*/
	languageMapJa["E_NO_LIKE"] = "いいねが存在しません"
	/*collect*/
	languageMapJa["E_NO_COLLECT"] = "お気に入りに追加していません"
	/*ota*/
	languageMapJa["E_NO_VERSION"] = "バージョンを入力してください"
	/*voice*/
	languageMapJa["E_NO_VOICE_TEXT"] = "音声コマンドを入力してください"
	languageMapJa["E_NO_COLOR_NAME"] = "色名が認識されませんでした"
	languageMapJa["V_UNKNOWN_COMMAND"] = "不明な音声コマンドです"
	languageMapJa["V_ACTION_HELP"] = "コマンドを認識できません。お試しください：\n· 色を変更：一番目をミッドナイトネイビーに変更"
	/*checkin*/
	languageMapJa["E_ALREADY_CHECKIN"] = "本日は既にチェックイン済みです"
	/*anti_fake*/
	languageMapJa["E_INVALID_PRODUCT_MODEL"] = "製品型番の形式が無効です（英数字2文字が必要）"
	languageMapJa["E_INVALID_BATCH"] = "生産ロットの形式が無効です（英数字6文字が必要）"
	languageMapJa["E_INVALID_NUMBER"] = "番号の形式が無効です（英数字2文字が必要）"
	languageMapJa["E_INVALID_ANTI_FAKE_CODE"] = "偽造防止コードの形式が無効です（英数字24文字が必要）"
	languageMapJa["E_GENERATE_CODE_FAILED"] = "偽造防止コードの生成に失敗しました"
	languageMapJa["E_ANTI_FAKE_CODE_ACTIVATED"] = "偽造防止コードは既に有効化されています"
	languageMapJa["E_NO_ANTI_FAKE_CODE"] = "偽造防止コードが存在しません"
	languageMapJa["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "この偽造防止コードは既に他のユーザーが検証済みです"
	languageMapJa["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "偽造防止コードが未有効化のため検証できません"
	/*mail*/
	languageMapJa["MAIL_VERIFY_SUBJECT"] = "TintaShift 認証コード"
	languageMapJa["MAIL_VERIFY_TITLE"] = "TintaShift へようこそ"
	languageMapJa["MAIL_VERIFY_DESC"] = "こんにちは。アカウントのセキュリティ認証を行っています。認証コードは次のとおりです："
	languageMapJa["MAIL_VERIFY_SECURITY_TITLE"] = "セキュリティのヒント："
	languageMapJa["MAIL_VERIFY_SECURITY_DESC"] = "TintaShift のスタッフがこの認証コードを求めることはありません。誰にも教えないでください。"
	languageMapJa["MAIL_VERIFY_IGNORE"] = "この認証コードをリクエストしていない場合は、このメールを無視してください。"
	languageMapJa["MAIL_COPYRIGHT"] = "© 2026 TintaShift. All rights reserved."
	/*social login*/
	languageMapJa["E_NO_GOOGLE_ACCESS_TOKEN"] = "Google アクセストークンが必要です"
	languageMapJa["E_GOOGLE_NOT_CONFIGURED"] = "サーバーに Google ログインが設定されていません"
	languageMapJa["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Google アクセストークンが無効または期限切れです"
	languageMapJa["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "Google のメールアドレスが未確認のため、ログインできません"
	languageMapJa["E_GOOGLE_EMAIL_BOUND_OTHER"] = "このメールアドレスは別の Google アカウントに既にリンクされています"
	languageMapJa["E_NO_APPLE_ID_TOKEN"] = "Apple identity token が必要です"
	languageMapJa["E_APPLE_NOT_CONFIGURED"] = "サーバーに Apple ログインが設定されていません"
	languageMapJa["E_APPLE_ID_TOKEN_INVALID"] = "Apple identity token が無効または期限切れです"
	languageMapJa["E_APPLE_EMAIL_NOT_VERIFIED"] = "Apple のメールアドレスが未確認のため、ログインできません"
	languageMapJa["E_APPLE_EMAIL_BOUND_OTHER"] = "このメールアドレスは別の Apple アカウントに既にリンクされています"
	languageMapJa["E_SOCIAL_NOT_CONFIGURED"] = "サーバーにソーシャルログインの署名シークレットが設定されていません"
	languageMapJa["E_NO_PROVIDER"] = "provider が必要です"
	languageMapJa["E_SOCIAL_PROVIDER_INVALID"] = "サポートされていないソーシャルログイン provider です"
	languageMapJa["E_NO_OPENID"] = "openid が必要です"
	languageMapJa["E_NO_NONCE"] = "nonce が必要です"
	languageMapJa["E_NO_TIMESTAMP"] = "timestamp が必要です"
	languageMapJa["E_NO_SIGNATURE"] = "signature が必要です"
	languageMapJa["E_NO_DEVICE_ID"] = "device_id が必要です"
	languageMapJa["E_REQUEST_EXPIRED"] = "リクエストの有効期限が切れました。もう一度お試しください"
	languageMapJa["E_CHALLENGE_INVALID"] = "challenge が無効です"
	languageMapJa["E_CHALLENGE_EXPIRED"] = "challenge の有効期限が切れました"
	languageMapJa["E_CHALLENGE_USED"] = "challenge は既に使用されています"
	languageMapJa["E_SIGNATURE_INVALID"] = "署名の検証に失敗しました"
	languageMapJa["E_SOCIAL_EMAIL_BOUND_OTHER"] = "このメールアドレスは別のソーシャルアカウントに既にリンクされています"
	languageMapJa["E_INVALID_VIDEO"] = "正しい動画をアップロードしてください(MP4/MOV)"
	registerLocale("ja-JP", languageMapJa)
}
