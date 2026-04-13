package language

var languageMapEn map[string]string
/*初始化*/
func initLanguageEnus() {
	languageMapEn = make(map[string]string)
	/*user*/
	languageMapEn["E_NO_IV"] = "Please enter the initialization vector"
	languageMapEn["E_NO_CODE"] = "Please enter the authorization code"
	languageMapEn["E_NO_DATA"] = "Please enter the data"
	languageMapEn["E_NO_PHONE"] = "Please enter the phone number"
	languageMapEn["E_NO_PWD"] = "Please enter the password"
	languageMapEn["E_INVALID_USER"] = "No this user"
	languageMapEn["E_INVALID_PWD"] = "Password error"
	languageMapEn["E_NO_NAME"] = "Please enter the name"
	languageMapEn["E_USER_EXIST"] = "User already exists"
	languageMapEn["E_NO_TOKEN"] = "Please login first"
	languageMapEn["E_NO_AUTH"] = "No auth"
	languageMapEn["E_NO_AVATAR"] = "No this avatar"
	languageMapEn["E_NO_EMAIL"] = "Please Enter the email"
	languageMapEn["E_NO_CERT"] = "Please Enter the certification"
	languageMapEn["E_INVALID_EMAIL"] = "Email format error"
	languageMapEn["E_DATA_TOO_SHORT"] = "The encrypted data is too short"
	languageMapEn["E_INVALID_IV"] = "Invalid IV"
	languageMapEn["E_INVALID_SESSION_KEY"] = "Invalid session key"
	languageMapEn["E_DECRYPT_FAIL"] = "Decrypt failed, ensure code and phone data are from the same session"
	languageMapEn["E_INVALID_PHONE"] = "Phone number format error"
	languageMapEn["VERIFICATION_CODE"] = "verification code"
	languageMapEn["DO_NOT_TELL"] = "Please do not tell others"
	/*mail*/
	languageMapEn["MAIL_VERIFY_SUBJECT"] = "TintaShift Verification Code"
	languageMapEn["MAIL_VERIFY_TITLE"] = "Welcome to TintaShift"
	languageMapEn["MAIL_VERIFY_DESC"] = "Hello, you are performing an account security verification. Your code is:"
	languageMapEn["MAIL_VERIFY_SECURITY_TITLE"] = "Security tip: "
	languageMapEn["MAIL_VERIFY_SECURITY_DESC"] = "TintaShift staff will never ask for this code. Do not share it with anyone."
	languageMapEn["MAIL_VERIFY_IGNORE"] = "If you didn’t request this code, please ignore this email."
	languageMapEn["MAIL_COPYRIGHT"] = "© 2026 TintaShift. All rights reserved."
	languageMapEn["E_CERT_ERR"] = "Certification number error"
	languageMapEn["E_CERT_FIRST"] = "Please get certification number first"
	languageMapEn["E_VERIFICATION_QUICKLY"] = "You are requesting verification codes too frequently. Please wait a moment"
	languageMapEn["E_ACCOUNT_LOCKED"] = "Too many attempts. Please try again later"
	languageMapEn["E_VERIFICATION_LIMIT"] = "Too many verification code requests. Please try again later"
	/*common*/
	languageMapEn["E_TOO_LONG"] = "Character length exceeds the limit"
	languageMapEn["E_INVALID_PARAM"] = "Invalid parameter"
	languageMapEn["E_FILE_NOT_FOUND"] = "File not found"
	/*color*/
	languageMapEn["E_INVALID_COLOR"] = "No this color"
	languageMapEn["E_INVALID_FILE"] = "Please upload correct file"
	languageMapEn["E_INVALID_PICTURE"] = "Please upload picture(JPG/PNG/BMP)"
	/*device*/
	languageMapEn["E_NO_ID"] = "Please enter the ID"
	languageMapEn["E_NO_MAC"] = "Please enter the MAC address"
	languageMapEn["E_NO_INFO"] = "No this record"
	languageMapEn["E_NO_DEVICE"] = "No this device"
	languageMapEn["E_INVALID_MAC"] = "Please enter the MAC address in the correct format"
	/*suggest*/
	languageMapEn["E_NO_CONTENT"] = "Please enter the content"
	languageMapEn["E_NO_SUGGEST"] = "No this suggest"
	/*word*/
	languageMapEn["E_NO_WORD"] = "No this sensitive word"
	languageMapEn["E_WORD_EXIST"] = "The sensitive word already exists"
	/*follow*/
	languageMapEn["E_FOLLOWED"] = "Already followed​"
	languageMapEn["E_UNFOLLOWED"] = "Already unfollowed"
	languageMapEn["E_NOT_FOLLOWING"] = "Not following"
	languageMapEn["E_NO_SELF_FOLLOW"] = "Cannot follow yourself"
	/*circle*/
	languageMapEn["E_NO_CIRCLE_POST"] = "Post Not Found"
	languageMapEn["E_SENSITIVEWORD"] = "Your message contains sensitive information. To maintain a friendly communication environment, please adjust the relevant expressions before reposting"
	/*comment*/
	languageMapEn["E_NO_COMMENT"] = "Comment Not Found"
	/*like*/
	languageMapEn["E_NO_LIKE"] = "Like Not Found"
	/*collect*/
	languageMapEn["E_NO_COLLECT"] = "Not collected"
	/*ota*/
	languageMapEn["E_NO_VERSION"] = "Please enter the version"
}