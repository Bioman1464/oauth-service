package responses

type AuthorizationResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
