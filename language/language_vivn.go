package language

var languageMapVi map[string]string

/*初始化*/
func initLanguageVivn() {
	languageMapVi = make(map[string]string)
	/*user*/
	languageMapVi["E_NO_IV"] = "Vui lòng nhập vector khởi tạo"
	languageMapVi["E_NO_CODE"] = "Vui lòng nhập mã ủy quyền"
	languageMapVi["E_NO_DATA"] = "Vui lòng nhập dữ liệu"
	languageMapVi["E_NO_PHONE"] = "Vui lòng nhập số điện thoại"
	languageMapVi["E_NO_PWD"] = "Vui lòng nhập mật khẩu"
	languageMapVi["E_INVALID_USER"] = "Không tồn tại người dùng này"
	languageMapVi["E_INVALID_PWD"] = "Mật khẩu không đúng"
	languageMapVi["E_NO_NAME"] = "Vui lòng nhập tên"
	languageMapVi["E_USER_EXIST"] = "Tài khoản đã tồn tại"
	languageMapVi["E_NO_TOKEN"] = "Vui lòng đăng nhập trước"
	languageMapVi["E_NO_AUTH"] = "Không có quyền"
	languageMapVi["E_NO_AVATAR"] = "Không tồn tại ảnh đại diện này"
	languageMapVi["E_NO_EMAIL"] = "Vui lòng nhập email"
	languageMapVi["E_NO_CERT"] = "Vui lòng nhập mã xác minh"
	languageMapVi["E_INVALID_EMAIL"] = "Định dạng email không đúng"
	languageMapVi["E_DATA_TOO_SHORT"] = "Dữ liệu mã hóa quá ngắn"
	languageMapVi["E_INVALID_IV"] = "Vector khởi tạo không hợp lệ"
	languageMapVi["E_INVALID_SESSION_KEY"] = "Khóa phiên không hợp lệ"
	languageMapVi["E_DECRYPT_FAIL"] = "Giải mã thất bại, vui lòng đảm bảo mã và số điện thoại cùng thuộc một phiên"
	languageMapVi["E_INVALID_PHONE"] = "Định dạng số điện thoại không đúng"
	languageMapVi["VERIFICATION_CODE"] = "mã xác minh"
	languageMapVi["DO_NOT_TELL"] = "Vui lòng không cho người khác biết"
	languageMapVi["E_CERT_ERR"] = "Mã xác minh không đúng"
	languageMapVi["E_CERT_EXPIRED"] = "Mã xác minh đã hết hạn. Vui lòng lấy mã mới"
	languageMapVi["E_CERT_FIRST"] = "Vui lòng lấy mã xác minh trước"
	languageMapVi["E_VERIFICATION_QUICKLY"] = "Bạn đang yêu cầu mã xác minh quá thường xuyên. Vui lòng đợi một lát"
	languageMapVi["E_ACCOUNT_LOCKED"] = "Quá nhiều lần thử. Vui lòng thử lại sau"
	languageMapVi["E_ACCOUNT_LOCKED_RETRY"] = "Quá nhiều lần thử. Vui lòng thử lại sau %d giây"
	languageMapVi["E_RETRIES_REMAINING"] = ". Còn %d lần thử"
	languageMapVi["E_VERIFICATION_LIMIT"] = "Quá nhiều yêu cầu mã xác minh. Vui lòng thử lại sau"
	/*common*/
	languageMapVi["E_TOO_LONG"] = "Độ dài ký tự vượt quá giới hạn"
	languageMapVi["E_INVALID_PARAM"] = "Tham số không hợp lệ"
	languageMapVi["E_FILE_NOT_FOUND"] = "Không tìm thấy tệp"
	/*color*/
	languageMapVi["E_INVALID_COLOR"] = "Không tồn tại màu này"
	languageMapVi["E_INVALID_FILE"] = "Vui lòng tải lên tệp đúng"
	languageMapVi["E_INVALID_PICTURE"] = "Vui lòng tải lên hình ảnh (JPG/PNG)"
	languageMapVi["E_AI_RECOMMEND_FAILED"] = "Dịch vụ gợi ý màu tạm thời không khả dụng, vui lòng thử lại sau"
	/*device*/
	languageMapVi["E_NO_ID"] = "Vui lòng nhập ID"
	languageMapVi["E_NO_MAC"] = "Vui lòng nhập địa chỉ MAC"
	languageMapVi["E_NO_INFO"] = "Không tồn tại bản ghi này"
	languageMapVi["E_NO_DEVICE"] = "Không tồn tại thiết bị này"
	languageMapVi["E_INVALID_MAC"] = "Vui lòng nhập địa chỉ MAC đúng định dạng"
	/*suggest*/
	languageMapVi["E_NO_CONTENT"] = "Vui lòng nhập nội dung"
	languageMapVi["E_NO_SUGGEST"] = "Không tồn tại đề xuất này"
	/*word*/
	languageMapVi["E_NO_WORD"] = "Không tồn tại từ nhạy cảm này"
	languageMapVi["E_WORD_EXIST"] = "Từ nhạy cảm đã tồn tại"
	/*follow*/
	languageMapVi["E_FOLLOWED"] = "Đã theo dõi"
	languageMapVi["E_UNFOLLOWED"] = "Đã bỏ theo dõi"
	languageMapVi["E_NOT_FOLLOWING"] = "Chưa theo dõi"
	languageMapVi["E_NO_SELF_FOLLOW"] = "Không thể theo dõi chính mình"
	/*circle*/
	languageMapVi["E_NO_CIRCLE_POST"] = "Không tồn tại bài viết này"
	languageMapVi["E_INVALID_REVIEW_STATUS"] = "Trạng thái duyệt không hợp lệ"
	languageMapVi["E_SENSITIVEWORD"] = "Bài viết của bạn chứa thông tin nhạy cảm. Để duy trì môi trường giao tiếp thân thiện, vui lòng điều chỉnh cách diễn đạt trước khi đăng"
	/*topic*/
	languageMapVi["E_NO_TOPIC"] = "Không tồn tại chủ đề này"
	languageMapVi["E_TOPIC_EXIST"] = "Chủ đề đã tồn tại"
	/*activity*/
	languageMapVi["E_NO_ACTIVITY"] = "Không tồn tại hoạt động này"
	languageMapVi["E_NO_TITLE"] = "Vui lòng nhập tiêu đề"
	languageMapVi["E_NO_TIME"] = "Vui lòng nhập thời gian"
	languageMapVi["E_ALREADY_JOINED"] = "Đã tham gia hoạt động này"
	languageMapVi["E_NOT_JOINED"] = "Chưa tham gia hoạt động này"
	languageMapVi["E_ALREADY_VOTED"] = "Đã bình chọn"
	languageMapVi["E_NOT_VOTED"] = "Chưa bình chọn"
	/*comment*/
	languageMapVi["E_NO_COMMENT"] = "Không tồn tại bình luận này"
	/*like*/
	languageMapVi["E_NO_LIKE"] = "Không tồn tại lượt thích này"
	/*collect*/
	languageMapVi["E_NO_COLLECT"] = "Chưa lưu"
	/*ota*/
	languageMapVi["E_NO_VERSION"] = "Vui lòng nhập phiên bản"
	/*voice*/
	languageMapVi["E_NO_VOICE_TEXT"] = "Vui lòng nhập lệnh thoại"
	languageMapVi["E_NO_COLOR_NAME"] = "Không nhận diện được tên màu"
	languageMapVi["V_UNKNOWN_COMMAND"] = "Lệnh thoại không xác định"
	languageMapVi["V_ACTION_HELP"] = "Không nhận diện được lệnh. Hãy thử:\n· Đổi màu: đổi màu đầu tiên thành xanh nửa đêm"
	/*checkin*/
	languageMapVi["E_ALREADY_CHECKIN"] = "Hôm nay đã điểm danh"
	/*anti_fake*/
	languageMapVi["E_INVALID_PRODUCT_MODEL"] = "Định dạng mẫu sản phẩm không đúng (cần 2 ký tự chữ và số)"
	languageMapVi["E_INVALID_BATCH"] = "Định dạng lô sản xuất không đúng (cần 6 ký tự chữ và số)"
	languageMapVi["E_INVALID_NUMBER"] = "Định dạng số thứ tự không đúng (cần 2 ký tự chữ và số)"
	languageMapVi["E_INVALID_ANTI_FAKE_CODE"] = "Định dạng mã chống hàng giả không đúng (cần 24 ký tự chữ và số)"
	languageMapVi["E_GENERATE_CODE_FAILED"] = "Tạo mã chống hàng giả thất bại"
	languageMapVi["E_ANTI_FAKE_CODE_ACTIVATED"] = "Mã chống hàng giả đã kích hoạt"
	languageMapVi["E_NO_ANTI_FAKE_CODE"] = "Không tồn tại mã chống hàng giả"
	languageMapVi["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "Mã chống hàng giả này đã được người dùng khác xác minh"
	languageMapVi["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "Mã chống hàng giả chưa kích hoạt, không thể xác minh"
	/*mail*/
	languageMapVi["MAIL_VERIFY_SUBJECT"] = "Mã xác minh TintaShift"
	languageMapVi["MAIL_VERIFY_TITLE"] = "Chào mừng bạn đến với TintaShift"
	languageMapVi["MAIL_VERIFY_DESC"] = "Xin chào, bạn đang thực hiện xác minh bảo mật tài khoản. Mã của bạn là:"
	languageMapVi["MAIL_VERIFY_SECURITY_TITLE"] = "Mẹo bảo mật: "
	languageMapVi["MAIL_VERIFY_SECURITY_DESC"] = "Nhân viên TintaShift sẽ không bao giờ yêu cầu mã này. Không chia sẻ mã với bất kỳ ai."
	languageMapVi["MAIL_VERIFY_IGNORE"] = "Nếu bạn không yêu cầu mã này, vui lòng bỏ qua email này."
	languageMapVi["MAIL_COPYRIGHT"] = "© 2026 TintaShift. Bảo lưu mọi quyền."
	/*social login*/
	languageMapVi["E_NO_GOOGLE_ACCESS_TOKEN"] = "Cần có access token Google"
	languageMapVi["E_GOOGLE_NOT_CONFIGURED"] = "Đăng nhập Google chưa được cấu hình trên máy chủ"
	languageMapVi["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Access token Google không hợp lệ hoặc đã hết hạn"
	languageMapVi["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "Email Google chưa được xác minh và không thể dùng để đăng nhập"
	languageMapVi["E_GOOGLE_EMAIL_BOUND_OTHER"] = "Email này đã được liên kết với một tài khoản Google khác"
	languageMapVi["E_NO_APPLE_ID_TOKEN"] = "Cần có identity token Apple"
	languageMapVi["E_APPLE_NOT_CONFIGURED"] = "Đăng nhập Apple chưa được cấu hình trên máy chủ"
	languageMapVi["E_APPLE_ID_TOKEN_INVALID"] = "Identity token Apple không hợp lệ hoặc đã hết hạn"
	languageMapVi["E_APPLE_EMAIL_NOT_VERIFIED"] = "Email Apple chưa được xác minh và không thể dùng để đăng nhập"
	languageMapVi["E_APPLE_EMAIL_BOUND_OTHER"] = "Email này đã được liên kết với một tài khoản Apple khác"
	languageMapVi["E_SOCIAL_NOT_CONFIGURED"] = "Khóa bí mật ký đăng nhập mạng xã hội chưa được cấu hình trên máy chủ"
	languageMapVi["E_NO_PROVIDER"] = "cần có provider"
	languageMapVi["E_SOCIAL_PROVIDER_INVALID"] = "provider đăng nhập mạng xã hội không được hỗ trợ"
	languageMapVi["E_NO_OPENID"] = "cần có openid"
	languageMapVi["E_NO_NONCE"] = "cần có nonce"
	languageMapVi["E_NO_TIMESTAMP"] = "cần có timestamp"
	languageMapVi["E_NO_SIGNATURE"] = "cần có signature"
	languageMapVi["E_NO_DEVICE_ID"] = "cần có device_id"
	languageMapVi["E_REQUEST_EXPIRED"] = "yêu cầu đã hết hạn, vui lòng thử lại"
	languageMapVi["E_CHALLENGE_INVALID"] = "challenge không hợp lệ"
	languageMapVi["E_CHALLENGE_EXPIRED"] = "challenge đã hết hạn"
	languageMapVi["E_CHALLENGE_USED"] = "challenge đã được sử dụng"
	languageMapVi["E_SIGNATURE_INVALID"] = "xác minh chữ ký thất bại"
	languageMapVi["E_SOCIAL_EMAIL_BOUND_OTHER"] = "Email này đã được liên kết với một tài khoản mạng xã hội khác"
	languageMapVi["E_INVALID_VIDEO"] = "Vui lòng tải lên video hợp lệ (MP4/MOV)"
	registerLocale("vi-VN", languageMapVi)
}
