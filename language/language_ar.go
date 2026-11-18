package language

var languageMapAr map[string]string

/*初始化*/
func initLanguageAr() {
	languageMapAr = make(map[string]string)
	/*user*/
	languageMapAr["E_NO_IV"] = "يرجى إدخال متجه التهيئة"
	languageMapAr["E_NO_CODE"] = "يرجى إدخال رمز التفويض"
	languageMapAr["E_NO_DATA"] = "يرجى إدخال البيانات"
	languageMapAr["E_NO_PHONE"] = "يرجى إدخال رقم الهاتف"
	languageMapAr["E_NO_PWD"] = "يرجى إدخال كلمة المرور"
	languageMapAr["E_INVALID_USER"] = "لا يوجد مستخدم بهذا الاسم"
	languageMapAr["E_INVALID_PWD"] = "كلمة المرور خاطئة"
	languageMapAr["E_NO_NAME"] = "يرجى إدخال الاسم"
	languageMapAr["E_USER_EXIST"] = "الحساب موجود بالفعل"
	languageMapAr["E_NO_TOKEN"] = "يرجى تسجيل الدخول أولاً"
	languageMapAr["E_NO_AUTH"] = "لا توجد صلاحيات"
	languageMapAr["E_NO_AVATAR"] = "لا يوجد صورة رمزية بهذا الاسم"
	languageMapAr["E_NO_EMAIL"] = "يرجى إدخال البريد الإلكتروني"
	languageMapAr["E_NO_CERT"] = "يرجى إدخال رمز التحقق"
	languageMapAr["E_INVALID_EMAIL"] = "صيغة البريد الإلكتروني غير صحيحة"
	languageMapAr["E_DATA_TOO_SHORT"] = "البيانات المشفرة قصيرة جداً"
	languageMapAr["E_INVALID_IV"] = "متجه التهيئة غير صالح"
	languageMapAr["E_INVALID_SESSION_KEY"] = "مفتاح الجلسة غير صالح"
	languageMapAr["E_DECRYPT_FAIL"] = "فشل فك التشفير، تأكد من أن الرمز ورقم الهاتف من نفس الجلسة"
	languageMapAr["E_INVALID_PHONE"] = "صيغة رقم الهاتف غير صحيحة"
	languageMapAr["VERIFICATION_CODE"] = "رمز التحقق"
	languageMapAr["DO_NOT_TELL"] = "لا تخبر أحداً"
	languageMapAr["E_CERT_ERR"] = "رمز التحقق خاطئ"
	languageMapAr["E_CERT_EXPIRED"] = "انتهت صلاحية رمز التحقق. يرجى طلب رمز جديد"
	languageMapAr["E_CERT_FIRST"] = "يرجى طلب رمز التحقق أولاً"
	languageMapAr["E_VERIFICATION_QUICKLY"] = "أنت تطلب رموز تحقق بشكل متكرر جداً. يرجى الانتظار قليلاً"
	languageMapAr["E_ACCOUNT_LOCKED"] = "محاولات كثيرة جداً. يرجى المحاولة مرة أخرى لاحقاً"
	languageMapAr["E_ACCOUNT_LOCKED_RETRY"] = "محاولات كثيرة جداً. يرجى المحاولة مرة أخرى بعد %d ثانية"
	languageMapAr["E_RETRIES_REMAINING"] = ". يتبقى %d محاولة"
	languageMapAr["E_VERIFICATION_LIMIT"] = "طلبات رموز تحقق كثيرة جداً. يرجى المحاولة مرة أخرى لاحقاً"
	/*common*/
	languageMapAr["E_TOO_LONG"] = "طول النص يتجاوز الحد المسموح"
	languageMapAr["E_INVALID_PARAM"] = "معامل غير صالح"
	languageMapAr["E_FILE_NOT_FOUND"] = "الملف غير موجود"
	/*color*/
	languageMapAr["E_INVALID_COLOR"] = "لا يوجد لون بهذا الاسم"
	languageMapAr["E_INVALID_FILE"] = "يرجى رفع ملف صحيح"
	languageMapAr["E_INVALID_PICTURE"] = "يرجى رفع صورة (JPG/PNG)"
	languageMapAr["E_AI_RECOMMEND_FAILED"] = "خدمة توصية الألوان غير متاحة مؤقتاً، يرجى المحاولة مرة أخرى لاحقاً"
	/*device*/
	languageMapAr["E_NO_ID"] = "يرجى إدخال المعرّف"
	languageMapAr["E_NO_MAC"] = "يرجى إدخال عنوان MAC"
	languageMapAr["E_NO_INFO"] = "لا يوجد سجل بهذا الاسم"
	languageMapAr["E_NO_DEVICE"] = "لا يوجد جهاز بهذا الاسم"
	languageMapAr["E_INVALID_MAC"] = "يرجى إدخال عنوان MAC بالصيغة الصحيحة"
	/*suggest*/
	languageMapAr["E_NO_CONTENT"] = "يرجى إدخال المحتوى"
	languageMapAr["E_NO_SUGGEST"] = "لا يوجد اقتراح بهذا الاسم"
	/*word*/
	languageMapAr["E_NO_WORD"] = "لا يوجد كلمة حساسة بهذا الاسم"
	languageMapAr["E_WORD_EXIST"] = "الكلمة الحساسة موجودة بالفعل"
	/*follow*/
	languageMapAr["E_FOLLOWED"] = "تمت المتابعة بالفعل"
	languageMapAr["E_UNFOLLOWED"] = "تم إلغاء المتابعة بالفعل"
	languageMapAr["E_NOT_FOLLOWING"] = "غير متابع"
	languageMapAr["E_NO_SELF_FOLLOW"] = "لا يمكن متابعة نفسك"
	/*circle*/
	languageMapAr["E_NO_CIRCLE_POST"] = "لا يوجد منشور بهذا الاسم"
	languageMapAr["E_INVALID_REVIEW_STATUS"] = "حالة المراجعة غير صالحة"
	languageMapAr["E_SENSITIVEWORD"] = "رسالتك تحتوي على معلومات حساسة. للحفاظ على بيئة تواصل ودية، يرجى تعديل التعابير ذات الصلة قبل النشر"
	/*topic*/
	languageMapAr["E_NO_TOPIC"] = "لا يوجد موضوع بهذا الاسم"
	languageMapAr["E_TOPIC_EXIST"] = "الموضوع موجود بالفعل"
	/*activity*/
	languageMapAr["E_NO_ACTIVITY"] = "لا يوجد نشاط بهذا الاسم"
	languageMapAr["E_NO_TITLE"] = "يرجى إدخال العنوان"
	languageMapAr["E_NO_TIME"] = "يرجى إدخال الوقت"
	languageMapAr["E_ALREADY_JOINED"] = "لقد انضممت إلى هذا النشاط بالفعل"
	languageMapAr["E_NOT_JOINED"] = "لم تنضم إلى هذا النشاط"
	languageMapAr["E_ALREADY_VOTED"] = "لقد تصويت بالفعل"
	languageMapAr["E_NOT_VOTED"] = "لم تصوت بعد"
	/*comment*/
	languageMapAr["E_NO_COMMENT"] = "لا يوجد تعليق بهذا الاسم"
	/*like*/
	languageMapAr["E_NO_LIKE"] = "لا يوجد إعجاب بهذا الاسم"
	/*collect*/
	languageMapAr["E_NO_COLLECT"] = "لم يتم الحفظ بعد"
	/*ota*/
	languageMapAr["E_NO_VERSION"] = "يرجى إدخال الإصدار"
	/*voice*/
	languageMapAr["E_NO_VOICE_TEXT"] = "يرجى إدخال أمر صوتي"
	languageMapAr["E_NO_COLOR_NAME"] = "لم يتم التعرف على اسم اللون"
	languageMapAr["V_UNKNOWN_COMMAND"] = "أمر صوتي غير معروف"
	languageMapAr["V_ACTION_HELP"] = "تعذّر التعرف على الأمر. جرّب:\n· تغيير اللون: غيّر الأول إلى الأزرق الليلي"
	/*checkin*/
	languageMapAr["E_ALREADY_CHECKIN"] = "لقد سجلت حضورك اليوم بالفعل"
	/*anti_fake*/
	languageMapAr["E_INVALID_PRODUCT_MODEL"] = "صيغة طراز المنتج غير صحيحة (يلزم حرفان أبجديان أو رقمان)"
	languageMapAr["E_INVALID_BATCH"] = "صيغة دفعة الإنتاج غير صحيحة (يلزم 6 أحرف أبجدية أو أرقام)"
	languageMapAr["E_INVALID_NUMBER"] = "صيغة الرقم غير صحيحة (يلزم حرفان أبجديان أو رقمان)"
	languageMapAr["E_INVALID_ANTI_FAKE_CODE"] = "صيغة رمز مكافحة التزييف غير صحيحة (يلزم 24 حرفاً أرقامياً)"
	languageMapAr["E_GENERATE_CODE_FAILED"] = "فشل إنشاء رمز مكافحة التزييف"
	languageMapAr["E_ANTI_FAKE_CODE_ACTIVATED"] = "تم تفعيل رمز مكافحة التزييف بالفعل"
	languageMapAr["E_NO_ANTI_FAKE_CODE"] = "لا يوجد رمز مكافحة التزييف"
	languageMapAr["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "تم التحقق من رمز مكافحة التزييف هذا بواسطة مستخدم آخر بالفعل"
	languageMapAr["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "رمز مكافحة التزييف غير مفعّل، لا يمكن التحقق منه"
	/*mail*/
	languageMapAr["MAIL_VERIFY_SUBJECT"] = "رمز التحقق من TintaShift"
	languageMapAr["MAIL_VERIFY_TITLE"] = "مرحباً بك في TintaShift"
	languageMapAr["MAIL_VERIFY_DESC"] = "مرحباً، أنت تقوم بإجراء التحقق من أمان الحساب. رمزك هو:"
	languageMapAr["MAIL_VERIFY_SECURITY_TITLE"] = "نصيحة أمنية: "
	languageMapAr["MAIL_VERIFY_SECURITY_DESC"] = "لن يطلب موظفو TintaShift هذا الرمز أبداً. لا تشاركه مع أي شخص."
	languageMapAr["MAIL_VERIFY_IGNORE"] = "إذا لم تطلب هذا الرمز، يرجى تجاهل هذا البريد الإلكتروني."
	languageMapAr["MAIL_COPYRIGHT"] = "© 2026 TintaShift. جميع الحقوق محفوظة."
	/*social login*/
	languageMapAr["E_NO_GOOGLE_ACCESS_TOKEN"] = "مطلوب رمز وصول Google"
	languageMapAr["E_GOOGLE_NOT_CONFIGURED"] = "لم يتم تكوين تسجيل الدخول عبر Google على الخادم"
	languageMapAr["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "رمز وصول Google غير صالح أو منتهي الصلاحية"
	languageMapAr["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "البريد الإلكتروني لـ Google غير موثّق ولا يمكن استخدامه لتسجيل الدخول"
	languageMapAr["E_GOOGLE_EMAIL_BOUND_OTHER"] = "هذا البريد الإلكتروني مرتبط بالفعل بحساب Google آخر"
	languageMapAr["E_NO_APPLE_ID_TOKEN"] = "مطلوب رمز هوية Apple"
	languageMapAr["E_APPLE_NOT_CONFIGURED"] = "لم يتم تكوين تسجيل الدخول عبر Apple على الخادم"
	languageMapAr["E_APPLE_ID_TOKEN_INVALID"] = "رمز هوية Apple غير صالح أو منتهي الصلاحية"
	languageMapAr["E_APPLE_EMAIL_NOT_VERIFIED"] = "البريد الإلكتروني لـ Apple غير موثّق ولا يمكن استخدامه لتسجيل الدخول"
	languageMapAr["E_APPLE_EMAIL_BOUND_OTHER"] = "هذا البريد الإلكتروني مرتبط بالفعل بحساب Apple آخر"
	languageMapAr["E_SOCIAL_NOT_CONFIGURED"] = "لم يتم تكوين سر توقيع تسجيل الدخول الاجتماعي على الخادم"
	languageMapAr["E_NO_PROVIDER"] = "provider مطلوب"
	languageMapAr["E_SOCIAL_PROVIDER_INVALID"] = "provider تسجيل الدخول الاجتماعي غير مدعوم"
	languageMapAr["E_NO_OPENID"] = "openid مطلوب"
	languageMapAr["E_NO_NONCE"] = "nonce مطلوب"
	languageMapAr["E_NO_TIMESTAMP"] = "timestamp مطلوب"
	languageMapAr["E_NO_SIGNATURE"] = "signature مطلوب"
	languageMapAr["E_NO_DEVICE_ID"] = "device_id مطلوب"
	languageMapAr["E_REQUEST_EXPIRED"] = "انتهت صلاحية الطلب، يرجى إعادة المحاولة"
	languageMapAr["E_CHALLENGE_INVALID"] = "challenge غير صالح"
	languageMapAr["E_CHALLENGE_EXPIRED"] = "انتهت صلاحية challenge"
	languageMapAr["E_CHALLENGE_USED"] = "تم استخدام challenge بالفعل"
	languageMapAr["E_SIGNATURE_INVALID"] = "فشل التحقق من التوقيع"
	languageMapAr["E_SOCIAL_EMAIL_BOUND_OTHER"] = "هذا البريد الإلكتروني مرتبط بالفعل بحساب اجتماعي آخر"
	languageMapAr["E_INVALID_VIDEO"] = "يرجى رفع فيديو صحيح (MP4/MOV)"
	registerLocale("ar", languageMapAr)
}
