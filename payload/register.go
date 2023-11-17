package payload

type InfoRegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CodeRegisterPayload struct {
	IdTemporaryCredential uint   `json:"idTemporaryCredential"`
	Code                  string `json:"code"`
}
