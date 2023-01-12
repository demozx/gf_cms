package admin

import (
	"context"
	"crypto/md5"
	"fmt"
	"gf_cms/api/backendApi"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/captcha"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/model/do"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/util/gconv"

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

	//角色id们
	roleIds, err := dao.CmsRoleAccount.Ctx(ctx).Where(do.CmsRoleAccount{AccountId: admin.Id}).Array(dao.CmsRoleAccount.Columns().RoleId)
	if err != nil {
		return nil, err
	}
	if len(roleIds) == 0 {
		return nil, gerror.New("该用户没有任何角色，无法登录")
	}
	//用户拥有的角色启用的数量
	count, err := dao.CmsRole.Ctx(ctx).WhereIn(dao.CmsRole.Columns().Id, roleIds).Where(dao.CmsRole.Columns().IsEnable, 1).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("该用户的角色均未启用，无法登录")
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

// md5加密
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

	err = dao.CmsAdmin.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
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
	if admin == nil {
		return nil, gerror.New("管理员不存在")
	}
	if admin.IsSystem == 1 && in.Status == 0 {
		return nil, gerror.New("系统管理员无法被禁用")
	}
	count, err := dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Username, in.Username).WhereNot(dao.CmsAdmin.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("用户名已存在")
	}
	count, err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Tel, in.Tel).WhereNot(dao.CmsAdmin.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("手机号已存在")
	}
	count, err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Email, in.Email).WhereNot(dao.CmsAdmin.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("邮箱号已存在")
	}
	adminData := g.Map{
		"username": in.Username,
		"name":     in.Name,
		"tel":      in.Tel,
		"email":    in.Email,
		"status":   in.Status,
	}
	if len(in.Password) > 0 {
		adminData["password"] = Admin().passMd5(in.Password)
	}
	_, err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Id, in.Id).Data(adminData).Update()
	if err != nil {
		return nil, err
	}
	_, err = dao.CmsRoleAccount.Ctx(ctx).Where(dao.CmsRoleAccount.Columns().AccountId, in.Id).Delete()
	if err != nil {
		return nil, err
	}
	for _, roleId := range in.Role {
		roleData := g.Map{
			"account_id": in.Id,
			"role_id":    roleId,
		}
		_, err = dao.CmsRoleAccount.Ctx(ctx).Where(dao.CmsRoleAccount.Columns().AccountId, in.Id).Data(roleData).Insert()
	}
	// 如果修改了密码，立刻退出当前用户登录
	if len(in.Password) > 0 {
		get, err := g.RequestFromCtx(ctx).Session.Get(consts.AdminSessionKeyPrefix)
		if err != nil {
			return nil, err
		}
		var cmsAdmin *entity.CmsAdmin
		err = get.Scan(&cmsAdmin)
		if err != nil {
			return nil, err
		}
		if cmsAdmin != nil && in.Id == int(cmsAdmin.Id) {
			err = g.RequestFromCtx(ctx).Session.Remove(consts.AdminSessionKeyPrefix)
			if err != nil {
				return nil, err
			}
		}
	}
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
	//删除当前管理员的角色id
	_, err = dao.CmsRoleAccount.Ctx(ctx).Where(dao.CmsRoleAccount.Columns().AccountId, in.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

// BackendApiAdminDeleteBatch 批量删除
func (s *sAdmin) BackendApiAdminDeleteBatch(ctx context.Context, in *backendApi.AdminDeleteBatchReq) (out interface{}, err error) {
	err = dao.CmsAdmin.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
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
			//删除当前管理员的角色id
			_, err = tx.Ctx(ctx).Model(entity.CmsRoleAccount{}).Where(dao.CmsRoleAccount.Columns().AccountId, id).Delete()
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

// InitAdminUser 初始化系统管理员
func (s *sAdmin) InitAdminUser(ctx context.Context) {
	//系统角色
	systemRoleOne, err := dao.CmsRole.Ctx(ctx).Where(do.CmsRole{IsSystem: 1}).One()
	if err != nil {
		return
	}
	//系统用户
	systemAdminOne, err := dao.CmsAdmin.Ctx(ctx).Where(do.CmsAdmin{IsSystem: 1}).One()
	if err != nil {
		return
	}

	if !systemRoleOne.IsEmpty() && !systemAdminOne.IsEmpty() {
		g.Log("InitAdminUser").Notice(ctx, "┌────────────────────────────┐")
		g.Log("InitAdminUser").Notice(ctx, "│无需初始化系统管理员，已跳过│")
		g.Log("InitAdminUser").Notice(ctx, "└────────────────────────────┘")
	}

	//没有系统角色
	var roleId = systemRoleOne.GMap().Get("id")
	if systemRoleOne.IsEmpty() {
		//创建系统角色
		roleId, err = dao.CmsRole.Ctx(ctx).Data(g.Map{
			dao.CmsRole.Columns().Type:        "backend",
			dao.CmsRole.Columns().IsEnable:    1,
			dao.CmsRole.Columns().IsSystem:    1,
			dao.CmsRole.Columns().Title:       "超级管理员",
			dao.CmsRole.Columns().Description: "超级管理员",
		}).InsertAndGetId()
		if err != nil {
			return
		}
		g.Log("InitAdminUser").Warning(ctx, "┌────────────────────────────")
		g.Log("InitAdminUser").Warning(ctx, "│自动初始化系统角色ID："+gconv.String(roleId)+"")
		g.Log("InitAdminUser").Warning(ctx, "└────────────────────────────")
	}
	// 没有系统管理员
	if systemAdminOne.IsEmpty() {
		//创建系统管理员
		var name = grand.Str("abcdefghijklmnopqrstuvwxyz0123456789", 6)
		var username = name
		var password = grand.Str("abcdefghijklmnopqrstuvwxyz0123456789", 10)
		adminId, err := dao.CmsAdmin.Ctx(ctx).Data(g.Map{
			dao.CmsAdmin.Columns().IsSystem: 1,
			dao.CmsAdmin.Columns().Status:   1,
			dao.CmsAdmin.Columns().Name:     name,
			dao.CmsAdmin.Columns().Username: username,
			dao.CmsAdmin.Columns().Password: Admin().passMd5(password),
		}).InsertAndGetId()
		if err != nil {
			return
		}

		//绑定用户和角色
		_, err = dao.CmsRoleAccount.Ctx(ctx).Data(g.Map{
			dao.CmsRoleAccount.Columns().RoleId:    roleId,
			dao.CmsRoleAccount.Columns().AccountId: adminId,
		}).Insert()
		if err != nil {
			return
		}

		g.Log("InitAdminUser").Warning(ctx, "┌────────────────────────────────────────────────────")
		g.Log("InitAdminUser").Warning(ctx, "│自动初始化系统管理员ID："+gconv.String(adminId)+"")
		g.Log("InitAdminUser").Warning(ctx, "│请使用如下信息登录后台，并修改密码")
		g.Log("InitAdminUser").Warning(ctx, "│用户名："+gconv.String(username)+"")
		g.Log("InitAdminUser").Warning(ctx, "│密码："+gconv.String(password)+"")
		g.Log("InitAdminUser").Warning(ctx, "└────────────────────────────────────────────────────")
	}
}
