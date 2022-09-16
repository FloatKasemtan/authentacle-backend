package request

type ApplicationRequest struct {
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	UnhashKey string `json:"unhash_key"`
}
