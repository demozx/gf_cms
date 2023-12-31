// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsAdDao is the data access object for table cms_ad.
type CmsAdDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns CmsAdColumns // columns contains all the column names of Table for convenient usage.
}

// CmsAdColumns defines and stores column names for table cms_ad.
type CmsAdColumns struct {
	Id        string // 广告id
	ChannelId string // 栏目id
	Name      string // 广告名称
	Link      string // 链接地址
	ImgUrl    string // 图片
	Status    string // 状态(0停用,1显示)
	StartTime string // 广告开始时间
	EndTime   string // 广告结束时间
	Sort      string // 排序
	Remarks   string // 备注
	CreatedAt string //
	UpdatedAt string //
}

// cmsAdColumns holds the columns for table cms_ad.
var cmsAdColumns = CmsAdColumns{
	Id:        "id",
	ChannelId: "channel_id",
	Name:      "name",
	Link:      "link",
	ImgUrl:    "img_url",
	Status:    "status",
	StartTime: "start_time",
	EndTime:   "end_time",
	Sort:      "sort",
	Remarks:   "remarks",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCmsAdDao creates and returns a new DAO object for table data access.
func NewCmsAdDao() *CmsAdDao {
	return &CmsAdDao{
		group:   "default",
		table:   "cms_ad",
		columns: cmsAdColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsAdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsAdDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsAdDao) Columns() CmsAdColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsAdDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsAdDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsAdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
