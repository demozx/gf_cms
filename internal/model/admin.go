package model

type AdminLoginInput struct {
	Username   string
	Password   string
	CaptchaStr string
	CaptchaId  string
}
