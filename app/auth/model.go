package auth

type responseAccessToken struct {
	IsSuccess   bool   `json:"is_success"`
	AccessToken string `json:"access_token"`
}

type requestRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	PhoneNo  string `json:"phone_no"`
	Email    string `json:"email"`
}
type requestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	Message string `json:"message"`
}
