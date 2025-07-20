package auth

type LoginRequest struct {
	Phone string `json:"phone" example:"09123456789"`
	OTP   string `json:"otp" example:"123456"`
}
type LoginResponse struct {
	Token string `json:"token" example:"xxxxx.yyyyy.zzzzz"`
}
type LoginInitRequest struct {
	Phone string `json:"phone" example:"09123456789"`
}
