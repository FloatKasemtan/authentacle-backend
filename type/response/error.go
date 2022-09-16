package response

type ErrorResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type GenericError struct {
	Code    string
	Message string
	Err     error
}

func (v *GenericError) Error() string {
	return v.Message
}
