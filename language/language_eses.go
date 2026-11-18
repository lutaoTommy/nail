package language

var languageMapEs map[string]string

/*初始化*/
func initLanguageEses() {
	languageMapEs = make(map[string]string)
	/*user*/
	languageMapEs["E_NO_IV"] = "Introduzca el vector de inicialización"
	languageMapEs["E_NO_CODE"] = "Introduzca el código de autorización"
	languageMapEs["E_NO_DATA"] = "Introduzca los datos"
	languageMapEs["E_NO_PHONE"] = "Introduzca el número de teléfono"
	languageMapEs["E_NO_PWD"] = "Introduzca la contraseña"
	languageMapEs["E_INVALID_USER"] = "Usuario no encontrado"
	languageMapEs["E_INVALID_PWD"] = "Contraseña incorrecta"
	languageMapEs["E_NO_NAME"] = "Introduzca el nombre"
	languageMapEs["E_USER_EXIST"] = "La cuenta ya existe"
	languageMapEs["E_NO_TOKEN"] = "Inicie sesión primero"
	languageMapEs["E_NO_AUTH"] = "Sin permisos"
	languageMapEs["E_NO_AVATAR"] = "Avatar no encontrado"
	languageMapEs["E_NO_EMAIL"] = "Introduzca el correo electrónico"
	languageMapEs["E_NO_CERT"] = "Introduzca el código de verificación"
	languageMapEs["E_INVALID_EMAIL"] = "Formato de correo electrónico no válido"
	languageMapEs["E_DATA_TOO_SHORT"] = "Los datos cifrados son demasiado cortos"
	languageMapEs["E_INVALID_IV"] = "Vector de inicialización no válido"
	languageMapEs["E_INVALID_SESSION_KEY"] = "Clave de sesión no válida"
	languageMapEs["E_DECRYPT_FAIL"] = "Error al descifrar; asegúrese de que el código y el número de teléfono provengan de la misma sesión"
	languageMapEs["E_INVALID_PHONE"] = "Formato de número de teléfono no válido"
	languageMapEs["VERIFICATION_CODE"] = "código de verificación"
	languageMapEs["DO_NOT_TELL"] = "No se lo diga a nadie"
	languageMapEs["E_CERT_ERR"] = "Código de verificación incorrecto"
	languageMapEs["E_CERT_EXPIRED"] = "El código de verificación ha caducado. Solicite uno nuevo"
	languageMapEs["E_CERT_FIRST"] = "Solicite primero el código de verificación"
	languageMapEs["E_VERIFICATION_QUICKLY"] = "Está solicitando códigos de verificación con demasiada frecuencia. Espere un momento"
	languageMapEs["E_ACCOUNT_LOCKED"] = "Demasiados intentos. Inténtelo de nuevo más tarde"
	languageMapEs["E_ACCOUNT_LOCKED_RETRY"] = "Demasiados intentos. Inténtelo de nuevo en %d segundos"
	languageMapEs["E_RETRIES_REMAINING"] = ". Quedan %d intento(s)"
	languageMapEs["E_VERIFICATION_LIMIT"] = "Demasiadas solicitudes de códigos de verificación. Inténtelo de nuevo más tarde"
	/*common*/
	languageMapEs["E_TOO_LONG"] = "La longitud del texto supera el límite"
	languageMapEs["E_INVALID_PARAM"] = "Parámetro no válido"
	languageMapEs["E_FILE_NOT_FOUND"] = "Archivo no encontrado"
	/*color*/
	languageMapEs["E_INVALID_COLOR"] = "Color no encontrado"
	languageMapEs["E_INVALID_FILE"] = "Suba un archivo correcto"
	languageMapEs["E_INVALID_PICTURE"] = "Suba una imagen (JPG/PNG)"
	languageMapEs["E_AI_RECOMMEND_FAILED"] = "El servicio de recomendación de colores no está disponible temporalmente, inténtelo de nuevo más tarde"
	/*device*/
	languageMapEs["E_NO_ID"] = "Introduzca el ID"
	languageMapEs["E_NO_MAC"] = "Introduzca la dirección MAC"
	languageMapEs["E_NO_INFO"] = "Registro no encontrado"
	languageMapEs["E_NO_DEVICE"] = "Dispositivo no encontrado"
	languageMapEs["E_INVALID_MAC"] = "Introduzca la dirección MAC en el formato correcto"
	/*suggest*/
	languageMapEs["E_NO_CONTENT"] = "Introduzca el contenido"
	languageMapEs["E_NO_SUGGEST"] = "Sugerencia no encontrada"
	/*word*/
	languageMapEs["E_NO_WORD"] = "Palabra sensible no encontrada"
	languageMapEs["E_WORD_EXIST"] = "La palabra sensible ya existe"
	/*follow*/
	languageMapEs["E_FOLLOWED"] = "Ya lo sigue"
	languageMapEs["E_UNFOLLOWED"] = "Ya no lo sigue"
	languageMapEs["E_NOT_FOLLOWING"] = "No lo sigue"
	languageMapEs["E_NO_SELF_FOLLOW"] = "No puede seguirse a sí mismo"
	/*circle*/
	languageMapEs["E_NO_CIRCLE_POST"] = "Publicación no encontrada"
	languageMapEs["E_INVALID_REVIEW_STATUS"] = "Estado de revisión no válido"
	languageMapEs["E_SENSITIVEWORD"] = "Su mensaje contiene información sensible. Para mantener un ambiente de comunicación amigable, ajuste las expresiones relevantes antes de publicar"
	/*topic*/
	languageMapEs["E_NO_TOPIC"] = "Tema no encontrado"
	languageMapEs["E_TOPIC_EXIST"] = "El tema ya existe"
	/*activity*/
	languageMapEs["E_NO_ACTIVITY"] = "Actividad no encontrada"
	languageMapEs["E_NO_TITLE"] = "Introduzca el título"
	languageMapEs["E_NO_TIME"] = "Introduzca la hora"
	languageMapEs["E_ALREADY_JOINED"] = "Ya se unió a esta actividad"
	languageMapEs["E_NOT_JOINED"] = "No se ha unido a esta actividad"
	languageMapEs["E_ALREADY_VOTED"] = "Ya ha votado"
	languageMapEs["E_NOT_VOTED"] = "No ha votado"
	/*comment*/
	languageMapEs["E_NO_COMMENT"] = "Comentario no encontrado"
	/*like*/
	languageMapEs["E_NO_LIKE"] = "Me gusta no encontrado"
	/*collect*/
	languageMapEs["E_NO_COLLECT"] = "No guardado en favoritos"
	/*ota*/
	languageMapEs["E_NO_VERSION"] = "Introduzca la versión"
	/*voice*/
	languageMapEs["E_NO_VOICE_TEXT"] = "Introduzca un comando de voz"
	languageMapEs["E_NO_COLOR_NAME"] = "No se reconoció ningún nombre de color"
	languageMapEs["V_UNKNOWN_COMMAND"] = "Comando de voz desconocido"
	languageMapEs["V_ACTION_HELP"] = "No se pudo reconocer el comando. Pruebe:\n· Cambiar color: cambie el primero a azul medianoche"
	/*checkin*/
	languageMapEs["E_ALREADY_CHECKIN"] = "Ya ha registrado su presencia hoy"
	/*anti_fake*/
	languageMapEs["E_INVALID_PRODUCT_MODEL"] = "Formato de modelo de producto no válido (se requieren 2 caracteres alfanuméricos)"
	languageMapEs["E_INVALID_BATCH"] = "Formato de lote de producción no válido (se requieren 6 caracteres alfanuméricos)"
	languageMapEs["E_INVALID_NUMBER"] = "Formato de número no válido (se requieren 2 caracteres alfanuméricos)"
	languageMapEs["E_INVALID_ANTI_FAKE_CODE"] = "Formato de código antifalsificación no válido (se requieren 24 caracteres alfanuméricos)"
	languageMapEs["E_GENERATE_CODE_FAILED"] = "Error al generar el código antifalsificación"
	languageMapEs["E_ANTI_FAKE_CODE_ACTIVATED"] = "El código antifalsificación ya está activado"
	languageMapEs["E_NO_ANTI_FAKE_CODE"] = "El código antifalsificación no existe"
	languageMapEs["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "Este código antifalsificación ya ha sido verificado por otro usuario"
	languageMapEs["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "El código antifalsificación no está activado, no se puede verificar"
	/*mail*/
	languageMapEs["MAIL_VERIFY_SUBJECT"] = "Código de verificación de TintaShift"
	languageMapEs["MAIL_VERIFY_TITLE"] = "Bienvenido a TintaShift"
	languageMapEs["MAIL_VERIFY_DESC"] = "Hola, está realizando una verificación de seguridad de la cuenta. Su código es:"
	languageMapEs["MAIL_VERIFY_SECURITY_TITLE"] = "Consejo de seguridad: "
	languageMapEs["MAIL_VERIFY_SECURITY_DESC"] = "El personal de TintaShift nunca le pedirá este código. No lo comparta con nadie."
	languageMapEs["MAIL_VERIFY_IGNORE"] = "Si no solicitó este código, ignore este correo electrónico."
	languageMapEs["MAIL_COPYRIGHT"] = "© 2026 TintaShift. Todos los derechos reservados."
	/*social login*/
	languageMapEs["E_NO_GOOGLE_ACCESS_TOKEN"] = "Se requiere el token de acceso de Google"
	languageMapEs["E_GOOGLE_NOT_CONFIGURED"] = "El inicio de sesión con Google no está configurado en el servidor"
	languageMapEs["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Token de acceso de Google no válido o caducado"
	languageMapEs["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "El correo de Google no está verificado y no se puede usar para iniciar sesión"
	languageMapEs["E_GOOGLE_EMAIL_BOUND_OTHER"] = "Este correo ya está vinculado a otra cuenta de Google"
	languageMapEs["E_NO_APPLE_ID_TOKEN"] = "Se requiere el identity token de Apple"
	languageMapEs["E_APPLE_NOT_CONFIGURED"] = "El inicio de sesión con Apple no está configurado en el servidor"
	languageMapEs["E_APPLE_ID_TOKEN_INVALID"] = "Identity token de Apple no válido o caducado"
	languageMapEs["E_APPLE_EMAIL_NOT_VERIFIED"] = "El correo de Apple no está verificado y no se puede usar para iniciar sesión"
	languageMapEs["E_APPLE_EMAIL_BOUND_OTHER"] = "Este correo ya está vinculado a otra cuenta de Apple"
	languageMapEs["E_SOCIAL_NOT_CONFIGURED"] = "El secreto de firma del inicio de sesión social no está configurado en el servidor"
	languageMapEs["E_NO_PROVIDER"] = "se requiere provider"
	languageMapEs["E_SOCIAL_PROVIDER_INVALID"] = "provider de inicio de sesión social no admitido"
	languageMapEs["E_NO_OPENID"] = "se requiere openid"
	languageMapEs["E_NO_NONCE"] = "se requiere nonce"
	languageMapEs["E_NO_TIMESTAMP"] = "se requiere timestamp"
	languageMapEs["E_NO_SIGNATURE"] = "se requiere signature"
	languageMapEs["E_NO_DEVICE_ID"] = "se requiere device_id"
	languageMapEs["E_REQUEST_EXPIRED"] = "la solicitud ha caducado, inténtelo de nuevo"
	languageMapEs["E_CHALLENGE_INVALID"] = "challenge no válido"
	languageMapEs["E_CHALLENGE_EXPIRED"] = "challenge caducado"
	languageMapEs["E_CHALLENGE_USED"] = "challenge ya utilizado"
	languageMapEs["E_SIGNATURE_INVALID"] = "fallo en la verificación de la firma"
	languageMapEs["E_SOCIAL_EMAIL_BOUND_OTHER"] = "Este correo ya está vinculado a otra cuenta social"
	languageMapEs["E_INVALID_VIDEO"] = "Suba un vídeo válido (MP4/MOV)"
	registerLocale("es-ES", languageMapEs)
}
