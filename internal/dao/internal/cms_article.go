// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsArticleDao is the data access object for table cms_article.
type CmsArticleDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CmsArticleColumns // columns contains all the column names of Table for convenient usage.
}

// CmsArticleColumns defines and stores column names for table cms_article.
type CmsArticleColumns struct {
	Id          string // 文章id
	Title       string // 标题
	ChannelId   string // 所属栏目id
	Keyword     string // 关键词
	Description string // 文章摘要
	Flag        string // 属性(p:带图,r:推荐,t:置顶)
	Status      string // 审核状态(1:已审核,0:未审核)
	Thumb       string // 缩略图
	CopyFrom    string // 文章来源
	ClickNum    string // 点击数
	Sort        string // 排序
	CreatedAt   string // 发布时间
	UpdatedAt   string // 编辑时间
	DeletedAt   string // 删除时间
}

// cmsArticleColumns holds the columns for table cms_article.
var cmsArticleColumns = CmsArticleColumns{
	Id:          "id",
	Title:       "title",
	ChannelId:   "channel_id",
	Keyword:     "keyword",
	Description: "description",
	Flag:        "flag",
	Status:      "status",
	Thumb:       "thumb",
	CopyFrom:    "copy_from",
	ClickNum:    "click_num",
	Sort:        "sort",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewCmsArticleDao creates and returns a new DAO object for table data access.
func NewCmsArticleDao() *CmsArticleDao {
	return &CmsArticleDao{
		group:   "default",
		table:   "cms_article",
		columns: cmsArticleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsArticleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsArticleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsArticleDao) Columns() CmsArticleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsArticleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsArticleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsArticleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
