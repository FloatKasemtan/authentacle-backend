package response

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(arg1 any, arg2 ...any) SuccessResponse {
	if message, ok := arg1.(string); ok {
		if len(arg2) == 0 {
			return SuccessResponse{
				Success: true,
				Message: message,
			}
		} else if message2, ok := arg2[0].(string); ok {
			return SuccessResponse{
				Success: true,
				Message: message2,
				Data:    message,
			}
		} else {
			return SuccessResponse{
				Success: true,
				Message: message,
				Data:    arg2,
			}
		}
	}
	return SuccessResponse{
		Success: true,
		Data:    arg1,
	}
}
