// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsShortcutDao is the data access object for table cms_shortcut.
type CmsShortcutDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CmsShortcutColumns // columns contains all the column names of Table for convenient usage.
}

// CmsShortcutColumns defines and stores column names for table cms_shortcut.
type CmsShortcutColumns struct {
	Id        string //
	AccountId string // 用户id
	Name      string // 快捷方式名称
	Route     string // 路由
	Sort      string // 排序
	CreatedAt string //
	UpdatedAt string //
}

// cmsShortcutColumns holds the columns for table cms_shortcut.
var cmsShortcutColumns = CmsShortcutColumns{
	Id:        "id",
	AccountId: "account_id",
	Name:      "name",
	Route:     "route",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCmsShortcutDao creates and returns a new DAO object for table data access.
func NewCmsShortcutDao() *CmsShortcutDao {
	return &CmsShortcutDao{
		group:   "default",
		table:   "cms_shortcut",
		columns: cmsShortcutColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsShortcutDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsShortcutDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsShortcutDao) Columns() CmsShortcutColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsShortcutDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsShortcutDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsShortcutDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
