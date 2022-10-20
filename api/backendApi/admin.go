package backendApi

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type AdminLoginReq struct {
	g.Meta     `tags:"BackendApi" method:"post" summary:"后台登录"`
	Username   string `name:"username" brief:"用户名" des:"用户名"  arg:"true" v:"required#请输入用户名"`
	Password   string `name:"password" brief:"密码" des:"密码"  arg:"true" v:"required#请输入密码"`
	CaptchaStr string `name:"captcha_str" brief:"验证码字符串" des:"验证码字符串"  arg:"true" v:"required#请输入验证码"`
	CaptchaId  string `name:"captcha_id" brief:"验证码id" des:"验证码id"  arg:"true" v:"required#请输入验证码id"`
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

// AdminAddReq 添加管理员
type AdminAddReq struct {
	g.Meta     `tags:"Backend" method:"post" summary:"添加管理员"`
	Username   string `name:"username" brief:"用户名" des:"用户名"  arg:"true" v:"required#请输入用户名"`
	Name       string `name:"name" brief:"姓名" des:"姓名"  arg:"true" v:"required#请输入姓名"`
	Tel        string `name:"tel" brief:"手机" des:"手机"  arg:"true" v:"required|phone#请输入手机|手机格式错误"`
	Email      string `name:"email" brief:"邮箱" des:"邮箱"  arg:"true" v:"required|email#请输入邮箱|邮箱格式错误"`
	Role       g.Map  `name:"role_ids" brief:"角色ids" des:"角色ids"  arg:"true" v:"required#请选择角色"`
	Status     int    `name:"status" brief:"状态" des:"状态"  arg:"true" v:"in:0,1#状态不合法"`
	Password   string `name:"password" brief:"密码" des:"密码"  arg:"true" v:"required|length:6,16#请输入密码|密码必须是6-16位"`
	RePassword string `name:"re_password" brief:"确认密码" des:"确认密码"  arg:"true" v:"required|same:Password#请输入确认密码|确认密码错误"`
}
type AdminAddRes struct{}

// AdminEditReq 修改管理员
type AdminEditReq struct {
	g.Meta     `tags:"Backend" method:"post" summary:"修改管理员"`
	Id         int    `name:"id" brief:"管理员ID" des:"管理员ID"  arg:"true" v:"required#请输入管理员ID"`
	Username   string `name:"username" brief:"用户名" des:"用户名"  arg:"true" v:"required#请输入用户名"`
	Name       string `name:"name" brief:"姓名" des:"姓名"  arg:"true" v:"required#请输入姓名"`
	Tel        string `name:"tel" brief:"手机" des:"手机"  arg:"true" v:"required|phone#请输入手机|手机格式错误"`
	Email      string `name:"email" brief:"邮箱" des:"邮箱"  arg:"true" v:"required|email#请输入邮箱|邮箱格式错误"`
	Role       g.Map  `name:"role_ids" brief:"角色ids" des:"角色ids"  arg:"true" v:"required#请选择角色"`
	Status     int    `name:"status" brief:"状态" des:"状态"  arg:"true" v:"in:0,1#状态不合法"`
	Password   string `name:"password" brief:"密码" des:"密码"  arg:"true" v:"length:6,16#密码必须是6-16位"`
	RePassword string `name:"re_password" brief:"确认密码" des:"确认密码"  arg:"true" v:"same:Password#确认密码错误"`
}
type AdminEditRes struct{}

type AdminStatusReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"修改启动状态"`
	Id     int `name:"id" brief:"管理员ID" des:"管理员ID"  arg:"true" v:"required#请输入管理员ID"`
}
type AdminStatusRes struct {
}

type AdminDeleteReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"删除管理员"`
	Id     int `name:"id" brief:"管理员ID" des:"管理员ID"  arg:"true" v:"required#请输入管理员ID"`
}
type AdminDeleteRes struct {
}

type AdminDeleteBatchReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"批量删除管理员"`
	Ids    []string `name:"ids" brief:"管理员ID们" des:"管理员ID们"  arg:"true" v:"required#请输入管理员ID们"`
}
type AdminDeleteBatchRes struct {
}
