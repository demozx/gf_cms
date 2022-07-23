package admin

import (
	"context"
	"crypto/md5"
	"fmt"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/captcha"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/model/do"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sAdmin struct{}
)

var (
	insAdmin = sAdmin{}
)

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

func Admin() *sAdmin {
	return &insAdmin
}

// LoginVerify 登录验证
func (s *sAdmin) LoginVerify(ctx context.Context, in model.AdminLoginInput) (admin *entity.CmsAdmin, err error) {
	// 验证验证码
	if !captcha.Captcha().Verify(in.CaptchaId, in.CaptchaStr) {
		return admin, gerror.New(`验证码错误`)
	}
	md5Password := Admin().passMd5(in.Password)
	err = dao.CmsAdmin.Ctx(ctx).Where(do.CmsAdmin{
		Username: in.Username,
		Password: md5Password,
	}).Scan(&admin)
	if err != nil {
		return admin, err
	}
	if admin == nil {
		return admin, gerror.New(`用户名或密码错误`)
	}

	if admin.Status == 0 {
		return admin, gerror.New(`用户已被封禁`)
	}

	return admin, nil
}

func (s *sAdmin) GetUserByUserNamePassword(ctx context.Context, in model.AdminLoginInput) g.Map {
	var admin *entity.CmsAdmin
	md5Password := Admin().passMd5(in.Password)
	dao.CmsAdmin.Ctx(ctx).Where(do.CmsAdmin{
		Username: in.Username,
		Password: md5Password,
	}).Scan(&admin)
	return g.Map{
		"id":       admin.Id,
		"username": admin.Username,
	}
}

//md5加密
func (s *sAdmin) passMd5(password string) string {
	bytePassword := []byte(password)
	md5Password := fmt.Sprintf("%x", md5.Sum(bytePassword))
	return md5Password
}

// GetRoleIdsByAccountId 获取用户的所有角色id
func (s *sAdmin) GetRoleIdsByAccountId(accountId string) []gdb.Value {
	roleIds, err := dao.CmsRoleAccount.Ctx(util.Ctx).Where("account_id", accountId).Fields("role_id").Array()
	if err != nil {
		panic(err)
	}
	return roleIds
}

// BackendAdminGetList 后台获取管理员列表
func (s *sAdmin) BackendAdminGetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	var (
		m = dao.CmsAdmin.Ctx(ctx).OrderDesc(dao.CmsAdmin.Columns().Id)
	)
	out = &model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	var list []*entity.CmsAdmin

	err = listModel.Scan(&list)
	if err != nil {
		return nil, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	if err := listModel.WithAll().Scan(&out.List); err != nil {
		return out, err
	}

	return
}
