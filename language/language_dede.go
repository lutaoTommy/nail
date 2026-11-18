package language

var languageMapDe map[string]string

/*初始化*/
func initLanguageDede() {
	languageMapDe = make(map[string]string)
	/*user*/
	languageMapDe["E_NO_IV"] = "Bitte Initialisierungsvektor eingeben"
	languageMapDe["E_NO_CODE"] = "Bitte Autorisierungscode eingeben"
	languageMapDe["E_NO_DATA"] = "Bitte Daten eingeben"
	languageMapDe["E_NO_PHONE"] = "Bitte Telefonnummer eingeben"
	languageMapDe["E_NO_PWD"] = "Bitte Passwort eingeben"
	languageMapDe["E_INVALID_USER"] = "Benutzer nicht gefunden"
	languageMapDe["E_INVALID_PWD"] = "Falsches Passwort"
	languageMapDe["E_NO_NAME"] = "Bitte Namen eingeben"
	languageMapDe["E_USER_EXIST"] = "Konto existiert bereits"
	languageMapDe["E_NO_TOKEN"] = "Bitte zuerst anmelden"
	languageMapDe["E_NO_AUTH"] = "Keine Berechtigung"
	languageMapDe["E_NO_AVATAR"] = "Avatar nicht gefunden"
	languageMapDe["E_NO_EMAIL"] = "Bitte E-Mail eingeben"
	languageMapDe["E_NO_CERT"] = "Bitte Bestätigungscode eingeben"
	languageMapDe["E_INVALID_EMAIL"] = "E-Mail-Format ungültig"
	languageMapDe["E_DATA_TOO_SHORT"] = "Verschlüsselte Daten zu kurz"
	languageMapDe["E_INVALID_IV"] = "Ungültiger Initialisierungsvektor"
	languageMapDe["E_INVALID_SESSION_KEY"] = "Ungültiger Sitzungsschlüssel"
	languageMapDe["E_DECRYPT_FAIL"] = "Entschlüsselung fehlgeschlagen. Stellen Sie sicher, dass Code und Telefonnummer aus derselben Sitzung stammen"
	languageMapDe["E_INVALID_PHONE"] = "Telefonnummer-Format ungültig"
	languageMapDe["VERIFICATION_CODE"] = "Bestätigungscode"
	languageMapDe["DO_NOT_TELL"] = "Bitte nicht weitergeben"
	languageMapDe["E_CERT_ERR"] = "Bestätigungscode falsch"
	languageMapDe["E_CERT_EXPIRED"] = "Bestätigungscode abgelaufen. Bitte einen neuen anfordern"
	languageMapDe["E_CERT_FIRST"] = "Bitte zuerst Bestätigungscode anfordern"
	languageMapDe["E_VERIFICATION_QUICKLY"] = "Sie fordern Bestätigungscodes zu häufig an. Bitte einen Moment warten"
	languageMapDe["E_ACCOUNT_LOCKED"] = "Zu viele Versuche. Bitte später erneut versuchen"
	languageMapDe["E_ACCOUNT_LOCKED_RETRY"] = "Zu viele Versuche. Bitte in %d Sekunden erneut versuchen"
	languageMapDe["E_RETRIES_REMAINING"] = ". Noch %d Versuch(e) verbleiben"
	languageMapDe["E_VERIFICATION_LIMIT"] = "Zu viele Anfragen nach Bestätigungscodes. Bitte später erneut versuchen"
	/*common*/
	languageMapDe["E_TOO_LONG"] = "Zeichenlänge überschreitet das Limit"
	languageMapDe["E_INVALID_PARAM"] = "Ungültiger Parameter"
	languageMapDe["E_FILE_NOT_FOUND"] = "Datei nicht gefunden"
	/*color*/
	languageMapDe["E_INVALID_COLOR"] = "Farbe nicht gefunden"
	languageMapDe["E_INVALID_FILE"] = "Bitte korrekte Datei hochladen"
	languageMapDe["E_INVALID_PICTURE"] = "Bitte Bild hochladen (JPG/PNG)"
	languageMapDe["E_AI_RECOMMEND_FAILED"] = "Farbempfehlungsdienst vorübergehend nicht verfügbar, bitte später erneut versuchen"
	/*device*/
	languageMapDe["E_NO_ID"] = "Bitte ID eingeben"
	languageMapDe["E_NO_MAC"] = "Bitte MAC-Adresse eingeben"
	languageMapDe["E_NO_INFO"] = "Eintrag nicht gefunden"
	languageMapDe["E_NO_DEVICE"] = "Gerät nicht gefunden"
	languageMapDe["E_INVALID_MAC"] = "Bitte MAC-Adresse im korrekten Format eingeben"
	/*suggest*/
	languageMapDe["E_NO_CONTENT"] = "Bitte Inhalt eingeben"
	languageMapDe["E_NO_SUGGEST"] = "Vorschlag nicht gefunden"
	/*word*/
	languageMapDe["E_NO_WORD"] = "Schimpfwort nicht gefunden"
	languageMapDe["E_WORD_EXIST"] = "Schimpfwort existiert bereits"
	/*follow*/
	languageMapDe["E_FOLLOWED"] = "Bereits gefolgt"
	languageMapDe["E_UNFOLLOWED"] = "Nicht mehr gefolgt"
	languageMapDe["E_NOT_FOLLOWING"] = "Folgt nicht"
	languageMapDe["E_NO_SELF_FOLLOW"] = "Kann sich selbst nicht folgen"
	/*circle*/
	languageMapDe["E_NO_CIRCLE_POST"] = "Beitrag nicht gefunden"
	languageMapDe["E_INVALID_REVIEW_STATUS"] = "Ungültiger Prüfstatus"
	languageMapDe["E_SENSITIVEWORD"] = "Ihre Nachricht enthält sensible Informationen. Um eine freundliche Kommunikationsumgebung zu wahren, passen Sie die betreffenden Ausdrücke vor dem Veröffentlichen an"
	/*topic*/
	languageMapDe["E_NO_TOPIC"] = "Thema nicht gefunden"
	languageMapDe["E_TOPIC_EXIST"] = "Thema existiert bereits"
	/*activity*/
	languageMapDe["E_NO_ACTIVITY"] = "Aktivität nicht gefunden"
	languageMapDe["E_NO_TITLE"] = "Bitte Titel eingeben"
	languageMapDe["E_NO_TIME"] = "Bitte Zeit eingeben"
	languageMapDe["E_ALREADY_JOINED"] = "Bereits an dieser Aktivität teilgenommen"
	languageMapDe["E_NOT_JOINED"] = "Nicht an dieser Aktivität teilgenommen"
	languageMapDe["E_ALREADY_VOTED"] = "Bereits abgestimmt"
	languageMapDe["E_NOT_VOTED"] = "Nicht abgestimmt"
	/*comment*/
	languageMapDe["E_NO_COMMENT"] = "Kommentar nicht gefunden"
	/*like*/
	languageMapDe["E_NO_LIKE"] = "Gefällt mir nicht gefunden"
	/*collect*/
	languageMapDe["E_NO_COLLECT"] = "Nicht gespeichert"
	/*ota*/
	languageMapDe["E_NO_VERSION"] = "Bitte Version eingeben"
	/*voice*/
	languageMapDe["E_NO_VOICE_TEXT"] = "Bitte Sprachbefehl eingeben"
	languageMapDe["E_NO_COLOR_NAME"] = "Kein Farbname erkannt"
	languageMapDe["V_UNKNOWN_COMMAND"] = "Unbekannter Sprachbefehl"
	languageMapDe["V_ACTION_HELP"] = "Befehl konnte nicht erkannt werden. Versuchen Sie:\n· Farbe ändern: den ersten zu Mitternachtsblau ändern"
	/*checkin*/
	languageMapDe["E_ALREADY_CHECKIN"] = "Heute bereits eingecheckt"
	/*anti_fake*/
	languageMapDe["E_INVALID_PRODUCT_MODEL"] = "Produktmodell-Format ungültig (2 alphanumerische Zeichen erforderlich)"
	languageMapDe["E_INVALID_BATCH"] = "Produktionslos-Format ungültig (6 alphanumerische Zeichen erforderlich)"
	languageMapDe["E_INVALID_NUMBER"] = "Nummernformat ungültig (2 alphanumerische Zeichen erforderlich)"
	languageMapDe["E_INVALID_ANTI_FAKE_CODE"] = "Fälschungsschutzcode-Format ungültig (24 alphanumerische Zeichen erforderlich)"
	languageMapDe["E_GENERATE_CODE_FAILED"] = "Fälschungsschutzcode konnte nicht erstellt werden"
	languageMapDe["E_ANTI_FAKE_CODE_ACTIVATED"] = "Fälschungsschutzcode bereits aktiviert"
	languageMapDe["E_NO_ANTI_FAKE_CODE"] = "Fälschungsschutzcode nicht gefunden"
	languageMapDe["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "Dieser Fälschungsschutzcode wurde bereits von einem anderen Benutzer verifiziert"
	languageMapDe["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "Fälschungsschutzcode nicht aktiviert, Verifizierung nicht möglich"
	/*mail*/
	languageMapDe["MAIL_VERIFY_SUBJECT"] = "TintaShift Bestätigungscode"
	languageMapDe["MAIL_VERIFY_TITLE"] = "Willkommen bei TintaShift"
	languageMapDe["MAIL_VERIFY_DESC"] = "Hallo, Sie führen eine Kontosicherheitsbestätigung durch. Ihr Code lautet:"
	languageMapDe["MAIL_VERIFY_SECURITY_TITLE"] = "Sicherheitstipp: "
	languageMapDe["MAIL_VERIFY_SECURITY_DESC"] = "Mitarbeiter von TintaShift werden diesen Code niemals erfragen. Geben Sie ihn an niemanden weiter."
	languageMapDe["MAIL_VERIFY_IGNORE"] = "Wenn Sie diesen Code nicht angefordert haben, ignorieren Sie bitte diese E-Mail."
	languageMapDe["MAIL_COPYRIGHT"] = "© 2026 TintaShift. Alle Rechte vorbehalten."
	/*social login*/
	languageMapDe["E_NO_GOOGLE_ACCESS_TOKEN"] = "Google-Zugriffstoken ist erforderlich"
	languageMapDe["E_GOOGLE_NOT_CONFIGURED"] = "Google-Anmeldung ist auf dem Server nicht konfiguriert"
	languageMapDe["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Ungültiges oder abgelaufenes Google-Zugriffstoken"
	languageMapDe["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "Google-E-Mail ist nicht bestätigt und kann nicht zur Anmeldung verwendet werden"
	languageMapDe["E_GOOGLE_EMAIL_BOUND_OTHER"] = "Diese E-Mail ist bereits mit einem anderen Google-Konto verknüpft"
	languageMapDe["E_NO_APPLE_ID_TOKEN"] = "Apple-Identitätstoken ist erforderlich"
	languageMapDe["E_APPLE_NOT_CONFIGURED"] = "Apple-Anmeldung ist auf dem Server nicht konfiguriert"
	languageMapDe["E_APPLE_ID_TOKEN_INVALID"] = "Ungültiges oder abgelaufenes Apple-Identitätstoken"
	languageMapDe["E_APPLE_EMAIL_NOT_VERIFIED"] = "Apple-E-Mail ist nicht bestätigt und kann nicht zur Anmeldung verwendet werden"
	languageMapDe["E_APPLE_EMAIL_BOUND_OTHER"] = "Diese E-Mail ist bereits mit einem anderen Apple-Konto verknüpft"
	languageMapDe["E_SOCIAL_NOT_CONFIGURED"] = "Das Signaturgeheimnis für die soziale Anmeldung ist auf dem Server nicht konfiguriert"
	languageMapDe["E_NO_PROVIDER"] = "provider ist erforderlich"
	languageMapDe["E_SOCIAL_PROVIDER_INVALID"] = "Nicht unterstützter provider für soziale Anmeldung"
	languageMapDe["E_NO_OPENID"] = "openid ist erforderlich"
	languageMapDe["E_NO_NONCE"] = "nonce ist erforderlich"
	languageMapDe["E_NO_TIMESTAMP"] = "timestamp ist erforderlich"
	languageMapDe["E_NO_SIGNATURE"] = "signature ist erforderlich"
	languageMapDe["E_NO_DEVICE_ID"] = "device_id ist erforderlich"
	languageMapDe["E_REQUEST_EXPIRED"] = "Anfrage abgelaufen, bitte erneut versuchen"
	languageMapDe["E_CHALLENGE_INVALID"] = "ungültige challenge"
	languageMapDe["E_CHALLENGE_EXPIRED"] = "challenge abgelaufen"
	languageMapDe["E_CHALLENGE_USED"] = "challenge bereits verwendet"
	languageMapDe["E_SIGNATURE_INVALID"] = "Signaturprüfung fehlgeschlagen"
	languageMapDe["E_SOCIAL_EMAIL_BOUND_OTHER"] = "Diese E-Mail ist bereits mit einem anderen sozialen Konto verknüpft"
	languageMapDe["E_INVALID_VIDEO"] = "Bitte laden Sie ein gültiges Video hoch (MP4/MOV)"
	registerLocale("de-DE", languageMapDe)
}
