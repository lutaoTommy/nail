package language

var languageMapPt map[string]string

/*初始化*/
func initLanguagePtpt() {
	languageMapPt = make(map[string]string)
	/*user*/
	languageMapPt["E_NO_IV"] = "Insira o vetor de inicialização"
	languageMapPt["E_NO_CODE"] = "Insira o código de autorização"
	languageMapPt["E_NO_DATA"] = "Insira os dados"
	languageMapPt["E_NO_PHONE"] = "Insira o número de telefone"
	languageMapPt["E_NO_PWD"] = "Insira a palavra-passe"
	languageMapPt["E_INVALID_USER"] = "Utilizador não encontrado"
	languageMapPt["E_INVALID_PWD"] = "Palavra-passe incorreta"
	languageMapPt["E_NO_NAME"] = "Insira o nome"
	languageMapPt["E_USER_EXIST"] = "A conta já existe"
	languageMapPt["E_NO_TOKEN"] = "Inicie sessão primeiro"
	languageMapPt["E_NO_AUTH"] = "Sem autorização"
	languageMapPt["E_NO_AVATAR"] = "Avatar não encontrado"
	languageMapPt["E_NO_EMAIL"] = "Insira o e-mail"
	languageMapPt["E_NO_CERT"] = "Insira o código de verificação"
	languageMapPt["E_INVALID_EMAIL"] = "Formato de e-mail inválido"
	languageMapPt["E_DATA_TOO_SHORT"] = "Os dados encriptados são demasiado curtos"
	languageMapPt["E_INVALID_IV"] = "Vetor de inicialização inválido"
	languageMapPt["E_INVALID_SESSION_KEY"] = "Chave de sessão inválida"
	languageMapPt["E_DECRYPT_FAIL"] = "Falha ao desencriptar; certifique-se de que o código e o número de telefone são da mesma sessão"
	languageMapPt["E_INVALID_PHONE"] = "Formato de número de telefone inválido"
	languageMapPt["VERIFICATION_CODE"] = "código de verificação"
	languageMapPt["DO_NOT_TELL"] = "Não o diga a ninguém"
	languageMapPt["E_CERT_ERR"] = "Código de verificação incorreto"
	languageMapPt["E_CERT_EXPIRED"] = "O código de verificação expirou. Solicite um novo"
	languageMapPt["E_CERT_FIRST"] = "Solicite primeiro o código de verificação"
	languageMapPt["E_VERIFICATION_QUICKLY"] = "Está a solicitar códigos de verificação com muita frequência. Aguarde um momento"
	languageMapPt["E_ACCOUNT_LOCKED"] = "Demasiadas tentativas. Tente novamente mais tarde"
	languageMapPt["E_ACCOUNT_LOCKED_RETRY"] = "Demasiadas tentativas. Tente novamente em %d segundos"
	languageMapPt["E_RETRIES_REMAINING"] = ". Restam %d tentativa(s)"
	languageMapPt["E_VERIFICATION_LIMIT"] = "Demasiados pedidos de códigos de verificação. Tente novamente mais tarde"
	/*common*/
	languageMapPt["E_TOO_LONG"] = "O comprimento do texto excede o limite"
	languageMapPt["E_INVALID_PARAM"] = "Parâmetro inválido"
	languageMapPt["E_FILE_NOT_FOUND"] = "Ficheiro não encontrado"
	/*color*/
	languageMapPt["E_INVALID_COLOR"] = "Cor não encontrada"
	languageMapPt["E_INVALID_FILE"] = "Carregue um ficheiro correto"
	languageMapPt["E_INVALID_PICTURE"] = "Carregue uma imagem (JPG/PNG)"
	languageMapPt["E_AI_RECOMMEND_FAILED"] = "O serviço de recomendação de cores está temporariamente indisponível, tente novamente mais tarde"
	/*device*/
	languageMapPt["E_NO_ID"] = "Insira o ID"
	languageMapPt["E_NO_MAC"] = "Insira o endereço MAC"
	languageMapPt["E_NO_INFO"] = "Registo não encontrado"
	languageMapPt["E_NO_DEVICE"] = "Dispositivo não encontrado"
	languageMapPt["E_INVALID_MAC"] = "Insira o endereço MAC no formato correto"
	/*suggest*/
	languageMapPt["E_NO_CONTENT"] = "Insira o conteúdo"
	languageMapPt["E_NO_SUGGEST"] = "Sugestão não encontrada"
	/*word*/
	languageMapPt["E_NO_WORD"] = "Palavra sensível não encontrada"
	languageMapPt["E_WORD_EXIST"] = "A palavra sensível já existe"
	/*follow*/
	languageMapPt["E_FOLLOWED"] = "Já segue"
	languageMapPt["E_UNFOLLOWED"] = "Deixou de seguir"
	languageMapPt["E_NOT_FOLLOWING"] = "Não segue"
	languageMapPt["E_NO_SELF_FOLLOW"] = "Não é possível seguir a si próprio"
	/*circle*/
	languageMapPt["E_NO_CIRCLE_POST"] = "Publicação não encontrada"
	languageMapPt["E_INVALID_REVIEW_STATUS"] = "Estado de revisão inválido"
	languageMapPt["E_SENSITIVEWORD"] = "A sua mensagem contém informações sensíveis. Para manter um ambiente de comunicação amigável, ajuste as expressões relevantes antes de publicar"
	/*topic*/
	languageMapPt["E_NO_TOPIC"] = "Tópico não encontrado"
	languageMapPt["E_TOPIC_EXIST"] = "O tópico já existe"
	/*activity*/
	languageMapPt["E_NO_ACTIVITY"] = "Atividade não encontrada"
	languageMapPt["E_NO_TITLE"] = "Insira o título"
	languageMapPt["E_NO_TIME"] = "Insira a hora"
	languageMapPt["E_ALREADY_JOINED"] = "Já participou nesta atividade"
	languageMapPt["E_NOT_JOINED"] = "Não participou nesta atividade"
	languageMapPt["E_ALREADY_VOTED"] = "Já votou"
	languageMapPt["E_NOT_VOTED"] = "Não votou"
	/*comment*/
	languageMapPt["E_NO_COMMENT"] = "Comentário não encontrado"
	/*like*/
	languageMapPt["E_NO_LIKE"] = "Gosto não encontrado"
	/*collect*/
	languageMapPt["E_NO_COLLECT"] = "Não guardado"
	/*ota*/
	languageMapPt["E_NO_VERSION"] = "Insira a versão"
	/*voice*/
	languageMapPt["E_NO_VOICE_TEXT"] = "Insira um comando de voz"
	languageMapPt["E_NO_COLOR_NAME"] = "Nenhum nome de cor reconhecido"
	languageMapPt["V_UNKNOWN_COMMAND"] = "Comando de voz desconhecido"
	languageMapPt["V_ACTION_HELP"] = "Não foi possível reconhecer o comando. Experimente:\n· Mudar cor: mudar o primeiro para azul meia-noite"
	/*checkin*/
	languageMapPt["E_ALREADY_CHECKIN"] = "Já fez check-in hoje"
	/*anti_fake*/
	languageMapPt["E_INVALID_PRODUCT_MODEL"] = "Formato de modelo de produto inválido (necessários 2 caracteres alfanuméricos)"
	languageMapPt["E_INVALID_BATCH"] = "Formato de lote de produção inválido (necessários 6 caracteres alfanuméricos)"
	languageMapPt["E_INVALID_NUMBER"] = "Formato de número inválido (necessários 2 caracteres alfanuméricos)"
	languageMapPt["E_INVALID_ANTI_FAKE_CODE"] = "Formato de código antifalsificação inválido (necessários 24 caracteres alfanuméricos)"
	languageMapPt["E_GENERATE_CODE_FAILED"] = "Falha ao gerar o código antifalsificação"
	languageMapPt["E_ANTI_FAKE_CODE_ACTIVATED"] = "Código antifalsificação já ativado"
	languageMapPt["E_NO_ANTI_FAKE_CODE"] = "Código antifalsificação não encontrado"
	languageMapPt["E_ANTI_FAKE_CODE_VERIFIED_BY_OTHER"] = "Este código antifalsificação já foi verificado por outro utilizador"
	languageMapPt["E_ANTI_FAKE_CODE_NOT_ACTIVATED"] = "Código antifalsificação não ativado, não é possível verificar"
	/*mail*/
	languageMapPt["MAIL_VERIFY_SUBJECT"] = "Código de verificação TintaShift"
	languageMapPt["MAIL_VERIFY_TITLE"] = "Bem-vindo ao TintaShift"
	languageMapPt["MAIL_VERIFY_DESC"] = "Olá, está a realizar uma verificação de segurança da conta. O seu código é:"
	languageMapPt["MAIL_VERIFY_SECURITY_TITLE"] = "Dica de segurança: "
	languageMapPt["MAIL_VERIFY_SECURITY_DESC"] = "A equipa da TintaShift nunca lhe pedirá este código. Não o partilhe com ninguém."
	languageMapPt["MAIL_VERIFY_IGNORE"] = "Se não solicitou este código, ignore este e-mail."
	languageMapPt["MAIL_COPYRIGHT"] = "© 2026 TintaShift. Todos os direitos reservados."
	/*social login*/
	languageMapPt["E_NO_GOOGLE_ACCESS_TOKEN"] = "É necessário o token de acesso Google"
	languageMapPt["E_GOOGLE_NOT_CONFIGURED"] = "O início de sessão com Google não está configurado no servidor"
	languageMapPt["E_GOOGLE_ACCESS_TOKEN_INVALID"] = "Token de acesso Google inválido ou expirado"
	languageMapPt["E_GOOGLE_EMAIL_NOT_VERIFIED"] = "O e-mail Google não está verificado e não pode ser usado para iniciar sessão"
	languageMapPt["E_GOOGLE_EMAIL_BOUND_OTHER"] = "Este e-mail já está associado a outra conta Google"
	languageMapPt["E_NO_APPLE_ID_TOKEN"] = "É necessário o token de identidade Apple"
	languageMapPt["E_APPLE_NOT_CONFIGURED"] = "O início de sessão com Apple não está configurado no servidor"
	languageMapPt["E_APPLE_ID_TOKEN_INVALID"] = "Token de identidade Apple inválido ou expirado"
	languageMapPt["E_APPLE_EMAIL_NOT_VERIFIED"] = "O e-mail Apple não está verificado e não pode ser usado para iniciar sessão"
	languageMapPt["E_APPLE_EMAIL_BOUND_OTHER"] = "Este e-mail já está associado a outra conta Apple"
	languageMapPt["E_SOCIAL_NOT_CONFIGURED"] = "O segredo de assinatura do início de sessão social não está configurado no servidor"
	languageMapPt["E_NO_PROVIDER"] = "provider é necessário"
	languageMapPt["E_SOCIAL_PROVIDER_INVALID"] = "provider de início de sessão social não suportado"
	languageMapPt["E_NO_OPENID"] = "openid é necessário"
	languageMapPt["E_NO_NONCE"] = "nonce é necessário"
	languageMapPt["E_NO_TIMESTAMP"] = "timestamp é necessário"
	languageMapPt["E_NO_SIGNATURE"] = "signature é necessário"
	languageMapPt["E_NO_DEVICE_ID"] = "device_id é necessário"
	languageMapPt["E_REQUEST_EXPIRED"] = "o pedido expirou, tente novamente"
	languageMapPt["E_CHALLENGE_INVALID"] = "challenge inválido"
	languageMapPt["E_CHALLENGE_EXPIRED"] = "challenge expirado"
	languageMapPt["E_CHALLENGE_USED"] = "challenge já utilizado"
	languageMapPt["E_SIGNATURE_INVALID"] = "falha na verificação da assinatura"
	languageMapPt["E_SOCIAL_EMAIL_BOUND_OTHER"] = "Este e-mail já está associado a outra conta social"
	languageMapPt["E_INVALID_VIDEO"] = "Carregue um vídeo válido (MP4/MOV)"
	registerLocale("pt-PT", languageMapPt)
}
