// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsFriendlyLinkDao is the data access object for table cms_friendly_link.
type CmsFriendlyLinkDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns CmsFriendlyLinkColumns // columns contains all the column names of Table for convenient usage.
}

// CmsFriendlyLinkColumns defines and stores column names for table cms_friendly_link.
type CmsFriendlyLinkColumns struct {
	Id        string //
	Name      string // 链接名称
	Url       string // 链接地址
	Status    string // 状态
	Sort      string // 排序
	CreatedAt string //
	UpdatedAt string //
}

// cmsFriendlyLinkColumns holds the columns for table cms_friendly_link.
var cmsFriendlyLinkColumns = CmsFriendlyLinkColumns{
	Id:        "id",
	Name:      "name",
	Url:       "url",
	Status:    "status",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCmsFriendlyLinkDao creates and returns a new DAO object for table data access.
func NewCmsFriendlyLinkDao() *CmsFriendlyLinkDao {
	return &CmsFriendlyLinkDao{
		group:   "default",
		table:   "cms_friendly_link",
		columns: cmsFriendlyLinkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsFriendlyLinkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsFriendlyLinkDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsFriendlyLinkDao) Columns() CmsFriendlyLinkColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsFriendlyLinkDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsFriendlyLinkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsFriendlyLinkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
