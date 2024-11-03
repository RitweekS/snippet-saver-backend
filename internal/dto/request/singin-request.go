package request

type SingInRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}