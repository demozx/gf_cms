package backendApi

import "github.com/gogf/gf/v2/frame/g"

type CaptchaGetApiReq struct {
	g.Meta `tags:"BackendApi" method:"get" summary:"验证码"`
}
type CaptchaGetApiRes struct {
	Id   string `json:"id" "dc": "验证码id"`
	B64s string `json:"b64s" "dc": "验证码图"`
}
