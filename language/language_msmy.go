package language

var languageMapMs map[string]string

/*初始化*/
func initLanguageMsmy() {
	languageMapMs = make(map[string]string)
	/*user*/
	languageMapMs["E_NO_IV"] = "Sila masukkan vektor permulaan"
	languageMapMs["E_NO_CODE"] = "Sila masukkan kod pengesahan"
	languageMapMs["E_NO_DATA"] = "Sila masukkan data"
	languageMapMs["E_NO_PHONE"] = "Sila masukkan nombor telefon"
	languageMapMs["E_NO_PWD"] = "Sila masukkan kata laluan"
	languageMapMs["E_INVALID_USER"] = "Pengguna tidak dijumpai"
	languageMapMs["E_INVALID_PWD"] = "Kata laluan salah"
	languageMapMs["E_NO_NAME"] = "Sila masukkan nama"
	languageMapMs["E_USER_EXIST"] = "Akaun sudah wujud"
	languageMapMs["E_NO_TOKEN"] = "Sila log masuk terlebih dahulu"
	languageMapMs["E_NO_AUTH"] = "Tiada kebenaran"
	languageMapMs["E_NO_AVATAR"] = "Avatar tidak dijumpai"
	languageMapMs["E_NO_EMAIL"] = "Sila masukkan e-mel"
	languageMapMs["E_NO_CERT"] = "Sila masukkan kod pengesahan"
	languageMapMs["E_INVALID_EMAIL"] = "Format e-mel tidak sah"
	languageMapMs["E_DATA_TOO_SHORT"] = "Data yang disulitkan terlalu pendek"
	languageMapMs["E_INVALID_IV"] = "Vektor permulaan tidak sah"
	languageMapMs["E_INVALID_SESSION_KEY"] = "Kunci sesi tidak sah"
	languageMapMs["E_DECRYPT_FAIL"] = "Penguraian gagal, pastikan kod dan nombor telefon daripada sesi yang sama"
	languageMapMs["E_INVALID_PHONE"] = "Format nombor telefon tidak sah"
	languageMapMs["VERIFICATION_CODE"] = "kod pengesahan"
	languageMapMs["DO_NOT_TELL"] = "Jangan beritahu sesiapa"
	languageMapMs["E_CERT_ERR"] = "Kod pengesahan salah"
	languageMapMs["E_CERT_EXPIRED"] = "Kod pengesahan telah tamat tempoh. Sila minta yang baharu"
	languageMapMs["E_CERT_FIRST"] = "Sila minta kod pengesahan terlebih dahulu"
	languageMapMs["E_VERIFICATION_QUICKLY"] = "Anda meminta kod pengesahan terlalu kerap. Sila tunggu seketika"
	languageMapMs["E_ACCOUNT_LOCKED"] = "Terlalu banyak percubaan. Sila cuba lagi kemudian"
	languageMapMs["E_ACCOUNT_LOCKED_RETRY"] = "Terlalu banyak percubaan. Sila cuba lagi dalam %d saat"
	languageMapMs["E_RETRIES_REMAINING"] = ". Baki %d percubaan"
	languageMapMs["E_VERIFICATION_LIMIT"] = "Terlalu banyak permintaan kod pengesahan. Sila cuba lagi kemudian"
	/*common*/
	languageMapMs["E_TOO_LONG"] = "Panjang teks melebihi had"
	languageMapMs["E_INVALID_PARAM"] = "Parameter tidak sah"
	languageMapMs["E_FILE_NOT_FOUND"] = "Fail tidak dijumpai"
	/*color*/
	languageMapMs["E_INVALID_COLOR"] = "Warna tidak dijumpai"
	languageMapMs["E_INVALID_FILE"] = "Sila muat naik fail yang betul"
	languageMapMs["E_INVALID_PICTURE"] = "Sila muat naik gambar (JPG/PNG)"
	languageMapMs["E_AI_RECOMMEND_FAILED"] = "Perkhidmatan cadangan warna buat sementara tidak tersedia, sila cuba lagi kemudian"
	/*device*/
	languageMapMs["E_NO_ID"] = "Sila masukkan ID"
	languageMapMs["E_NO_MAC"] = "Sila masukkan alamat MAC"
	languageMapMs["E_NO_INFO"] = "Rekod tidak dijumpai"
	languageMapMs["E_NO_DEVICE"] = "Peranti tidak dijumpai"
	languageMapMs["E_INVALID_MAC"] = "Sila masukkan alamat MAC dalam format yang betul"
	/*suggest*/
	languageMapMs["E_NO_CONTENT"] = "Sila masukkan kandungan"
	languageMapMs["E_NO_SUGGEST"] = "Cadangan tidak dijumpai"
	/*word*/
	languageMapMs["E_NO_WORD"] = "Perkataan sensitif tidak dijumpai"
	languageMapMs["E_WORD_EXIST"] = "Perkataan sensitif sudah wujud"
	/*follow*/
	languageMapMs["E_FOLLOWED"] = "Sudah diikuti"
	languageMapMs["E_UNFOLLOWED"] = "Sudah tidak diikuti"
	languageMapMs["E_NOT_FOLLOWING"] = "Tidak diikuti"
	languageMapMs["E_NO_SELF_FOLLOW"] = "Tidak boleh mengikuti diri sendiri"
	/*circle*/
	languageMapMs["E_NO_CIRCLE_POST"] = "Catatan tidak dijumpai"
	languageMapMs["E_INVALID_REVIEW_STATUS"] = "Status semakan tidak sah"
	languageMapMs["E_SENSITIVEWORD"] = "Mesej anda mengandungi maklumat sensitif. Untuk mengekalkan persekitaran komunikasi yang mesra, sila laraskan ungkapan yang berkaitan sebelum menerbitkan"
	/*topic*/
	languageMapMs["E_NO_TOPIC"] = "Topik tidak dijumpai"
	languageMapMs["E_TOPIC_EXIST"] = "Topik sudah wujud"
	/*activity*/
	languageMapMs["E_NO_ACTIVITY"] = "Aktiviti tidak dijumpai"
	languageMapMs["E_NO_TITLE"] = "Sila masukkan tajuk"
	languageMapMs["E_NO_TIME"] = "Sila masukkan masa"
	languageMapMs["E_ALREADY_JOINED"] = "Sudah menyertai aktiviti ini"
	languageMapMs["E_NOT_JOINED"] = "Belum menyertai aktiviti ini"
	languageMapMs["E_ALREADY_VOTED"] = "Sudah mengundi"
	languageMapMs["E_NOT_VOTED"] = "Belum mengundi"
	/*comment*/
	languageMapMs["E_NO_COMMENT"] = "Komen tidak dijumpai"
	/*like*/
	languageMapMs["E_NO_LIKE"] = "Suka tidak dijumpai"
	/*collect*/
	languageMapMs["E_NO_COLLECT"] = "Belum disimpan"
	/*ota*/
	languageMapMs["E_NO_VERSION"] = "Sila masukkan versi"
	/*voice*/
	languageMapMs["E_NO_VOICE_TEXT"] = "Sila masukkan arahan suara"
	languageMapMs["E_NO_COLOR_NAME"] = "Nama warna tidak dikenali"
	languageMapMs["V_UNKNOWN_COMMAND"] = "Arahan suara tidak diketahui"
	languageMapMs["V_ACTION_HELP"] = "Tidak dapat mengenal arahan. Cuba:\n· Tukar warna: tukar yang pertama kepada biru tengah malam"
	/*checkin*/
	languageMapMs["E_ALREADY_CHECKIN"] = "Sudah mendaftar masuk hari ini"
	/*anti_fake*/
	languageMapMs["E_INVALID_PRODUCT_MODEL"] = "Format model produk tidak sah (perlu 2 aksara alphanumeric)"
	languageMapMs["E_INVALID_BATCH"] = "Format batch pengeluaran tidak sah (perlu 6 aksara alphanumeric)"
	languageMapMs["E_INVALID_NUMBER"] = "Format nombor tidak sah (perlu 2 aksara alphanumeric)"
	languageMapMs["E_INVALID_ANTI_FAKE_CODE"] = "Format kod anti-hijau tidak sah (perlu 24 aksara alphanumeric)"
	languageMapMs["E_GENERATE_CODE_FAILED"] = "Gagal menjana kod anti-hijau"
	languageMapMs["E_ANTI_FAKE_CODE_ACTIVATED"] = "Kod anti-hijau sudah diaktifkan"
	languageMapMs["E_NO_ANTI_FAKE_CODE"] = "Kod anti-hijau tidak dijumpai"
	languageMapMs["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "Kod anti-hijau ini telah disahkan oleh pengguna lain"
	languageMapMs["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "Kod anti-hijau tidak diaktifkan, tidak boleh disahkan"
	/*mail*/
	languageMapMs["MAIL_VERIFY_SUBJECT"] = "Kod Pengesahan TintaShift"
	languageMapMs["MAIL_VERIFY_TITLE"] = "Selamat datang ke TintaShift"
	languageMapMs["MAIL_VERIFY_DESC"] = "Helo, anda sedang melakukan pengesahan keselamatan akaun. Kod anda ialah:"
	languageMapMs["MAIL_VERIFY_SECURITY_TITLE"] = "Petua keselamatan: "
	languageMapMs["MAIL_VERIFY_SECURITY_DESC"] = "Kakitangan TintaShift tidak akan sekali-kali meminta kod ini. Jangan kongsi dengan sesiapa."
	languageMapMs["MAIL_VERIFY_IGNORE"] = "Jika anda tidak meminta kod ini, sila abaikan e-mel ini."
	languageMapMs["MAIL_COPYRIGHT"] = "© 2026 TintaShift. Hak cipta terpelihara."
	/*social login*/
	languageMapMs["E_NO_GOOGLE_ACCESS_TOKEN"] = "Token akses Google diperlukan"
	languageMapMs["E_GOOGLE_NOT_CONFIGURED"] = "Log masuk Google tidak dikonfigurasikan pada pelayan"
	languageMapMs["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Token akses Google tidak sah atau telah tamat tempoh"
	languageMapMs["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "E-mel Google tidak disahkan dan tidak boleh digunakan untuk log masuk"
	languageMapMs["E_GOOGLE_EMAIL_BOUND_OTHER"] = "E-mel ini sudah dipautkan ke akaun Google lain"
	languageMapMs["E_NO_APPLE_ID_TOKEN"] = "Token identiti Apple diperlukan"
	languageMapMs["E_APPLE_NOT_CONFIGURED"] = "Log masuk Apple tidak dikonfigurasikan pada pelayan"
	languageMapMs["E_APPLE_ID_TOKEN_INVALID"] = "Token identiti Apple tidak sah atau telah tamat tempoh"
	languageMapMs["E_APPLE_EMAIL_NOT_VERIFIED"] = "E-mel Apple tidak disahkan dan tidak boleh digunakan untuk log masuk"
	languageMapMs["E_APPLE_EMAIL_BOUND_OTHER"] = "E-mel ini sudah dipautkan ke akaun Apple lain"
	languageMapMs["E_SOCIAL_NOT_CONFIGURED"] = "Rahsia tandatangan log masuk sosial tidak dikonfigurasikan pada pelayan"
	languageMapMs["E_NO_PROVIDER"] = "provider diperlukan"
	languageMapMs["E_SOCIAL_PROVIDER_INVALID"] = "provider log masuk sosial tidak disokong"
	languageMapMs["E_NO_OPENID"] = "openid diperlukan"
	languageMapMs["E_NO_NONCE"] = "nonce diperlukan"
	languageMapMs["E_NO_TIMESTAMP"] = "timestamp diperlukan"
	languageMapMs["E_NO_SIGNATURE"] = "signature diperlukan"
	languageMapMs["E_NO_DEVICE_ID"] = "device_id diperlukan"
	languageMapMs["E_REQUEST_EXPIRED"] = "permintaan telah tamat tempoh, sila cuba lagi"
	languageMapMs["E_CHALLENGE_INVALID"] = "challenge tidak sah"
	languageMapMs["E_CHALLENGE_EXPIRED"] = "challenge telah tamat tempoh"
	languageMapMs["E_CHALLENGE_USED"] = "challenge telah digunakan"
	languageMapMs["E_SIGNATURE_INVALID"] = "pengesahan tandatangan gagal"
	languageMapMs["E_SOCIAL_EMAIL_BOUND_OTHER"] = "E-mel ini sudah dipautkan ke akaun sosial lain"
	languageMapMs["E_INVALID_VIDEO"] = "Sila muat naik video yang sah (MP4/MOV)"
	registerLocale("ms-MY", languageMapMs)
}
