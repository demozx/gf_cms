// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsImageDao is the data access object for table cms_image.
type CmsImageDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CmsImageColumns // columns contains all the column names of Table for convenient usage.
}

// CmsImageColumns defines and stores column names for table cms_image.
type CmsImageColumns struct {
	Id          string // 图片id
	Title       string // 标题
	ChannelId   string // 所属栏目id
	Images      string // 图片们
	Description string // 图片描述
	Flag        string // 属性(r:推荐,t:置顶)
	Status      string // 审核状态(1:启用,0:停用)
	ClickNum    string // 点击数
	Sort        string // 排序
	CreatedAt   string // 发布时间
	UpdatedAt   string // 编辑时间
	DeletedAt   string // 删除时间
}

// cmsImageColumns holds the columns for table cms_image.
var cmsImageColumns = CmsImageColumns{
	Id:          "id",
	Title:       "title",
	ChannelId:   "channel_id",
	Images:      "images",
	Description: "description",
	Flag:        "flag",
	Status:      "status",
	ClickNum:    "click_num",
	Sort:        "sort",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewCmsImageDao creates and returns a new DAO object for table data access.
func NewCmsImageDao() *CmsImageDao {
	return &CmsImageDao{
		group:   "default",
		table:   "cms_image",
		columns: cmsImageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsImageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsImageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsImageDao) Columns() CmsImageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsImageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsImageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsImageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
