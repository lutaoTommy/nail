package language

var languageMapPl map[string]string

/*初始化*/
func initLanguagePlpl() {
	languageMapPl = make(map[string]string)
	/*user*/
	languageMapPl["E_NO_IV"] = "Wprowadź wektor inicjalizacji"
	languageMapPl["E_NO_CODE"] = "Wprowadź kod autoryzacji"
	languageMapPl["E_NO_DATA"] = "Wprowadź dane"
	languageMapPl["E_NO_PHONE"] = "Wprowadź numer telefonu"
	languageMapPl["E_NO_PWD"] = "Wprowadź hasło"
	languageMapPl["E_INVALID_USER"] = "Nie znaleziono użytkownika"
	languageMapPl["E_INVALID_PWD"] = "Nieprawidłowe hasło"
	languageMapPl["E_NO_NAME"] = "Wprowadź nazwę"
	languageMapPl["E_USER_EXIST"] = "Konto już istnieje"
	languageMapPl["E_NO_TOKEN"] = "Najpierw się zaloguj"
	languageMapPl["E_NO_AUTH"] = "Brak uprawnień"
	languageMapPl["E_NO_AVATAR"] = "Nie znaleziono awatara"
	languageMapPl["E_NO_EMAIL"] = "Wprowadź e-mail"
	languageMapPl["E_NO_CERT"] = "Wprowadź kod weryfikacyjny"
	languageMapPl["E_INVALID_EMAIL"] = "Nieprawidłowy format e-mail"
	languageMapPl["E_DATA_TOO_SHORT"] = "Zaszyfrowane dane są zbyt krótkie"
	languageMapPl["E_INVALID_IV"] = "Nieprawidłowy wektor inicjalizacji"
	languageMapPl["E_INVALID_SESSION_KEY"] = "Nieprawidłowy klucz sesji"
	languageMapPl["E_DECRYPT_FAIL"] = "Odszyfrowanie nie powiodło się; upewnij się, że kod i numer telefonu pochodzą z tej samej sesji"
	languageMapPl["E_INVALID_PHONE"] = "Nieprawidłowy format numeru telefonu"
	languageMapPl["VERIFICATION_CODE"] = "kod weryfikacyjny"
	languageMapPl["DO_NOT_TELL"] = "Nie mów nikomu"
	languageMapPl["E_CERT_ERR"] = "Nieprawidłowy kod weryfikacyjny"
	languageMapPl["E_CERT_EXPIRED"] = "Kod weryfikacyjny wygasł. Poproś o nowy"
	languageMapPl["E_CERT_FIRST"] = "Najpierw poproś o kod weryfikacyjny"
	languageMapPl["E_VERIFICATION_QUICKLY"] = "Zbyt często żądasz kodów weryfikacyjnych. Poczekaj chwilę"
	languageMapPl["E_ACCOUNT_LOCKED"] = "Zbyt wiele prób. Spróbuj ponownie później"
	languageMapPl["E_ACCOUNT_LOCKED_RETRY"] = "Zbyt wiele prób. Spróbuj ponownie za %d sekund"
	languageMapPl["E_RETRIES_REMAINING"] = ". Pozostało prób: %d"
	languageMapPl["E_VERIFICATION_LIMIT"] = "Zbyt wiele żądań kodów weryfikacyjnych. Spróbuj ponownie później"
	/*common*/
	languageMapPl["E_TOO_LONG"] = "Długość tekstu przekracza limit"
	languageMapPl["E_INVALID_PARAM"] = "Nieprawidłowy parametr"
	languageMapPl["E_FILE_NOT_FOUND"] = "Nie znaleziono pliku"
	/*color*/
	languageMapPl["E_INVALID_COLOR"] = "Nie znaleziono koloru"
	languageMapPl["E_INVALID_FILE"] = "Prześlij poprawny plik"
	languageMapPl["E_INVALID_PICTURE"] = "Prześlij obraz (JPG/PNG)"
	languageMapPl["E_AI_RECOMMEND_FAILED"] = "Usługa rekomendacji kolorów jest tymczasowo niedostępna, spróbuj ponownie później"
	/*device*/
	languageMapPl["E_NO_ID"] = "Wprowadź ID"
	languageMapPl["E_NO_MAC"] = "Wprowadź adres MAC"
	languageMapPl["E_NO_INFO"] = "Nie znaleziono rekordu"
	languageMapPl["E_NO_DEVICE"] = "Nie znaleziono urządzenia"
	languageMapPl["E_INVALID_MAC"] = "Wprowadź adres MAC w poprawnym formacie"
	/*suggest*/
	languageMapPl["E_NO_CONTENT"] = "Wprowadź treść"
	languageMapPl["E_NO_SUGGEST"] = "Nie znaleziono sugestii"
	/*word*/
	languageMapPl["E_NO_WORD"] = "Nie znaleziono słowa wrażliwego"
	languageMapPl["E_WORD_EXIST"] = "Słowo wrażliwe już istnieje"
	/*follow*/
	languageMapPl["E_FOLLOWED"] = "Już obserwowane"
	languageMapPl["E_UNFOLLOWED"] = "Już nieobserwowane"
	languageMapPl["E_NOT_FOLLOWING"] = "Nieobserwowane"
	languageMapPl["E_NO_SELF_FOLLOW"] = "Nie można obserwować siebie"
	/*circle*/
	languageMapPl["E_NO_CIRCLE_POST"] = "Nie znaleziono wpisu"
	languageMapPl["E_INVALID_REVIEW_STATUS"] = "Nieprawidłowy status przeglądu"
	languageMapPl["E_SENSITIVEWORD"] = "Twoja wiadomość zawiera wrażliwe informacje. Aby utrzymać przyjazną atmosferę komunikacji, dostosuj odpowiednie sformułowania przed publikacją"
	/*topic*/
	languageMapPl["E_NO_TOPIC"] = "Nie znaleziono tematu"
	languageMapPl["E_TOPIC_EXIST"] = "Temat już istnieje"
	/*activity*/
	languageMapPl["E_NO_ACTIVITY"] = "Nie znaleziono aktywności"
	languageMapPl["E_NO_TITLE"] = "Wprowadź tytuł"
	languageMapPl["E_NO_TIME"] = "Wprowadź czas"
	languageMapPl["E_ALREADY_JOINED"] = "Już dołączono do tej aktywności"
	languageMapPl["E_NOT_JOINED"] = "Nie dołączono do tej aktywności"
	languageMapPl["E_ALREADY_VOTED"] = "Już oddano głos"
	languageMapPl["E_NOT_VOTED"] = "Nie oddano głosu"
	/*comment*/
	languageMapPl["E_NO_COMMENT"] = "Nie znaleziono komentarza"
	/*like*/
	languageMapPl["E_NO_LIKE"] = "Nie znaleziono polubienia"
	/*collect*/
	languageMapPl["E_NO_COLLECT"] = "Niezapisane"
	/*ota*/
	languageMapPl["E_NO_VERSION"] = "Wprowadź wersję"
	/*voice*/
	languageMapPl["E_NO_VOICE_TEXT"] = "Wprowadź polecenie głosowe"
	languageMapPl["E_NO_COLOR_NAME"] = "Nie rozpoznano nazwy koloru"
	languageMapPl["V_UNKNOWN_COMMAND"] = "Nieznane polecenie głosowe"
	languageMapPl["V_ACTION_HELP"] = "Nie udało się rozpoznać polecenia. Spróbuj:\n· Zmień kolor: zmień pierwszy na granatowy"
	/*checkin*/
	languageMapPl["E_ALREADY_CHECKIN"] = "Już zalogowano się dzisiaj"
	/*anti_fake*/
	languageMapPl["E_INVALID_PRODUCT_MODEL"] = "Nieprawidłowy format modelu produktu (wymagane 2 znaki alfanumeryczne)"
	languageMapPl["E_INVALID_BATCH"] = "Nieprawidłowy format partii produkcyjnej (wymagane 6 znaków alfanumerycznych)"
	languageMapPl["E_INVALID_NUMBER"] = "Nieprawidłowy format numeru (wymagane 2 znaki alfanumeryczne)"
	languageMapPl["E_INVALID_ANTI_FAKE_CODE"] = "Nieprawidłowy format kodu antyfałszerskiego (wymagane 24 znaki alfanumeryczne)"
	languageMapPl["E_GENERATE_CODE_FAILED"] = "Nie udało się wygenerować kodu antyfałszerskiego"
	languageMapPl["E_ANTI_FAKE_CODE_ACTIVATED"] = "Kod antyfałszerski już aktywowany"
	languageMapPl["E_NO_ANTI_FAKE_CODE"] = "Nie znaleziono kodu antyfałszerskiego"
	languageMapPl["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "Ten kod antyfałszerski został już zweryfikowany przez innego użytkownika"
	languageMapPl["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "Kod antyfałszerski nieaktywny, weryfikacja niemożliwa"
	/*mail*/
	languageMapPl["MAIL_VERIFY_SUBJECT"] = "Kod weryfikacyjny TintaShift"
	languageMapPl["MAIL_VERIFY_TITLE"] = "Witamy w TintaShift"
	languageMapPl["MAIL_VERIFY_DESC"] = "Cześć, przeprowadzasz weryfikację bezpieczeństwa konta. Twój kod to:"
	languageMapPl["MAIL_VERIFY_SECURITY_TITLE"] = "Wskazówka bezpieczeństwa: "
	languageMapPl["MAIL_VERIFY_SECURITY_DESC"] = "Personel TintaShift nigdy nie poprosi o ten kod. Nie udostępniaj go nikomu."
	languageMapPl["MAIL_VERIFY_IGNORE"] = "Jeśli nie prosiłeś o ten kod, zignoruj tę wiadomość e-mail."
	languageMapPl["MAIL_COPYRIGHT"] = "© 2026 TintaShift. Wszelkie prawa zastrzeżone."
	/*social login*/
	languageMapPl["E_NO_GOOGLE_ACCESS_TOKEN"] = "Wymagany jest token dostępu Google"
	languageMapPl["E_GOOGLE_NOT_CONFIGURED"] = "Logowanie Google nie jest skonfigurowane na serwerze"
	languageMapPl["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Nieprawidłowy lub wygasły token dostępu Google"
	languageMapPl["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "Adres e-mail Google nie jest zweryfikowany i nie można go użyć do logowania"
	languageMapPl["E_GOOGLE_EMAIL_BOUND_OTHER"] = "Ten adres e-mail jest już powiązany z innym kontem Google"
	languageMapPl["E_NO_APPLE_ID_TOKEN"] = "Wymagany jest token tożsamości Apple"
	languageMapPl["E_APPLE_NOT_CONFIGURED"] = "Logowanie Apple nie jest skonfigurowane na serwerze"
	languageMapPl["E_APPLE_ID_TOKEN_INVALID"] = "Nieprawidłowy lub wygasły token tożsamości Apple"
	languageMapPl["E_APPLE_EMAIL_NOT_VERIFIED"] = "Adres e-mail Apple nie jest zweryfikowany i nie można go użyć do logowania"
	languageMapPl["E_APPLE_EMAIL_BOUND_OTHER"] = "Ten adres e-mail jest już powiązany z innym kontem Apple"
	languageMapPl["E_SOCIAL_NOT_CONFIGURED"] = "Sekret podpisu logowania społecznościowego nie jest skonfigurowany na serwerze"
	languageMapPl["E_NO_PROVIDER"] = "wymagany jest provider"
	languageMapPl["E_SOCIAL_PROVIDER_INVALID"] = "nieobsługiwany provider logowania społecznościowego"
	languageMapPl["E_NO_OPENID"] = "wymagany jest openid"
	languageMapPl["E_NO_NONCE"] = "wymagany jest nonce"
	languageMapPl["E_NO_TIMESTAMP"] = "wymagany jest timestamp"
	languageMapPl["E_NO_SIGNATURE"] = "wymagany jest signature"
	languageMapPl["E_NO_DEVICE_ID"] = "wymagany jest device_id"
	languageMapPl["E_REQUEST_EXPIRED"] = "żądanie wygasło, spróbuj ponownie"
	languageMapPl["E_CHALLENGE_INVALID"] = "nieprawidłowy challenge"
	languageMapPl["E_CHALLENGE_EXPIRED"] = "challenge wygasł"
	languageMapPl["E_CHALLENGE_USED"] = "challenge został już użyty"
	languageMapPl["E_SIGNATURE_INVALID"] = "weryfikacja podpisu nie powiodła się"
	languageMapPl["E_SOCIAL_EMAIL_BOUND_OTHER"] = "Ten adres e-mail jest już powiązany z innym kontem społecznościowym"
	languageMapPl["E_INVALID_VIDEO"] = "Prześlij prawidłowe wideo (MP4/MOV)"
	registerLocale("pl-PL", languageMapPl)
}
