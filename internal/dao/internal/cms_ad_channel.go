// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsAdChannelDao is the data access object for table cms_ad_channel.
type CmsAdChannelDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns CmsAdChannelColumns // columns contains all the column names of Table for convenient usage.
}

// CmsAdChannelColumns defines and stores column names for table cms_ad_channel.
type CmsAdChannelColumns struct {
	Id          string //
	ChannelName string //
	Remarks     string //
	Sort        string //
	CreatedAt   string //
	UpdatedAt   string //
}

// cmsAdChannelColumns holds the columns for table cms_ad_channel.
var cmsAdChannelColumns = CmsAdChannelColumns{
	Id:          "id",
	ChannelName: "channel_name",
	Remarks:     "remarks",
	Sort:        "sort",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewCmsAdChannelDao creates and returns a new DAO object for table data access.
func NewCmsAdChannelDao() *CmsAdChannelDao {
	return &CmsAdChannelDao{
		group:   "default",
		table:   "cms_ad_channel",
		columns: cmsAdChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsAdChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsAdChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsAdChannelDao) Columns() CmsAdChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsAdChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsAdChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsAdChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
