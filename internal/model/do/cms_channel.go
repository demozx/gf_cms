// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsChannel is the golang structure of table cms_channel for DAO operations like Where/Data.
type CmsChannel struct {
	g.Meta         `orm:"table:cms_channel, do:true"`
	Id             interface{} // 频道ID
	Pid            interface{} // 父级ID
	Tid            interface{} // 顶级id(已经是顶级则为自己)
	ChildrenIds    interface{} // 所有子栏目id们(不包含自己)
	Level          interface{} // 分类层次
	Name           interface{} // 名称
	Thumb          interface{} // 缩略图
	Sort           interface{} // 排名
	Status         interface{} // 状态(0:停用;1:启用;)
	Type           interface{} // 类型(1:频道;2:单页;3:链接)
	LinkUrl        interface{} // 链接地址
	LinkTrigger    interface{} // 链接打开方式(0:当前窗口打开;1:新窗口打开;)
	ListRouter     interface{} // 列表页路由
	DetailRouter   interface{} // 详情页路由
	ListTemplate   interface{} // 列表页模板
	DetailTemplate interface{} // 详情页模板
	Description    interface{} // 频道描述
	Model          interface{} // 模型
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 修改时间
	DeletedAt      *gtime.Time // 删除时间
}
