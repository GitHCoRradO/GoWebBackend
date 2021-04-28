package response

type LoginResponse struct {
	Status string `json:"status"`
	Err string `json:"error"`
}
