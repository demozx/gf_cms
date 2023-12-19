package captcha

import (
	"fmt"
	"gf_cms/internal/service"

	"github.com/mojocn/base64Captcha"
)

var (
	store      = base64Captcha.DefaultMemStore
	insCaptcha = sCaptcha{}
)

func init() {
	service.RegisterCaptcha(New())
}

func New() *sCaptcha {
	return &sCaptcha{}
}

func Captcha() *sCaptcha {
	return &insCaptcha
}

type sCaptcha struct{}

// Get GetCaptcha 获取验证码
func (s *sCaptcha) Get() (string, string) {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)

	// 获取
	id, b64s, _, err := c.Generate()
	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return "", ""
	}
	return id, b64s
}

// Verify 验证验证码
func (s *sCaptcha) Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	// 同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}
