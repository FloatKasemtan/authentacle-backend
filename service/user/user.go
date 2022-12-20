package user

type Response struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsVerify bool   `json:"isVerify"`
}

type Service interface {
	SignUp(username string, email string, password string) (*string, *string, error)
	SignIn(username string, password string, otp string) (*string, error)
	GetUser(userId string) (*Response, error)
	Verify(id string, role int8, otp string) (*string, error)
}
