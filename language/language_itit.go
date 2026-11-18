package language

var languageMapIt map[string]string

/*初始化*/
func initLanguageItit() {
	languageMapIt = make(map[string]string)
	/*user*/
	languageMapIt["E_NO_IV"] = "Immettere il vettore di inizializzazione"
	languageMapIt["E_NO_CODE"] = "Immettere il codice di autorizzazione"
	languageMapIt["E_NO_DATA"] = "Immettere i dati"
	languageMapIt["E_NO_PHONE"] = "Immettere il numero di telefono"
	languageMapIt["E_NO_PWD"] = "Immettere la password"
	languageMapIt["E_INVALID_USER"] = "Utente non trovato"
	languageMapIt["E_INVALID_PWD"] = "Password errata"
	languageMapIt["E_NO_NAME"] = "Immettere il nome"
	languageMapIt["E_USER_EXIST"] = "L'account esiste già"
	languageMapIt["E_NO_TOKEN"] = "Accedere prima dell'operazione"
	languageMapIt["E_NO_AUTH"] = "Nessuna autorizzazione"
	languageMapIt["E_NO_AVATAR"] = "Avatar non trovato"
	languageMapIt["E_NO_EMAIL"] = "Immettere l'e-mail"
	languageMapIt["E_NO_CERT"] = "Immettere il codice di verifica"
	languageMapIt["E_INVALID_EMAIL"] = "Formato dell'e-mail non valido"
	languageMapIt["E_DATA_TOO_SHORT"] = "I dati cifrati sono troppo corti"
	languageMapIt["E_INVALID_IV"] = "Vettore di inizializzazione non valido"
	languageMapIt["E_INVALID_SESSION_KEY"] = "Chiave di sessione non valida"
	languageMapIt["E_DECRYPT_FAIL"] = "Decrittazione non riuscita, assicurarsi che il codice e il numero di telefono provengano dalla stessa sessione"
	languageMapIt["E_INVALID_PHONE"] = "Formato del numero di telefono non valido"
	languageMapIt["VERIFICATION_CODE"] = "codice di verifica"
	languageMapIt["DO_NOT_TELL"] = "Non comunicarlo ad altri"
	languageMapIt["E_CERT_ERR"] = "Codice di verifica errato"
	languageMapIt["E_CERT_EXPIRED"] = "Il codice di verifica è scaduto. Richiederne uno nuovo"
	languageMapIt["E_CERT_FIRST"] = "Richiedere prima il codice di verifica"
	languageMapIt["E_VERIFICATION_QUICKLY"] = "Si stanno richiedendo troppi codici di verifica. Attendere un momento"
	languageMapIt["E_ACCOUNT_LOCKED"] = "Troppi tentativi. Riprovare più tardi"
	languageMapIt["E_ACCOUNT_LOCKED_RETRY"] = "Troppi tentativi. Riprovare tra %d secondi"
	languageMapIt["E_RETRIES_REMAINING"] = ". %d tentativo/i rimasto/i"
	languageMapIt["E_VERIFICATION_LIMIT"] = "Troppe richieste di codici di verifica. Riprovare più tardi"
	/*common*/
	languageMapIt["E_TOO_LONG"] = "La lunghezza del testo supera il limite"
	languageMapIt["E_INVALID_PARAM"] = "Parametro non valido"
	languageMapIt["E_FILE_NOT_FOUND"] = "File non trovato"
	/*color*/
	languageMapIt["E_INVALID_COLOR"] = "Colore non trovato"
	languageMapIt["E_INVALID_FILE"] = "Caricare un file corretto"
	languageMapIt["E_INVALID_PICTURE"] = "Caricare un'immagine (JPG/PNG)"
	languageMapIt["E_AI_RECOMMEND_FAILED"] = "Il servizio di raccomandazione colori è temporaneamente non disponibile, riprovare più tardi"
	/*device*/
	languageMapIt["E_NO_ID"] = "Immettere l'ID"
	languageMapIt["E_NO_MAC"] = "Immettere l'indirizzo MAC"
	languageMapIt["E_NO_INFO"] = "Record non trovato"
	languageMapIt["E_NO_DEVICE"] = "Dispositivo non trovato"
	languageMapIt["E_INVALID_MAC"] = "Immettere l'indirizzo MAC nel formato corretto"
	/*suggest*/
	languageMapIt["E_NO_CONTENT"] = "Immettere il contenuto"
	languageMapIt["E_NO_SUGGEST"] = "Suggerimento non trovato"
	/*word*/
	languageMapIt["E_NO_WORD"] = "Parola sensibile non trovata"
	languageMapIt["E_WORD_EXIST"] = "La parola sensibile esiste già"
	/*follow*/
	languageMapIt["E_FOLLOWED"] = "Già seguito"
	languageMapIt["E_UNFOLLOWED"] = "Non più seguito"
	languageMapIt["E_NOT_FOLLOWING"] = "Non seguito"
	languageMapIt["E_NO_SELF_FOLLOW"] = "Impossibile seguire se stessi"
	/*circle*/
	languageMapIt["E_NO_CIRCLE_POST"] = "Post non trovato"
	languageMapIt["E_INVALID_REVIEW_STATUS"] = "Stato di revisione non valido"
	languageMapIt["E_SENSITIVEWORD"] = "Il messaggio contiene informazioni sensibili. Per mantenere un ambiente di comunicazione amichevole, ajustare le espressioni pertinenti prima di pubblicare"
	/*topic*/
	languageMapIt["E_NO_TOPIC"] = "Argomento non trovato"
	languageMapIt["E_TOPIC_EXIST"] = "L'argomento esiste già"
	/*activity*/
	languageMapIt["E_NO_ACTIVITY"] = "Attività non trovata"
	languageMapIt["E_NO_TITLE"] = "Immettere il titolo"
	languageMapIt["E_NO_TIME"] = "Immettere l'ora"
	languageMapIt["E_ALREADY_JOINED"] = "Già iscritto a questa attività"
	languageMapIt["E_NOT_JOINED"] = "Non iscritto a questa attività"
	languageMapIt["E_ALREADY_VOTED"] = "Già votato"
	languageMapIt["E_NOT_VOTED"] = "Non votato"
	/*comment*/
	languageMapIt["E_NO_COMMENT"] = "Commento non trovato"
	/*like*/
	languageMapIt["E_NO_LIKE"] = "Mi piace non trovato"
	/*collect*/
	languageMapIt["E_NO_COLLECT"] = "Non salvato"
	/*ota*/
	languageMapIt["E_NO_VERSION"] = "Immettere la versione"
	/*voice*/
	languageMapIt["E_NO_VOICE_TEXT"] = "Immettere un comando vocale"
	languageMapIt["E_NO_COLOR_NAME"] = "Nessun nome di colore riconosciuto"
	languageMapIt["V_UNKNOWN_COMMAND"] = "Comando vocale sconosciuto"
	languageMapIt["V_ACTION_HELP"] = "Impossibile riconoscere il comando. Provare:\n· Cambia colore: cambia il primo in blu mezzanotte"
	/*checkin*/
	languageMapIt["E_ALREADY_CHECKIN"] = "Già registrato oggi"
	/*anti_fake*/
	languageMapIt["E_INVALID_PRODUCT_MODEL"] = "Formato del modello di prodotto non valido (richiesti 2 caratteri alfanumerici)"
	languageMapIt["E_INVALID_BATCH"] = "Formato del lotto di produzione non valido (richiesti 6 caratteri alfanumerici)"
	languageMapIt["E_INVALID_NUMBER"] = "Formato del numero non valido (richiesti 2 caratteri alfanumerici)"
	languageMapIt["E_INVALID_ANTI_FAKE_CODE"] = "Formato del codice anticontraffazione non valido (richiesti 24 caratteri alfanumerici)"
	languageMapIt["E_GENERATE_CODE_FAILED"] = "Generazione del codice anticontraffazione non riuscita"
	languageMapIt["E_ANTI_FAKE_CODE_ACTIVATED"] = "Codice anticontraffazione già attivato"
	languageMapIt["E_NO_ANTI_FAKE_CODE"] = "Codice anticontraffazione non trovato"
	languageMapIt["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "Questo codice anticontraffazione è già stato verificato da un altro utente"
	languageMapIt["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "Codice anticontraffazione non attivato, verifica impossibile"
	/*mail*/
	languageMapIt["MAIL_VERIFY_SUBJECT"] = "Codice di verifica TintaShift"
	languageMapIt["MAIL_VERIFY_TITLE"] = "Benvenuto in TintaShift"
	languageMapIt["MAIL_VERIFY_DESC"] = "Ciao, stai eseguendo una verifica di sicurezza dell'account. Il tuo codice è:"
	languageMapIt["MAIL_VERIFY_SECURITY_TITLE"] = "Suggerimento di sicurezza: "
	languageMapIt["MAIL_VERIFY_SECURITY_DESC"] = "Il personale di TintaShift non ti chiederà mai questo codice. Non condividerlo con nessuno."
	languageMapIt["MAIL_VERIFY_IGNORE"] = "Se non hai richiesto questo codice, ignora questa e-mail."
	languageMapIt["MAIL_COPYRIGHT"] = "© 2026 TintaShift. Tutti i diritti riservati."
	/*social login*/
	languageMapIt["E_NO_GOOGLE_ACCESS_TOKEN"] = "È richiesto il token di accesso Google"
	languageMapIt["E_GOOGLE_NOT_CONFIGURED"] = "L'accesso con Google non è configurato sul server"
	languageMapIt["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Token di accesso Google non valido o scaduto"
	languageMapIt["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "L'e-mail Google non è verificata e non può essere usata per accedere"
	languageMapIt["E_GOOGLE_EMAIL_BOUND_OTHER"] = "Questa e-mail è già collegata a un altro account Google"
	languageMapIt["E_NO_APPLE_ID_TOKEN"] = "È richiesto l'identity token di Apple"
	languageMapIt["E_APPLE_NOT_CONFIGURED"] = "L'accesso con Apple non è configurato sul server"
	languageMapIt["E_APPLE_ID_TOKEN_INVALID"] = "Identity token di Apple non valido o scaduto"
	languageMapIt["E_APPLE_EMAIL_NOT_VERIFIED"] = "L'e-mail Apple non è verificata e non può essere usata per accedere"
	languageMapIt["E_APPLE_EMAIL_BOUND_OTHER"] = "Questa e-mail è già collegata a un altro account Apple"
	languageMapIt["E_SOCIAL_NOT_CONFIGURED"] = "Il segreto di firma dell'accesso social non è configurato sul server"
	languageMapIt["E_NO_PROVIDER"] = "provider è richiesto"
	languageMapIt["E_SOCIAL_PROVIDER_INVALID"] = "provider di accesso social non supportato"
	languageMapIt["E_NO_OPENID"] = "openid è richiesto"
	languageMapIt["E_NO_NONCE"] = "nonce è richiesto"
	languageMapIt["E_NO_TIMESTAMP"] = "timestamp è richiesto"
	languageMapIt["E_NO_SIGNATURE"] = "signature è richiesta"
	languageMapIt["E_NO_DEVICE_ID"] = "device_id è richiesto"
	languageMapIt["E_REQUEST_EXPIRED"] = "richiesta scaduta, riprovare"
	languageMapIt["E_CHALLENGE_INVALID"] = "challenge non valido"
	languageMapIt["E_CHALLENGE_EXPIRED"] = "challenge scaduto"
	languageMapIt["E_CHALLENGE_USED"] = "challenge già utilizzato"
	languageMapIt["E_SIGNATURE_INVALID"] = "verifica della firma non riuscita"
	languageMapIt["E_SOCIAL_EMAIL_BOUND_OTHER"] = "Questa e-mail è già collegata a un altro account social"
	languageMapIt["E_INVALID_VIDEO"] = "Caricare un video valido (MP4/MOV)"
	registerLocale("it-IT", languageMapIt)
}
