package model

import "github.com/gogf/gf/v2/util/gmeta"

type AdminLoginInput struct {
	Username   string
	Password   string
	CaptchaStr string
	CaptchaId  string
}

type AdminGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

type AdminListItem struct {
	gmeta.Meta  `orm:"table:admin"`
	Id          int               `json:"id"`                //ID
	Username    string            `json:"username"`          //用户名
	Name        string            `json:"Name"`              //姓名
	Tel         string            `json:"tel"`               //手机号
	Email       string            `json:"email"`             //邮箱
	Status      uint              `json:"status"`            //状态
	CreatedAt   string            `json:"created_at"`        //创建时间
	UpdatedAt   string            `json:"updated_at"`        //修改时间
	RoleAccount []RoleAccountItem `orm:"with:account_id=id"` //用户角色id们
}

type AdminGetListOutput struct {
	List  []AdminListItem `json:"list" description:"列表"`
	Page  int             `json:"page" description:"分页码"`
	Size  int             `json:"size" description:"分页数量"`
	Total int             `json:"total" description:"数据总数"`
}

type AdminSession struct {
	Id       int    `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
