package request

type ApplicationRequest struct {
	Name      string `json:"name"validate`
	Logo      string `json:"logo"`
	UnhashKey string `json:"unhash_key"`
}

type GetApplicationRequest struct {
	Id string `json:"id"`
}
