package backendApi

import "github.com/gogf/gf/v2/frame/g"

type RoleStatusReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"修改启动状态"`
	Id     int `name:"id" brief:"角色ID" des:"角色ID"  arg:"true" v:"required#请输入角色ID"`
}
type RoleStatusRes struct {
}

type RoleDeleteReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"删除角色"`
	Id     int `name:"id" brief:"角色ID" des:"角色ID"  arg:"true" v:"required#请输入角色ID"`
}
type RoleDeleteRes struct {
}

type RoleDeleteBatchReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"批量删除角色"`
	Ids    []string `name:"ids" brief:"角色ID们" des:"角色ID们"  arg:"true" v:"required#请输入角色ID们"`
}
type RoleDeleteBatchRes struct {
}

type RoleAddReq struct {
	g.Meta      `tags:"Backend" method:"post" summary:"增加角色"`
	Title       string                            `name:"title" brief:"角色名称" des:"角色名称"  arg:"true" v:"required#请输入角色名称"`
	Description string                            `name:"description" brief:"角色描述" des:"角色描述"  arg:"true" v:"required#请输入角色描述"`
	Status      int                               `name:"status" brief:"状态" des:"状态"  arg:"true" v:"in:0,1#状态不合法"`
	Rules       map[string]map[string]interface{} `name:"rules" brief:"权限ID们" des:"权限ID们"  arg:"true" v:"required#请输入权限ID们"`
}
type RoleAddRes struct {
}

type RoleEditReq struct {
	g.Meta      `tags:"Backend" method:"post" summary:"编辑角色"`
	Id          int                               `name:"id" brief:"角色ID" des:"角色ID"  arg:"true" v:"required#请输入角色ID"`
	Title       string                            `name:"title" brief:"角色名称" des:"角色名称"  arg:"true" v:"required#请输入角色名称"`
	Description string                            `name:"description" brief:"角色描述" des:"角色描述"  arg:"true" v:"required#请输入角色描述"`
	Status      int                               `name:"status" brief:"状态" des:"状态"  arg:"true" v:"in:0,1#状态不合法"`
	Rules       map[string]map[string]interface{} `name:"rules" brief:"权限ID们" des:"权限ID们"  arg:"true" v:"required#请输入权限ID们"`
}
type RoleEditRes struct {
}
