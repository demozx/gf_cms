// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsSystemSettingDao is the data access object for table cms_system_setting.
type CmsSystemSettingDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns CmsSystemSettingColumns // columns contains all the column names of Table for convenient usage.
}

// CmsSystemSettingColumns defines and stores column names for table cms_system_setting.
type CmsSystemSettingColumns struct {
	Id        string //
	Group     string //
	Name      string // 名称
	Value     string // 值
	CreatedAt string //
	UpdatedAt string //
}

// cmsSystemSettingColumns holds the columns for table cms_system_setting.
var cmsSystemSettingColumns = CmsSystemSettingColumns{
	Id:        "id",
	Group:     "group",
	Name:      "name",
	Value:     "value",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCmsSystemSettingDao creates and returns a new DAO object for table data access.
func NewCmsSystemSettingDao() *CmsSystemSettingDao {
	return &CmsSystemSettingDao{
		group:   "default",
		table:   "cms_system_setting",
		columns: cmsSystemSettingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsSystemSettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsSystemSettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsSystemSettingDao) Columns() CmsSystemSettingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsSystemSettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsSystemSettingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsSystemSettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
