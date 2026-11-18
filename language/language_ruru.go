package language

var languageMapRu map[string]string

/*初始化*/
func initLanguageRuru() {
	languageMapRu = make(map[string]string)
	/*user*/
	languageMapRu["E_NO_IV"] = "Введите вектор инициализации"
	languageMapRu["E_NO_CODE"] = "Введите код авторизации"
	languageMapRu["E_NO_DATA"] = "Введите данные"
	languageMapRu["E_NO_PHONE"] = "Введите номер телефона"
	languageMapRu["E_NO_PWD"] = "Введите пароль"
	languageMapRu["E_INVALID_USER"] = "Пользователь не найден"
	languageMapRu["E_INVALID_PWD"] = "Неверный пароль"
	languageMapRu["E_NO_NAME"] = "Введите имя"
	languageMapRu["E_USER_EXIST"] = "Аккаунт уже существует"
	languageMapRu["E_NO_TOKEN"] = "Сначала войдите в систему"
	languageMapRu["E_NO_AUTH"] = "Нет прав доступа"
	languageMapRu["E_NO_AVATAR"] = "Аватар не найден"
	languageMapRu["E_NO_EMAIL"] = "Введите адрес электронной почты"
	languageMapRu["E_NO_CERT"] = "Введите код подтверждения"
	languageMapRu["E_INVALID_EMAIL"] = "Неверный формат адреса электронной почты"
	languageMapRu["E_DATA_TOO_SHORT"] = "Шифрованные данные слишком короткие"
	languageMapRu["E_INVALID_IV"] = "Недействительный вектор инициализации"
	languageMapRu["E_INVALID_SESSION_KEY"] = "Недействительный ключ сессии"
	languageMapRu["E_DECRYPT_FAIL"] = "Не удалось расшифровать; убедитесь, что код и номер телефона из одной сессии"
	languageMapRu["E_INVALID_PHONE"] = "Неверный формат номера телефона"
	languageMapRu["VERIFICATION_CODE"] = "код подтверждения"
	languageMapRu["DO_NOT_TELL"] = "Не сообщайте его другим"
	languageMapRu["E_CERT_ERR"] = "Неверный код подтверждения"
	languageMapRu["E_CERT_EXPIRED"] = "Срок действия кода подтверждения истёк. Запросите новый"
	languageMapRu["E_CERT_FIRST"] = "Сначала запросите код подтверждения"
	languageMapRu["E_VERIFICATION_QUICKLY"] = "Вы слишком часто запрашиваете коды подтверждения. Подождите немного"
	languageMapRu["E_ACCOUNT_LOCKED"] = "Слишком много попыток. Повторите попытку позже"
	languageMapRu["E_ACCOUNT_LOCKED_RETRY"] = "Слишком много попыток. Повторите попытку через %d сек."
	languageMapRu["E_RETRIES_REMAINING"] = ". Осталось попыток: %d"
	languageMapRu["E_VERIFICATION_LIMIT"] = "Слишком много запросов кодов подтверждения. Повторите попытку позже"
	/*common*/
	languageMapRu["E_TOO_LONG"] = "Длина текста превышает лимит"
	languageMapRu["E_INVALID_PARAM"] = "Недействительный параметр"
	languageMapRu["E_FILE_NOT_FOUND"] = "Файл не найден"
	/*color*/
	languageMapRu["E_INVALID_COLOR"] = "Цвет не найден"
	languageMapRu["E_INVALID_FILE"] = "Загрузите корректный файл"
	languageMapRu["E_INVALID_PICTURE"] = "Загрузите изображение (JPG/PNG)"
	languageMapRu["E_AI_RECOMMEND_FAILED"] = "Сервис рекомендации цветов временно недоступен, повторите попытку позже"
	/*device*/
	languageMapRu["E_NO_ID"] = "Введите ID"
	languageMapRu["E_NO_MAC"] = "Введите MAC-адрес"
	languageMapRu["E_NO_INFO"] = "Запись не найдена"
	languageMapRu["E_NO_DEVICE"] = "Устройство не найдено"
	languageMapRu["E_INVALID_MAC"] = "Введите MAC-адрес в правильном формате"
	/*suggest*/
	languageMapRu["E_NO_CONTENT"] = "Введите содержание"
	languageMapRu["E_NO_SUGGEST"] = "Предложение не найдено"
	/*word*/
	languageMapRu["E_NO_WORD"] = "Слово-фильтр не найдено"
	languageMapRu["E_WORD_EXIST"] = "Слово-фильтр уже существует"
	/*follow*/
	languageMapRu["E_FOLLOWED"] = "Вы уже подписаны"
	languageMapRu["E_UNFOLLOWED"] = "Подписка отменена"
	languageMapRu["E_NOT_FOLLOWING"] = "Вы не подписаны"
	languageMapRu["E_NO_SELF_FOLLOW"] = "Нельзя подписаться на себя"
	/*circle*/
	languageMapRu["E_NO_CIRCLE_POST"] = "Публикация не найдена"
	languageMapRu["E_INVALID_REVIEW_STATUS"] = "Недействительный статус проверки"
	languageMapRu["E_SENSITIVEWORD"] = "Ваше сообщение содержит конфиденциальную информацию. Чтобы поддерживать дружелюбную атмосферу общения, отредактируйте соответствующие выражения перед публикацией"
	/*topic*/
	languageMapRu["E_NO_TOPIC"] = "Тема не найдена"
	languageMapRu["E_TOPIC_EXIST"] = "Тема уже существует"
	/*activity*/
	languageMapRu["E_NO_ACTIVITY"] = "Активность не найдена"
	languageMapRu["E_NO_TITLE"] = "Введите заголовок"
	languageMapRu["E_NO_TIME"] = "Введите время"
	languageMapRu["E_ALREADY_JOINED"] = "Вы уже участвуете в этом мероприятии"
	languageMapRu["E_NOT_JOINED"] = "Вы не участвуете в этом мероприятии"
	languageMapRu["E_ALREADY_VOTED"] = "Вы уже проголосовали"
	languageMapRu["E_NOT_VOTED"] = "Вы не голосовали"
	/*comment*/
	languageMapRu["E_NO_COMMENT"] = "Комментарий не найден"
	/*like*/
	languageMapRu["E_NO_LIKE"] = "Отметка «Нравится» не найдена"
	/*collect*/
	languageMapRu["E_NO_COLLECT"] = "Не сохранено"
	/*ota*/
	languageMapRu["E_NO_VERSION"] = "Введите версию"
	/*voice*/
	languageMapRu["E_NO_VOICE_TEXT"] = "Введите голосовую команду"
	languageMapRu["E_NO_COLOR_NAME"] = "Название цвета не распознано"
	languageMapRu["V_UNKNOWN_COMMAND"] = "Неизвестная голосовая команда"
	languageMapRu["V_ACTION_HELP"] = "Не удалось распознать команду. Попробуйте:\n· Сменить цвет: изменить первый на полуночно-синий"
	/*checkin*/
	languageMapRu["E_ALREADY_CHECKIN"] = "Вы уже отметились сегодня"
	/*anti_fake*/
	languageMapRu["E_INVALID_PRODUCT_MODEL"] = "Неверный формат модели продукта (требуется 2 буквенно-цифровых символа)"
	languageMapRu["E_INVALID_BATCH"] = "Неверный формат производственной партии (требуется 6 буквенно-цифровых символов)"
	languageMapRu["E_INVALID_NUMBER"] = "Неверный формат номера (требуется 2 буквенно-цифровых символа)"
	languageMapRu["E_INVALID_ANTI_FAKE_CODE"] = "Неверный формат антиконтрафактного кода (требуется 24 буквенно-цифровых символа)"
	languageMapRu["E_GENERATE_CODE_FAILED"] = "Не удалось сгенерировать антиконтрафактный код"
	languageMapRu["E_ANTI_FAKE_CODE_ACTIVATED"] = "Антиконтрафактный код уже активирован"
	languageMapRu["E_NO_ANTI_FAKE_CODE"] = "Антиконтрафактный код не найден"
	languageMapRu["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "Этот антиконтрафактный код уже подтверждён другим пользователем"
	languageMapRu["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "Антиконтрафактный код не активирован, подтверждение невозможно"
	/*mail*/
	languageMapRu["MAIL_VERIFY_SUBJECT"] = "Код подтверждения TintaShift"
	languageMapRu["MAIL_VERIFY_TITLE"] = "Добро пожаловать в TintaShift"
	languageMapRu["MAIL_VERIFY_DESC"] = "Здравствуйте! Вы проходите проверку безопасности аккаунта. Ваш код:"
	languageMapRu["MAIL_VERIFY_SECURITY_TITLE"] = "Совет по безопасности: "
	languageMapRu["MAIL_VERIFY_SECURITY_DESC"] = "Сотрудники TintaShift никогда не запрашивают этот код. Не сообщайте его никому."
	languageMapRu["MAIL_VERIFY_IGNORE"] = "Если вы не запрашивали этот код, проигнорируйте это письмо."
	languageMapRu["MAIL_COPYRIGHT"] = "© 2026 TintaShift. Все права защищены."
	/*social login*/
	languageMapRu["E_NO_GOOGLE_ACCESS_TOKEN"] = "Требуется токен доступа Google"
	languageMapRu["E_GOOGLE_NOT_CONFIGURED"] = "Вход через Google не настроен на сервере"
	languageMapRu["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Недействительный или просроченный токен доступа Google"
	languageMapRu["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "Электронная почта Google не подтверждена, вход невозможен"
	languageMapRu["E_GOOGLE_EMAIL_BOUND_OTHER"] = "Эта почта уже привязана к другому аккаунту Google"
	languageMapRu["E_NO_APPLE_ID_TOKEN"] = "Требуется identity token Apple"
	languageMapRu["E_APPLE_NOT_CONFIGURED"] = "Вход через Apple не настроен на сервере"
	languageMapRu["E_APPLE_ID_TOKEN_INVALID"] = "Недействительный или просроченный identity token Apple"
	languageMapRu["E_APPLE_EMAIL_NOT_VERIFIED"] = "Электронная почта Apple не подтверждена, вход невозможен"
	languageMapRu["E_APPLE_EMAIL_BOUND_OTHER"] = "Эта почта уже привязана к другому аккаунту Apple"
	languageMapRu["E_SOCIAL_NOT_CONFIGURED"] = "Секрет подписи социального входа не настроен на сервере"
	languageMapRu["E_NO_PROVIDER"] = "требуется provider"
	languageMapRu["E_SOCIAL_PROVIDER_INVALID"] = "неподдерживаемый provider социального входа"
	languageMapRu["E_NO_OPENID"] = "требуется openid"
	languageMapRu["E_NO_NONCE"] = "требуется nonce"
	languageMapRu["E_NO_TIMESTAMP"] = "требуется timestamp"
	languageMapRu["E_NO_SIGNATURE"] = "требуется signature"
	languageMapRu["E_NO_DEVICE_ID"] = "требуется device_id"
	languageMapRu["E_REQUEST_EXPIRED"] = "срок действия запроса истёк, повторите попытку"
	languageMapRu["E_CHALLENGE_INVALID"] = "недействительный challenge"
	languageMapRu["E_CHALLENGE_EXPIRED"] = "срок действия challenge истёк"
	languageMapRu["E_CHALLENGE_USED"] = "challenge уже использован"
	languageMapRu["E_SIGNATURE_INVALID"] = "проверка подписи не удалась"
	languageMapRu["E_SOCIAL_EMAIL_BOUND_OTHER"] = "Эта почта уже привязана к другому социальному аккаунту"
	languageMapRu["E_INVALID_VIDEO"] = "Загрузите корректное видео (MP4/MOV)"
	registerLocale("ru-RU", languageMapRu)
}
