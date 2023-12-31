// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsChannelDao is the data access object for table cms_channel.
type CmsChannelDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CmsChannelColumns // columns contains all the column names of Table for convenient usage.
}

// CmsChannelColumns defines and stores column names for table cms_channel.
type CmsChannelColumns struct {
	Id             string // 频道ID
	Pid            string // 父级ID
	Tid            string // 顶级id(已经是顶级则为自己)
	ChildrenIds    string // 所有子栏目id们(不包含自己)
	Level          string // 分类层次
	Name           string // 名称
	Thumb          string // 缩略图
	Sort           string // 排名
	Status         string // 状态(0:停用;1:启用;)
	Type           string // 类型(1:频道;2:单页;3:链接)
	LinkUrl        string // 链接地址
	LinkTrigger    string // 链接打开方式(0:当前窗口打开;1:新窗口打开;)
	ListRouter     string // 列表页路由
	DetailRouter   string // 详情页路由
	ListTemplate   string // 列表页模板
	DetailTemplate string // 详情页模板
	Description    string // 频道描述
	Model          string // 模型
	CreatedAt      string // 创建时间
	UpdatedAt      string // 修改时间
	DeletedAt      string // 删除时间
}

// cmsChannelColumns holds the columns for table cms_channel.
var cmsChannelColumns = CmsChannelColumns{
	Id:             "id",
	Pid:            "pid",
	Tid:            "tid",
	ChildrenIds:    "children_ids",
	Level:          "level",
	Name:           "name",
	Thumb:          "thumb",
	Sort:           "sort",
	Status:         "status",
	Type:           "type",
	LinkUrl:        "link_url",
	LinkTrigger:    "link_trigger",
	ListRouter:     "list_router",
	DetailRouter:   "detail_router",
	ListTemplate:   "list_template",
	DetailTemplate: "detail_template",
	Description:    "description",
	Model:          "model",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewCmsChannelDao creates and returns a new DAO object for table data access.
func NewCmsChannelDao() *CmsChannelDao {
	return &CmsChannelDao{
		group:   "default",
		table:   "cms_channel",
		columns: cmsChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsChannelDao) Columns() CmsChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
