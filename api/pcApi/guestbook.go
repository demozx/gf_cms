package pcApi

import "github.com/gogf/gf/v2/frame/g"

type GuestbookReq struct {
	g.Meta  `tags:"Pc" method:"post" summary:"pc提交留言"`
	Name    string `json:"name" dc:"联系人" v:"required#联系人不能为空"`
	Tel     string `json:"tel" dc:"手机" v:"required|phone#手机不能为空|手机号格式错误"`
	Email   string `json:"email" dc:"邮箱" v:"required|email#邮箱不能为空|邮箱格式错误"`
	Content string `json:"content" dc:"内容" v:"required#内容不能为空"`
}
type GuestbookRes struct{}
