// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsAdminDao is the data access object for table cms_admin.
type CmsAdminDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CmsAdminColumns // columns contains all the column names of Table for convenient usage.
}

// CmsAdminColumns defines and stores column names for table cms_admin.
type CmsAdminColumns struct {
	Id        string //
	Username  string // 用户名
	Password  string // 密码
	Name      string // 姓名
	Tel       string // 手机号
	Email     string // 邮箱
	Status    string // 状态
	IsSystem  string // 是否系统用户
	CreatedAt string //
	UpdatedAt string //
}

// cmsAdminColumns holds the columns for table cms_admin.
var cmsAdminColumns = CmsAdminColumns{
	Id:        "id",
	Username:  "username",
	Password:  "password",
	Name:      "name",
	Tel:       "tel",
	Email:     "email",
	Status:    "status",
	IsSystem:  "is_system",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCmsAdminDao creates and returns a new DAO object for table data access.
func NewCmsAdminDao() *CmsAdminDao {
	return &CmsAdminDao{
		group:   "default",
		table:   "cms_admin",
		columns: cmsAdminColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsAdminDao) Columns() CmsAdminColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsAdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
