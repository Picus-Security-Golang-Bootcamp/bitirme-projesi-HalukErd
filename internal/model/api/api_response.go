package api

type APIResponse struct {
	Code    int64       `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
	Message string      `json:"message,omitempty"`
}
