// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsChannel is the golang structure for table cms_channel.
type CmsChannel struct {
	Id             uint        `json:"id"             ` // 频道ID
	Pid            int         `json:"pid"            ` // 父级ID
	Tid            int         `json:"tid"            ` // 顶级id(已经是顶级则为自己)
	ChildrenIds    string      `json:"childrenIds"    ` // 所有子栏目id们(不包含自己)
	Level          int         `json:"level"          ` // 分类层次
	Name           string      `json:"name"           ` // 名称
	Thumb          string      `json:"thumb"          ` // 缩略图
	Sort           int         `json:"sort"           ` // 排名
	Status         int         `json:"status"         ` // 状态(0:停用;1:启用;)
	Type           int         `json:"type"           ` // 类型(1:频道;2:单页;3:链接)
	LinkUrl        string      `json:"linkUrl"        ` // 链接地址
	LinkTrigger    int         `json:"linkTrigger"    ` // 链接打开方式(0:当前窗口打开;1:新窗口打开;)
	ListRouter     string      `json:"listRouter"     ` // 列表页路由
	DetailRouter   string      `json:"detailRouter"   ` // 详情页路由
	ListTemplate   string      `json:"listTemplate"   ` // 列表页模板
	DetailTemplate string      `json:"detailTemplate" ` // 详情页模板
	Description    string      `json:"description"    ` // 频道描述
	Model          string      `json:"model"          ` // 模型
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 修改时间
	DeletedAt      *gtime.Time `json:"deletedAt"      ` // 删除时间
}
