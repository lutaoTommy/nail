package language

var languageMapTh map[string]string

/*初始化*/
func initLanguageThth() {
	languageMapTh = make(map[string]string)
	/*user*/
	languageMapTh["E_NO_IV"] = "กรุณาใส่เวกเตอร์เริ่มต้น"
	languageMapTh["E_NO_CODE"] = "กรุณาใส่รหัสการอนุญาต"
	languageMapTh["E_NO_DATA"] = "กรุณาใส่ข้อมูล"
	languageMapTh["E_NO_PHONE"] = "กรุณาใส่หมายเลขโทรศัพท์"
	languageMapTh["E_NO_PWD"] = "กรุณาใส่รหัสผ่าน"
	languageMapTh["E_INVALID_USER"] = "ไม่พบผู้ใช้นี้"
	languageMapTh["E_INVALID_PWD"] = "รหัสผ่านไม่ถูกต้อง"
	languageMapTh["E_NO_NAME"] = "กรุณาใส่ชื่อ"
	languageMapTh["E_USER_EXIST"] = "บัญชีมีอยู่แล้ว"
	languageMapTh["E_NO_TOKEN"] = "กรุณาเข้าสู่ระบบก่อน"
	languageMapTh["E_NO_AUTH"] = "ไม่มีสิทธิ์"
	languageMapTh["E_NO_AVATAR"] = "ไม่พบอวาตาร์นี้"
	languageMapTh["E_NO_EMAIL"] = "กรุณาใส่อีเมล"
	languageMapTh["E_NO_CERT"] = "กรุณาใส่รหัสยืนยัน"
	languageMapTh["E_INVALID_EMAIL"] = "รูปแบบอีเมลไม่ถูกต้อง"
	languageMapTh["E_DATA_TOO_SHORT"] = "ข้อมูลที่เข้ารหัสสั้นเกินไป"
	languageMapTh["E_INVALID_IV"] = "เวกเตอร์เริ่มต้นไม่ถูกต้อง"
	languageMapTh["E_INVALID_SESSION_KEY"] = "คีย์เซสชันไม่ถูกต้อง"
	languageMapTh["E_DECRYPT_FAIL"] = "การถอดรหัสล้มเหลว โปรดตรวจสอบว่ารหัสและหมายเลขโทรศัพท์มาจากการอนุญาตครั้งเดียวกัน"
	languageMapTh["E_INVALID_PHONE"] = "รูปแบบหมายเลขโทรศัพท์ไม่ถูกต้อง"
	languageMapTh["VERIFICATION_CODE"] = "รหัสยืนยัน"
	languageMapTh["DO_NOT_TELL"] = "โปรดอย่าบอกผู้อื่น"
	languageMapTh["E_CERT_ERR"] = "รหัสยืนยันไม่ถูกต้อง"
	languageMapTh["E_CERT_EXPIRED"] = "รหัสยืนยันหมดอายุแล้ว โปรดขอรหัสใหม่"
	languageMapTh["E_CERT_FIRST"] = "โปรดขอรหัสยืนยันก่อน"
	languageMapTh["E_VERIFICATION_QUICKLY"] = "คุณขอรหัสยืนยันบ่อยเกินไป โปรดรอสักครู่"
	languageMapTh["E_ACCOUNT_LOCKED"] = "พยายามมากเกินไป โปรดลองอีกครั้งภายหลัง"
	languageMapTh["E_ACCOUNT_LOCKED_RETRY"] = "พยายามมากเกินไป โปรดลองอีกครั้งใน %d วินาที"
	languageMapTh["E_RETRIES_REMAINING"] = " เหลือความพยายามอีก %d ครั้ง"
	languageMapTh["E_VERIFICATION_LIMIT"] = "ขอรหัสยืนยันบ่อยเกินไป โปรดลองอีกครั้งภายหลัง"
	/*common*/
	languageMapTh["E_TOO_LONG"] = "ความยาวอักขระเกินขีดจำกัด"
	languageMapTh["E_INVALID_PARAM"] = "พารามิเตอร์ไม่ถูกต้อง"
	languageMapTh["E_FILE_NOT_FOUND"] = "ไม่พบไฟล์"
	/*color*/
	languageMapTh["E_INVALID_COLOR"] = "ไม่พบสีนี้"
	languageMapTh["E_INVALID_FILE"] = "โปรดอัปโหลดไฟล์ที่ถูกต้อง"
	languageMapTh["E_INVALID_PICTURE"] = "โปรดอัปโหลดรูปภาพ (JPG/PNG)"
	languageMapTh["E_AI_RECOMMEND_FAILED"] = "บริการแนะนำสีชั่วคราวไม่พร้อมใช้งาน โปรดลองอีกครั้งภายหลัง"
	/*device*/
	languageMapTh["E_NO_ID"] = "กรุณาใส่ ID"
	languageMapTh["E_NO_MAC"] = "กรุณาใส่ที่อยู่ MAC"
	languageMapTh["E_NO_INFO"] = "ไม่พบบันทึกนี้"
	languageMapTh["E_NO_DEVICE"] = "ไม่พบอุปกรณ์นี้"
	languageMapTh["E_INVALID_MAC"] = "กรุณาใส่ที่อยู่ MAC ในรูปแบบที่ถูกต้อง"
	/*suggest*/
	languageMapTh["E_NO_CONTENT"] = "กรุณาใส่เนื้อหา"
	languageMapTh["E_NO_SUGGEST"] = "ไม่พบข้อเสนอนี้"
	/*word*/
	languageMapTh["E_NO_WORD"] = "ไม่พบคำละเอียดอ่อนนี้"
	languageMapTh["E_WORD_EXIST"] = "คำละเอียดอ่อนมีอยู่แล้ว"
	/*follow*/
	languageMapTh["E_FOLLOWED"] = "ติดตามแล้ว"
	languageMapTh["E_UNFOLLOWED"] = "เลิกติดตามแล้ว"
	languageMapTh["E_NOT_FOLLOWING"] = "ยังไม่ได้ติดตาม"
	languageMapTh["E_NO_SELF_FOLLOW"] = "ไม่สามารถติดตามตัวเองได้"
	/*circle*/
	languageMapTh["E_NO_CIRCLE_POST"] = "ไม่พบโพสต์นี้"
	languageMapTh["E_INVALID_REVIEW_STATUS"] = "สถานะการตรวจสอบไม่ถูกต้อง"
	languageMapTh["E_SENSITIVEWORD"] = "ข้อความของคุณมีข้อมูลอ่อนไหว เพื่อรักษาบรรยากาศการสื่อสารที่เป็นมิตร โปรดปรับถ้อยคำที่เกี่ยวข้องก่อนเผยแพร่"
	/*topic*/
	languageMapTh["E_NO_TOPIC"] = "ไม่พบหัวข้อนี้"
	languageMapTh["E_TOPIC_EXIST"] = "หัวข้อมีอยู่แล้ว"
	/*activity*/
	languageMapTh["E_NO_ACTIVITY"] = "ไม่พบกิจกรรมนี้"
	languageMapTh["E_NO_TITLE"] = "กรุณาใส่ชื่อเรื่อง"
	languageMapTh["E_NO_TIME"] = "กรุณาใส่เวลา"
	languageMapTh["E_ALREADY_JOINED"] = "เข้าร่วมกิจกรรมนี้แล้ว"
	languageMapTh["E_NOT_JOINED"] = "ยังไม่ได้เข้าร่วมกิจกรรมนี้"
	languageMapTh["E_ALREADY_VOTED"] = "ลงคะแนนแล้ว"
	languageMapTh["E_NOT_VOTED"] = "ยังไม่ได้ลงคะแนน"
	/*comment*/
	languageMapTh["E_NO_COMMENT"] = "ไม่พบความคิดเห็นนี้"
	/*like*/
	languageMapTh["E_NO_LIKE"] = "ไม่พบการถูกใจนี้"
	/*collect*/
	languageMapTh["E_NO_COLLECT"] = "ยังไม่ได้บันทึก"
	/*ota*/
	languageMapTh["E_NO_VERSION"] = "กรุณาใส่เวอร์ชัน"
	/*voice*/
	languageMapTh["E_NO_VOICE_TEXT"] = "กรุณาใส่คำสั่งเสียง"
	languageMapTh["E_NO_COLOR_NAME"] = "ไม่พบชื่อสี"
	languageMapTh["V_UNKNOWN_COMMAND"] = "คำสั่งเสียงไม่รู้จัก"
	languageMapTh["V_ACTION_HELP"] = "ไม่สามารถจำคำสั่งได้ โปรดลอง:\n· เปลี่ยนสี: เปลี่ยนสีแรกเป็นสีน้ำเงินเที่ยงคืน"
	/*checkin*/
	languageMapTh["E_ALREADY_CHECKIN"] = "เช็คอินวันนี้แล้ว"
	/*anti_fake*/
	languageMapTh["E_INVALID_PRODUCT_MODEL"] = "รูปแบบรุ่นสินค้าไม่ถูกต้อง (ต้องใช้อักขระตัวอักษรและตัวเลข 2 ตัว)"
	languageMapTh["E_INVALID_BATCH"] = "รูปแบบล็อตการผลิตไม่ถูกต้อง (ต้องใช้อักขระตัวอักษรและตัวเลข 6 ตัว)"
	languageMapTh["E_INVALID_NUMBER"] = "รูปแบบหมายเลขไม่ถูกต้อง (ต้องใช้อักขระตัวอักษรและตัวเลข 2 ตัว)"
	languageMapTh["E_INVALID_ANTI_FAKE_CODE"] = "รูปแบบรหัสป้องกันการปลอมแปลงไม่ถูกต้อง (ต้องใช้อักขระตัวอักษรและตัวเลข 24 ตัว)"
	languageMapTh["E_GENERATE_CODE_FAILED"] = "สร้างรหัสป้องกันการปลอมแปลงล้มเหลว"
	languageMapTh["E_ANTI_FAKE_CODE_ACTIVATED"] = "รหัสป้องกันการปลอมแปลงเปิดใช้งานแล้ว"
	languageMapTh["E_NO_ANTI_FAKE_CODE"] = "ไม่พบรหัสป้องกันการปลอมแปลง"
	languageMapTh["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "รหัสป้องกันการปลอมแปลงนี้ถูกยืนยันโดยผู้ใช้อื่นแล้ว"
	languageMapTh["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "รหัสป้องกันการปลอมแปลงยังไม่เปิดใช้งาน ไม่สามารถยืนยันได้"
	/*mail*/
	languageMapTh["MAIL_VERIFY_SUBJECT"] = "รหัสยืนยัน TintaShift"
	languageMapTh["MAIL_VERIFY_TITLE"] = "ยินดีต้อนรับสู่ TintaShift"
	languageMapTh["MAIL_VERIFY_DESC"] = "สวัสดี คุณกำลังดำเนินการยืนยันความปลอดภัยของบัญชี รหัสของคุณคือ:"
	languageMapTh["MAIL_VERIFY_SECURITY_TITLE"] = "คำแนะนำด้านความปลอดภัย: "
	languageMapTh["MAIL_VERIFY_SECURITY_DESC"] = "เจ้าหน้าที่ TintaShift จะไม่ขอรหัสนี้จากคุณ โปรดอย่าแชร์ให้ใครก็ตาม"
	languageMapTh["MAIL_VERIFY_IGNORE"] = "หากคุณไม่ได้ขอรหัสนี้ โปรดเพิกเฉยต่ออีเมลนี้"
	languageMapTh["MAIL_COPYRIGHT"] = "© 2026 TintaShift สงวนลิขสิทธิ์"
	/*social login*/
	languageMapTh["E_NO_GOOGLE_ACCESS_TOKEN"] = "ต้องใช้โทเค็นการเข้าถึงของ Google"
	languageMapTh["E_GOOGLE_NOT_CONFIGURED"] = "ยังไม่ได้กำหนดค่าการเข้าสู่ระบบด้วย Google บนเซิร์ฟเวอร์"
	languageMapTh["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "โทเค็นการเข้าถึงของ Google ไม่ถูกต้องหรือหมดอายุ"
	languageMapTh["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "อีเมล Google ยังไม่ได้รับการยืนยัน จึงไม่สามารถใช้เข้าสู่ระบบได้"
	languageMapTh["E_GOOGLE_EMAIL_BOUND_OTHER"] = "อีเมลนี้เชื่อมโยงกับบัญชี Google อื่นอยู่แล้ว"
	languageMapTh["E_NO_APPLE_ID_TOKEN"] = "ต้องใช้ identity token ของ Apple"
	languageMapTh["E_APPLE_NOT_CONFIGURED"] = "ยังไม่ได้กำหนดค่าการเข้าสู่ระบบด้วย Apple บนเซิร์ฟเวอร์"
	languageMapTh["E_APPLE_ID_TOKEN_INVALID"] = "identity token ของ Apple ไม่ถูกต้องหรือหมดอายุ"
	languageMapTh["E_APPLE_EMAIL_NOT_VERIFIED"] = "อีเมล Apple ยังไม่ได้รับการยืนยัน จึงไม่สามารถใช้เข้าสู่ระบบได้"
	languageMapTh["E_APPLE_EMAIL_BOUND_OTHER"] = "อีเมลนี้เชื่อมโยงกับบัญชี Apple อื่นอยู่แล้ว"
	languageMapTh["E_SOCIAL_NOT_CONFIGURED"] = "ยังไม่ได้กำหนดค่ารหัสลับสำหรับลายเซ็นการเข้าสู่ระบบโซเชียลบนเซิร์ฟเวอร์"
	languageMapTh["E_NO_PROVIDER"] = "ต้องระบุ provider"
	languageMapTh["E_SOCIAL_PROVIDER_INVALID"] = "ไม่รองรับ provider การเข้าสู่ระบบโซเชียลนี้"
	languageMapTh["E_NO_OPENID"] = "ต้องระบุ openid"
	languageMapTh["E_NO_NONCE"] = "ต้องระบุ nonce"
	languageMapTh["E_NO_TIMESTAMP"] = "ต้องระบุ timestamp"
	languageMapTh["E_NO_SIGNATURE"] = "ต้องระบุ signature"
	languageMapTh["E_NO_DEVICE_ID"] = "ต้องระบุ device_id"
	languageMapTh["E_REQUEST_EXPIRED"] = "คำขอหมดอายุแล้ว โปรดลองอีกครั้ง"
	languageMapTh["E_CHALLENGE_INVALID"] = "challenge ไม่ถูกต้อง"
	languageMapTh["E_CHALLENGE_EXPIRED"] = "challenge หมดอายุแล้ว"
	languageMapTh["E_CHALLENGE_USED"] = "challenge ถูกใช้ไปแล้ว"
	languageMapTh["E_SIGNATURE_INVALID"] = "การตรวจสอบลายเซ็นล้มเหลว"
	languageMapTh["E_SOCIAL_EMAIL_BOUND_OTHER"] = "อีเมลนี้เชื่อมโยงกับบัญชีโซเชียลอื่นอยู่แล้ว"
	languageMapTh["E_INVALID_VIDEO"] = "โปรดอัปโหลดวิดีโอที่ถูกต้อง (MP4/MOV)"
	registerLocale("th-TH", languageMapTh)
}
