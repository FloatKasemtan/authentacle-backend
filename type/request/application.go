package request

type ApplicationRequest struct {
	Name      string `json:"name"binding:"required"`
	Logo      string `json:"logo"binding:"required"`
	UnhashKey string `json:"unhash_key"binding:"required"`
}

type GetApplicationRequest struct {
	Id string `json:"id"binding:"required"`
}
