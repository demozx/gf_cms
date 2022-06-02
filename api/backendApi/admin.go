package backendApi

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type AdminLoginReq struct {
	g.Meta     `tags:"BackendApi" method:"post" summary:"后台登录"`
	Username   string `p:"username" name:"username" brief:"用户名" des:"用户名"  arg:"true" v:"required#请输入用户名"`
	Password   string `p:"password" name:"password" brief:"密码" des:"密码"  arg:"true" v:"required#请输入密码"`
	CaptchaStr string `p:"captcha_str" name:"captcha_str" brief:"验证码字符串" des:"验证码字符串"  arg:"true" v:"required#请输入验证码"`
	CaptchaId  string `p:"captcha_id" name:"captcha_id" brief:"验证码id" des:"验证码id"  arg:"true" v:"required#请输入验证码id"`
}
type AdminLoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type AdminLogoutReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"后台退出"`
}
type AdminLogoutRes struct{}

type AdminClearCacheReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"清除缓存"`
}
type AdminClearCacheRes struct{}
