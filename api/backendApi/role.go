package backendApi

import "github.com/gogf/gf/v2/frame/g"

type RoleStatusReq struct {
	g.Meta `tags:"Backend" method:"all" summary:"修改启动状态"`
	Id     int `p:"id" name:"id" brief:"角色ID" des:"角色ID"  arg:"true" v:"required#请输入角色ID"`
}
type RoleStatusRes struct {
}

type RoleDeleteReq struct {
	g.Meta `tags:"Backend" method:"all" summary:"删除角色"`
	Id     int `p:"id" name:"id" brief:"角色ID" des:"角色ID"  arg:"true" v:"required#请输入角色ID"`
}
type RoleDeleteRes struct {
}

type RoleDeleteBatchReq struct {
	g.Meta `tags:"Backend" method:"all" summary:"批量删除角色"`
	Ids    []string `p:"ids" name:"ids" brief:"角色ID们" des:"角色ID们"  arg:"true" v:"required#请输入角色ID们"`
}
type RoleDeleteBatchRes struct {
}
