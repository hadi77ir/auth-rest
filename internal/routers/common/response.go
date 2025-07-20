package common

type SuccessResponse struct {
	Success bool `json:"success" example:"true"`
}
type ErrorResponse struct {
	Error string `json:"error" example:"this is an example error"`
}
