package language

var languageMapKo map[string]string

/*初始化*/
func initLanguageKokr() {
	languageMapKo = make(map[string]string)
	/*user*/
	languageMapKo["E_NO_IV"] = "초기화 벡터를 입력하세요"
	languageMapKo["E_NO_CODE"] = "인증 코드를 입력하세요"
	languageMapKo["E_NO_DATA"] = "데이터를 입력하세요"
	languageMapKo["E_NO_PHONE"] = "전화번호를 입력하세요"
	languageMapKo["E_NO_PWD"] = "비밀번호를 입력하세요"
	languageMapKo["E_INVALID_USER"] = "사용자가 존재하지 않습니다"
	languageMapKo["E_INVALID_PWD"] = "비밀번호가 잘못되었습니다"
	languageMapKo["E_NO_NAME"] = "이름을 입력하세요"
	languageMapKo["E_USER_EXIST"] = "이미 존재하는 계정입니다"
	languageMapKo["E_NO_TOKEN"] = "먼저 로그인하세요"
	languageMapKo["E_NO_AUTH"] = "권한이 없습니다"
	languageMapKo["E_NO_AVATAR"] = "아바타가 존재하지 않습니다"
	languageMapKo["E_NO_EMAIL"] = "이메일을 입력하세요"
	languageMapKo["E_NO_CERT"] = "인증 코드를 입력하세요"
	languageMapKo["E_INVALID_EMAIL"] = "이메일 형식이 잘못되었습니다"
	languageMapKo["E_DATA_TOO_SHORT"] = "암호화된 데이터가 너무 짧습니다"
	languageMapKo["E_INVALID_IV"] = "초기화 벡터가 유효하지 않습니다"
	languageMapKo["E_INVALID_SESSION_KEY"] = "세션 키가 유효하지 않습니다"
	languageMapKo["E_DECRYPT_FAIL"] = "복호화에 실패했습니다. 인증 코드와 전화번호가 같은 세션의 것인지 확인하세요"
	languageMapKo["E_INVALID_PHONE"] = "전화번호 형식이 잘못되었습니다"
	languageMapKo["VERIFICATION_CODE"] = "인증 코드"
	languageMapKo["DO_NOT_TELL"] = "다른 사람에게 알려주지 마세요"
	languageMapKo["E_CERT_ERR"] = "인증 코드가 잘못되었습니다"
	languageMapKo["E_CERT_EXPIRED"] = "인증 코드가 만료되었습니다. 새로 요청하세요"
	languageMapKo["E_CERT_FIRST"] = "먼저 인증 코드를 요청하세요"
	languageMapKo["E_VERIFICATION_QUICKLY"] = "인증 코드 요청이 너무 잦습니다. 잠시 후 다시 시도하세요"
	languageMapKo["E_ACCOUNT_LOCKED"] = "시도 횟수가 너무 많습니다. 잠시 후 다시 시도하세요"
	languageMapKo["E_ACCOUNT_LOCKED_RETRY"] = "시도 횟수가 너무 많습니다. %d초 후에 다시 시도하세요"
	languageMapKo["E_RETRIES_REMAINING"] = ". 남은 시도 횟수: %d회"
	languageMapKo["E_VERIFICATION_LIMIT"] = "인증 코드 요청이 너무 잦습니다. 잠시 후 다시 시도하세요"
	/*common*/
	languageMapKo["E_TOO_LONG"] = "문자 길이가 제한을 초과했습니다"
	languageMapKo["E_INVALID_PARAM"] = "잘못된 매개변수입니다"
	languageMapKo["E_FILE_NOT_FOUND"] = "파일을 찾을 수 없습니다"
	/*color*/
	languageMapKo["E_INVALID_COLOR"] = "해당 색상이 없습니다"
	languageMapKo["E_INVALID_FILE"] = "파일을 업로드하세요"
	languageMapKo["E_INVALID_PICTURE"] = "이미지를 업로드하세요(JPG/PNG)"
	languageMapKo["E_AI_RECOMMEND_FAILED"] = "색상 추천 서비스를 일시적으로 사용할 수 없습니다. 잠시 후 다시 시도하세요"
	/*device*/
	languageMapKo["E_NO_ID"] = "ID를 입력하세요"
	languageMapKo["E_NO_MAC"] = "MAC 주소를 입력하세요"
	languageMapKo["E_NO_INFO"] = "해당 기록이 없습니다"
	languageMapKo["E_NO_DEVICE"] = "해당 기기가 없습니다"
	languageMapKo["E_INVALID_MAC"] = "올바른 형식의 MAC 주소를 입력하세요"
	/*suggest*/
	languageMapKo["E_NO_CONTENT"] = "내용을 입력하세요"
	languageMapKo["E_NO_SUGGEST"] = "해당 제안이 없습니다"
	/*word*/
	languageMapKo["E_NO_WORD"] = "해당 금칙어가 없습니다"
	languageMapKo["E_WORD_EXIST"] = "이미 존재하는 금칙어입니다"
	/*follow*/
	languageMapKo["E_FOLLOWED"] = "팔로우됨"
	languageMapKo["E_UNFOLLOWED"] = "언팔로우됨"
	languageMapKo["E_NOT_FOLLOWING"] = "팔로우하지 않음"
	languageMapKo["E_NO_SELF_FOLLOW"] = "자기 자신은 팔로우할 수 없습니다"
	/*circle*/
	languageMapKo["E_NO_CIRCLE_POST"] = "해당 게시물이 없습니다"
	languageMapKo["E_INVALID_REVIEW_STATUS"] = "잘못된 검토 상태입니다"
	languageMapKo["E_SENSITIVEWORD"] = "메시지에 민감한 정보가 포함되어 있습니다. 친근한 소통 환경을 위해 게시하기 전에 관련 표현을 조정하세요"
	/*topic*/
	languageMapKo["E_NO_TOPIC"] = "해당 주제가 없습니다"
	languageMapKo["E_TOPIC_EXIST"] = "이미 존재하는 주제입니다"
	/*activity*/
	languageMapKo["E_NO_ACTIVITY"] = "해당 활동이 없습니다"
	languageMapKo["E_NO_TITLE"] = "제목을 입력하세요"
	languageMapKo["E_NO_TIME"] = "시간을 입력하세요"
	languageMapKo["E_ALREADY_JOINED"] = "이미 이 활동에 참여했습니다"
	languageMapKo["E_NOT_JOINED"] = "이 활동에 참여하지 않았습니다"
	languageMapKo["E_ALREADY_VOTED"] = "이미 투표했습니다"
	languageMapKo["E_NOT_VOTED"] = "투표하지 않았습니다"
	/*comment*/
	languageMapKo["E_NO_COMMENT"] = "해당 댓글이 없습니다"
	/*like*/
	languageMapKo["E_NO_LIKE"] = "해당 좋아요가 없습니다"
	/*collect*/
	languageMapKo["E_NO_COLLECT"] = "북마크에 추가하지 않았습니다"
	/*ota*/
	languageMapKo["E_NO_VERSION"] = "버전을 입력하세요"
	/*voice*/
	languageMapKo["E_NO_VOICE_TEXT"] = "음성 명령을 입력하세요"
	languageMapKo["E_NO_COLOR_NAME"] = "색상 이름이 인식되지 않았습니다"
	languageMapKo["V_UNKNOWN_COMMAND"] = "알 수 없는 음성 명령입니다"
	languageMapKo["V_ACTION_HELP"] = "명령을 인식할 수 없습니다. 다음을 시도해 보세요:\n· 색상 변경: 첫 번째를 미드나잇 네이비로 변경"
	/*checkin*/
	languageMapKo["E_ALREADY_CHECKIN"] = "오늘은 이미 체크인했습니다"
	/*anti_fake*/
	languageMapKo["E_INVALID_PRODUCT_MODEL"] = "제품 모델 형식이 잘못되었습니다 (영숫자 2자리 필요)"
	languageMapKo["E_INVALID_BATCH"] = "생산 배치 형식이 잘못되었습니다 (영숫자 6자리 필요)"
	languageMapKo["E_INVALID_NUMBER"] = "번호 형식이 잘못되었습니다 (영숫자 2자리 필요)"
	languageMapKo["E_INVALID_ANTI_FAKE_CODE"] = "위조 방지 코드 형식이 잘못되었습니다 (영숫자 24자리 필요)"
	languageMapKo["E_GENERATE_CODE_FAILED"] = "위조 방지 코드 생성에 실패했습니다"
	languageMapKo["E_ANTI_FAKE_CODE_ACTIVATED"] = "위조 방지 코드가 이미 활성화되었습니다"
	languageMapKo["E_NO_ANTI_FAKE_CODE"] = "위조 방지 코드가 존재하지 않습니다"
	languageMapKo["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "이 위조 방지 코드는 다른 사용자가 이미 확인했습니다"
	languageMapKo["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "위조 방지 코드가 활성화되지 않아 확인할 수 없습니다"
	/*mail*/
	languageMapKo["MAIL_VERIFY_SUBJECT"] = "TintaShift 인증 코드"
	languageMapKo["MAIL_VERIFY_TITLE"] = "TintaShift에 오신 것을 환영합니다"
	languageMapKo["MAIL_VERIFY_DESC"] = "안녕하세요. 계정 보안 인증을 진행 중입니다. 인증 코드는 다음과 같습니다:"
	languageMapKo["MAIL_VERIFY_SECURITY_TITLE"] = "보안 팁: "
	languageMapKo["MAIL_VERIFY_SECURITY_DESC"] = "TintaShift 직원은 이 인증 코드를 절대 요청하지 않습니다. 누구에게도 공유하지 마세요."
	languageMapKo["MAIL_VERIFY_IGNORE"] = "이 인증 코드를 요청하지 않으셨다면 이 이메일을 무시하세요."
	languageMapKo["MAIL_COPYRIGHT"] = "© 2026 TintaShift. All rights reserved."
	/*social login*/
	languageMapKo["E_NO_GOOGLE_ACCESS_TOKEN"] = "Google 액세스 토큰이 필요합니다"
	languageMapKo["E_GOOGLE_NOT_CONFIGURED"] = "서버에 Google 로그인이 구성되어 있지 않습니다"
	languageMapKo["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Google 액세스 토큰이 유효하지 않거나 만료되었습니다"
	languageMapKo["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "Google 이메일이 인증되지 않아 로그인할 수 없습니다"
	languageMapKo["E_GOOGLE_EMAIL_BOUND_OTHER"] = "이 이메일은 이미 다른 Google 계정에 연결되어 있습니다"
	languageMapKo["E_NO_APPLE_ID_TOKEN"] = "Apple identity token이 필요합니다"
	languageMapKo["E_APPLE_NOT_CONFIGURED"] = "서버에 Apple 로그인이 구성되어 있지 않습니다"
	languageMapKo["E_APPLE_ID_TOKEN_INVALID"] = "Apple identity token이 유효하지 않거나 만료되었습니다"
	languageMapKo["E_APPLE_EMAIL_NOT_VERIFIED"] = "Apple 이메일이 인증되지 않아 로그인할 수 없습니다"
	languageMapKo["E_APPLE_EMAIL_BOUND_OTHER"] = "이 이메일은 이미 다른 Apple 계정에 연결되어 있습니다"
	languageMapKo["E_SOCIAL_NOT_CONFIGURED"] = "서버에 소셜 로그인 서명 비밀키가 구성되어 있지 않습니다"
	languageMapKo["E_NO_PROVIDER"] = "provider가 필요합니다"
	languageMapKo["E_SOCIAL_PROVIDER_INVALID"] = "지원하지 않는 소셜 로그인 provider입니다"
	languageMapKo["E_NO_OPENID"] = "openid가 필요합니다"
	languageMapKo["E_NO_NONCE"] = "nonce가 필요합니다"
	languageMapKo["E_NO_TIMESTAMP"] = "timestamp가 필요합니다"
	languageMapKo["E_NO_SIGNATURE"] = "signature가 필요합니다"
	languageMapKo["E_NO_DEVICE_ID"] = "device_id가 필요합니다"
	languageMapKo["E_REQUEST_EXPIRED"] = "요청이 만료되었습니다. 다시 시도하세요"
	languageMapKo["E_CHALLENGE_INVALID"] = "challenge가 유효하지 않습니다"
	languageMapKo["E_CHALLENGE_EXPIRED"] = "challenge가 만료되었습니다"
	languageMapKo["E_CHALLENGE_USED"] = "challenge가 이미 사용되었습니다"
	languageMapKo["E_SIGNATURE_INVALID"] = "서명 검증에 실패했습니다"
	languageMapKo["E_SOCIAL_EMAIL_BOUND_OTHER"] = "이 이메일은 이미 다른 소셜 계정에 연결되어 있습니다"
	languageMapKo["E_INVALID_VIDEO"] = "올바른 동영상을 업로드하세요(MP4/MOV)"
	registerLocale("ko-KR", languageMapKo)
}
