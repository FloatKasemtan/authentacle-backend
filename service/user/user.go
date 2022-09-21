package user

type Response struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsVerify bool   `json:"isVerify"`
}

type Service interface {
	SignUp(username string, email string, password string) (string, error)
	SignIn(username string, password string) (string, error)
	GetUser(userId string) (*Response, error)
	SendVerificationForm(id string, email string) error
}
