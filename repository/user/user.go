package user

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsVerify bool   `json:"isVerify"`
}

type Repository interface {
	SignUp(username string, email string, password string) (string, error)
	SignIn(username string, password string) (string, error)
	GetById(userId string) (*User, error)
	SendVerificationForm(userId string) error
}
