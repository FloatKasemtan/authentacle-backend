package user

type Response struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsVerify bool   `json:"isVerify"`
}

type Service interface {
	SignUp(username string, email string, password string) (*string, *string, *string, error)
	SignIn(username string, password string) (*string, *bool, *string, *string, error)
	GetUser(userId string) (*Response, error)
	Verify(id string, role int8, otp string) (*string, error)
	CheckOTP(id string, role int8, otp string) (*string, error)
}
