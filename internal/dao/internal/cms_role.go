// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsRoleDao is the data access object for table cms_role.
type CmsRoleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns CmsRoleColumns // columns contains all the column names of Table for convenient usage.
}

// CmsRoleColumns defines and stores column names for table cms_role.
type CmsRoleColumns struct {
	Id          string //
	Title       string // 中文名称
	Description string //
	Type        string // 账户类型
	IsEnable    string // 是否可用
	IsSystem    string //
}

// cmsRoleColumns holds the columns for table cms_role.
var cmsRoleColumns = CmsRoleColumns{
	Id:          "id",
	Title:       "title",
	Description: "description",
	Type:        "type",
	IsEnable:    "is_enable",
	IsSystem:    "is_system",
}

// NewCmsRoleDao creates and returns a new DAO object for table data access.
func NewCmsRoleDao() *CmsRoleDao {
	return &CmsRoleDao{
		group:   "default",
		table:   "cms_role",
		columns: cmsRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsRoleDao) Columns() CmsRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
