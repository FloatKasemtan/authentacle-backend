package request

type UserRequest struct {
	Username string `json:"username"binding:"required"`
	Email    string `json:"email"binding:"required"`
	Password string `json:"password"binding:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username"binding:"required"`
	Password string `json:"password"binding:"required"`
}
