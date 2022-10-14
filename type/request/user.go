package request

type UserRequest struct {
	Username string `json:"username"validate:"required"`
	Email    string `json:"email"validate:"required"`
	Password string `json:"password"validate:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username"validate:"required"`
	Password string `json:"password"validate:"required"`
	Otp      string `json:"otp"validate:"required"`
}
