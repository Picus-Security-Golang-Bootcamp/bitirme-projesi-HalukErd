package api

type SignupResponse struct {
	Id    *string  `json:"id"`
	Email *string  `json:"email"`
	Roles []string `json:"roles"`
}
