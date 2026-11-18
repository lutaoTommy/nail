package language

var languageMapFr map[string]string

/*初始化*/
func initLanguageFrf() {
	languageMapFr = make(map[string]string)
	/*user*/
	languageMapFr["E_NO_IV"] = "Veuillez saisir le vecteur d'initialisation"
	languageMapFr["E_NO_CODE"] = "Veuillez saisir le code d'autorisation"
	languageMapFr["E_NO_DATA"] = "Veuillez saisir les données"
	languageMapFr["E_NO_PHONE"] = "Veuillez saisir le numéro de téléphone"
	languageMapFr["E_NO_PWD"] = "Veuillez saisir le mot de passe"
	languageMapFr["E_INVALID_USER"] = "Utilisateur introuvable"
	languageMapFr["E_INVALID_PWD"] = "Mot de passe incorrect"
	languageMapFr["E_NO_NAME"] = "Veuillez saisir le nom"
	languageMapFr["E_USER_EXIST"] = "Le compte existe déjà"
	languageMapFr["E_NO_TOKEN"] = "Veuillez vous connecter d'abord"
	languageMapFr["E_NO_AUTH"] = "Aucune autorisation"
	languageMapFr["E_NO_AVATAR"] = "Avatar introuvable"
	languageMapFr["E_NO_EMAIL"] = "Veuillez saisir l'e-mail"
	languageMapFr["E_NO_CERT"] = "Veuillez saisir le code de vérification"
	languageMapFr["E_INVALID_EMAIL"] = "Format d'e-mail invalide"
	languageMapFr["E_DATA_TOO_SHORT"] = "Les données chiffrées sont trop courtes"
	languageMapFr["E_INVALID_IV"] = "Vecteur d'initialisation invalide"
	languageMapFr["E_INVALID_SESSION_KEY"] = "Clé de session invalide"
	languageMapFr["E_DECRYPT_FAIL"] = "Échec du déchiffrement, assurez-vous que le code et le numéro de téléphone proviennent de la même session"
	languageMapFr["E_INVALID_PHONE"] = "Format de numéro de téléphone invalide"
	languageMapFr["VERIFICATION_CODE"] = "code de vérification"
	languageMapFr["DO_NOT_TELL"] = "Ne le dites à personne"
	languageMapFr["E_CERT_ERR"] = "Code de vérification incorrect"
	languageMapFr["E_CERT_EXPIRED"] = "Le code de vérification a expiré. Veuillez en demander un nouveau"
	languageMapFr["E_CERT_FIRST"] = "Veuillez d'abord demander un code de vérification"
	languageMapFr["E_VERIFICATION_QUICKLY"] = "Vous demandez des codes de vérification trop fréquemment. Veuillez patienter un instant"
	languageMapFr["E_ACCOUNT_LOCKED"] = "Trop de tentatives. Veuillez réessayer plus tard"
	languageMapFr["E_ACCOUNT_LOCKED_RETRY"] = "Trop de tentatives. Veuillez réessayer dans %d secondes"
	languageMapFr["E_RETRIES_REMAINING"] = ". %d tentative(s) restante(s)"
	languageMapFr["E_VERIFICATION_LIMIT"] = "Trop de demandes de codes de vérification. Veuillez réessayer plus tard"
	/*common*/
	languageMapFr["E_TOO_LONG"] = "La longueur du texte dépasse la limite"
	languageMapFr["E_INVALID_PARAM"] = "Paramètre invalide"
	languageMapFr["E_FILE_NOT_FOUND"] = "Fichier introuvable"
	/*color*/
	languageMapFr["E_INVALID_COLOR"] = "Couleur introuvable"
	languageMapFr["E_INVALID_FILE"] = "Veuillez télécharger un fichier correct"
	languageMapFr["E_INVALID_PICTURE"] = "Veuillez télécharger une image (JPG/PNG)"
	languageMapFr["E_AI_RECOMMEND_FAILED"] = "Le service de recommandation de couleurs est temporairement indisponible, veuillez réessayer plus tard"
	/*device*/
	languageMapFr["E_NO_ID"] = "Veuillez saisir l'ID"
	languageMapFr["E_NO_MAC"] = "Veuillez saisir l'adresse MAC"
	languageMapFr["E_NO_INFO"] = "Enregistrement introuvable"
	languageMapFr["E_NO_DEVICE"] = "Appareil introuvable"
	languageMapFr["E_INVALID_MAC"] = "Veuillez saisir l'adresse MAC au format correct"
	/*suggest*/
	languageMapFr["E_NO_CONTENT"] = "Veuillez saisir le contenu"
	languageMapFr["E_NO_SUGGEST"] = "Suggestion introuvable"
	/*word*/
	languageMapFr["E_NO_WORD"] = "Mot sensible introuvable"
	languageMapFr["E_WORD_EXIST"] = "Le mot sensible existe déjà"
	/*follow*/
	languageMapFr["E_FOLLOWED"] = "Déjà suivi"
	languageMapFr["E_UNFOLLOWED"] = "Plus suivi"
	languageMapFr["E_NOT_FOLLOWING"] = "Ne suit pas"
	languageMapFr["E_NO_SELF_FOLLOW"] = "Impossible de se suivre soi-même"
	/*circle*/
	languageMapFr["E_NO_CIRCLE_POST"] = "Publication introuvable"
	languageMapFr["E_INVALID_REVIEW_STATUS"] = "Statut de révision invalide"
	languageMapFr["E_SENSITIVEWORD"] = "Votre message contient des informations sensibles. Pour maintenir un environnement de communication amical, veuillez ajuster les expressions concernées avant de publier"
	/*topic*/
	languageMapFr["E_NO_TOPIC"] = "Sujet introuvable"
	languageMapFr["E_TOPIC_EXIST"] = "Le sujet existe déjà"
	/*activity*/
	languageMapFr["E_NO_ACTIVITY"] = "Activité introuvable"
	languageMapFr["E_NO_TITLE"] = "Veuillez saisir le titre"
	languageMapFr["E_NO_TIME"] = "Veuillez saisir l'heure"
	languageMapFr["E_ALREADY_JOINED"] = "Déjà inscrit à cette activité"
	languageMapFr["E_NOT_JOINED"] = "Pas inscrit à cette activité"
	languageMapFr["E_ALREADY_VOTED"] = "Déjà voté"
	languageMapFr["E_NOT_VOTED"] = "Pas voté"
	/*comment*/
	languageMapFr["E_NO_COMMENT"] = "Commentaire introuvable"
	/*like*/
	languageMapFr["E_NO_LIKE"] = "J'aime introuvable"
	/*collect*/
	languageMapFr["E_NO_COLLECT"] = "Pas enregistré"
	/*ota*/
	languageMapFr["E_NO_VERSION"] = "Veuillez saisir la version"
	/*voice*/
	languageMapFr["E_NO_VOICE_TEXT"] = "Veuillez saisir une commande vocale"
	languageMapFr["E_NO_COLOR_NAME"] = "Aucun nom de couleur reconnu"
	languageMapFr["V_UNKNOWN_COMMAND"] = "Commande vocale inconnue"
	languageMapFr["V_ACTION_HELP"] = "Impossible de reconnaître la commande. Essayez :\n· Changer la couleur : changer le premier en bleu minuit"
	/*checkin*/
	languageMapFr["E_ALREADY_CHECKIN"] = "Déjà pointé aujourd'hui"
	/*anti_fake*/
	languageMapFr["E_INVALID_PRODUCT_MODEL"] = "Format de modèle de produit invalide (2 caractères alphanumériques requis)"
	languageMapFr["E_INVALID_BATCH"] = "Format de lot de production invalide (6 caractères alphanumériques requis)"
	languageMapFr["E_INVALID_NUMBER"] = "Format de numéro invalide (2 caractères alphanumériques requis)"
	languageMapFr["E_INVALID_ANTI_FAKE_CODE"] = "Format de code anticontrefaçon invalide (24 caractères alphanumériques requis)"
	languageMapFr["E_GENERATE_CODE_FAILED"] = "Échec de la génération du code anticontrefaçon"
	languageMapFr["E_ANTI_FAKE_CODE_ACTIVATED"] = "Code anticontrefaçon déjà activé"
	languageMapFr["E_NO_ANTI_FAKE_CODE"] = "Code anticontrefaçon introuvable"
	languageMapFr["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "Ce code anticontrefaçon a déjà été vérifié par un autre utilisateur"
	languageMapFr["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "Code anticontrefaçon non activé, vérification impossible"
	/*mail*/
	languageMapFr["MAIL_VERIFY_SUBJECT"] = "Code de vérification TintaShift"
	languageMapFr["MAIL_VERIFY_TITLE"] = "Bienvenue sur TintaShift"
	languageMapFr["MAIL_VERIFY_DESC"] = "Bonjour, vous effectuez une vérification de sécurité du compte. Votre code est :"
	languageMapFr["MAIL_VERIFY_SECURITY_TITLE"] = "Conseil de sécurité : "
	languageMapFr["MAIL_VERIFY_SECURITY_DESC"] = "Le personnel de TintaShift ne vous demandera jamais ce code. Ne le partagez avec personne."
	languageMapFr["MAIL_VERIFY_IGNORE"] = "Si vous n'avez pas demandé ce code, veuillez ignorer cet e-mail."
	languageMapFr["MAIL_COPYRIGHT"] = "© 2026 TintaShift. Tous droits réservés."
	/*social login*/
	languageMapFr["E_NO_GOOGLE_ACCESS_TOKEN"] = "Le jeton d'accès Google est requis"
	languageMapFr["E_GOOGLE_NOT_CONFIGURED"] = "La connexion Google n'est pas configurée sur le serveur"
	languageMapFr["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Jeton d'accès Google invalide ou expiré"
	languageMapFr["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "L'e-mail Google n'est pas vérifié et ne peut pas être utilisé pour se connecter"
	languageMapFr["E_GOOGLE_EMAIL_BOUND_OTHER"] = "Cet e-mail est déjà lié à un autre compte Google"
	languageMapFr["E_NO_APPLE_ID_TOKEN"] = "Le jeton d'identité Apple est requis"
	languageMapFr["E_APPLE_NOT_CONFIGURED"] = "La connexion Apple n'est pas configurée sur le serveur"
	languageMapFr["E_APPLE_ID_TOKEN_INVALID"] = "Jeton d'identité Apple invalide ou expiré"
	languageMapFr["E_APPLE_EMAIL_NOT_VERIFIED"] = "L'e-mail Apple n'est pas vérifié et ne peut pas être utilisé pour se connecter"
	languageMapFr["E_APPLE_EMAIL_BOUND_OTHER"] = "Cet e-mail est déjà lié à un autre compte Apple"
	languageMapFr["E_SOCIAL_NOT_CONFIGURED"] = "Le secret de signature de la connexion sociale n'est pas configuré sur le serveur"
	languageMapFr["E_NO_PROVIDER"] = "provider est requis"
	languageMapFr["E_SOCIAL_PROVIDER_INVALID"] = "provider de connexion sociale non pris en charge"
	languageMapFr["E_NO_OPENID"] = "openid est requis"
	languageMapFr["E_NO_NONCE"] = "nonce est requis"
	languageMapFr["E_NO_TIMESTAMP"] = "timestamp est requis"
	languageMapFr["E_NO_SIGNATURE"] = "signature est requise"
	languageMapFr["E_NO_DEVICE_ID"] = "device_id est requis"
	languageMapFr["E_REQUEST_EXPIRED"] = "la requête a expiré, veuillez réessayer"
	languageMapFr["E_CHALLENGE_INVALID"] = "challenge invalide"
	languageMapFr["E_CHALLENGE_EXPIRED"] = "challenge expiré"
	languageMapFr["E_CHALLENGE_USED"] = "challenge déjà utilisé"
	languageMapFr["E_SIGNATURE_INVALID"] = "échec de la vérification de la signature"
	languageMapFr["E_SOCIAL_EMAIL_BOUND_OTHER"] = "Cet e-mail est déjà lié à un autre compte social"
	languageMapFr["E_INVALID_VIDEO"] = "Veuillez télécharger une vidéo valide (MP4/MOV)"
	registerLocale("fr-FR", languageMapFr)
}
