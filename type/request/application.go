package request

type ApplicationRequest struct {
	Name      string `json:"name"validate:"required"`
	Logo      string `json:"logo"validate:"required"`
	UnhashKey string `json:"unhash_key"validate:"required"`
}

type GetApplicationRequest struct {
	Id string `json:"id"validate:"required"`
}
