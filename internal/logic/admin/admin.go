package admin

import (
	"context"
	"crypto/md5"
	"fmt"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/captcha"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/model/do"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
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

// BackendApiAdminAdd 添加管理员
func (s *sAdmin) BackendApiAdminAdd(ctx context.Context, in *backendApi.AdminAddReq) (out interface{}, err error) {
	one, err := dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Username, in.Username).One()
	if err != nil {
		return nil, err
	}
	if !one.IsEmpty() {
		return nil, gerror.New("用户名已存在")
	}
	one, err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Tel, in.Tel).One()
	if err != nil {
		return nil, err
	}
	if !one.IsEmpty() {
		return nil, gerror.New("手机号已存在")
	}
	one, err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Email, in.Email).One()
	if err != nil {
		return nil, err
	}
	if !one.IsEmpty() {
		return nil, gerror.New("邮箱号已存在")
	}

	err = dao.CmsAdmin.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var err error
		//写入用户
		in.Password = Admin().passMd5(in.Password)
		id, err := tx.Ctx(ctx).Model(entity.CmsAdmin{}).Data(in).InsertAndGetId()
		if err != nil {
			return err
		}
		for _, roleId := range in.Role {
			//写入用户角色
			_, err = tx.Model(entity.CmsRoleAccount{}).Data(g.Map{
				"account_id": id,
				"role_id":    roleId,
			}).Insert()
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return
}

// BackendApiAdminEdit 编辑
func (s *sAdmin) BackendApiAdminEdit(ctx context.Context, in *backendApi.AdminEditReq) (out interface{}, err error) {
	var admin *entity.CmsAdmin
	err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Id, in.Id).Scan(&admin)
	if err != nil {
		return nil, err
	}
	g.Dump(admin)
	return
}

// BackendApiAdminStatus 修改自动状态
func (s *sAdmin) BackendApiAdminStatus(ctx context.Context, in *backendApi.AdminStatusReq) (out interface{}, err error) {
	var admin *entity.CmsAdmin
	err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Id, in.Id).Scan(&admin)
	if err != nil {
		return nil, err
	}
	if admin.IsSystem == 1 && admin.Status == 1 {
		return nil, gerror.New("系统用户无法被停用")
	}
	status := 0
	if admin.Status == 0 {
		status = 1
	}
	_, err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Id, in.Id).Update(g.Map{
		dao.CmsAdmin.Columns().Status: status,
	})
	if err != nil {
		return nil, err
	}
	return
}

// BackendApiAdminDelete 删除
func (s *sAdmin) BackendApiAdminDelete(ctx context.Context, in *backendApi.AdminDeleteReq) (out interface{}, err error) {
	var admin *entity.CmsAdmin
	err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Id, in.Id).Scan(&admin)
	if err != nil {
		return nil, err
	}
	if admin.IsSystem == 1 {
		return nil, gerror.New("系统用户无法被删除")
	}
	_, err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Id, in.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sAdmin) BackendApiAdminDeleteBatch(ctx context.Context, in *backendApi.AdminDeleteBatchReq) (out interface{}, err error) {
	err = dao.CmsAdmin.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var err error
		for _, id := range in.Ids {
			var admin *entity.CmsAdmin
			err = tx.Ctx(ctx).Model(entity.CmsAdmin{}).Where(dao.CmsAdmin.Columns().Id, id).Scan(&admin)
			if err != nil {
				return err
			}
			if admin.IsSystem == 1 {
				return gerror.New("删除失败，存在无法被删除的系统管理员：" + gvar.New(admin.Id).String())
			}
			_, err = tx.Ctx(ctx).Model(entity.CmsAdmin{}).Where(dao.CmsAdmin.Columns().Id, id).Delete()
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return
}
