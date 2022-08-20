package role

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sRole struct{}
)

var (
	insRole = sRole{}
)

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func Role() *sRole {
	return &insRole
}

func (s *sRole) BackendRoleGetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	var (
		m    = dao.CmsRole.Ctx(ctx).OrderAsc(dao.CmsRole.Columns().Id)
		list []*entity.CmsRole
	)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	listModel := m.Page(in.Page, in.Size)
	err = listModel.Scan(&list)
	if err != nil {
		return nil, err
	}
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

// BackendApiRoleStatus 修改角色状态
func (s *sRole) BackendApiRoleStatus(ctx context.Context, in *backendApi.RoleStatusReq) (out interface{}, err error) {
	var role *entity.CmsRole
	err = dao.CmsRole.Ctx(ctx).Where(dao.CmsRole.Columns().Id, in.Id).Scan(&role)
	if err != nil {
		return nil, err
	}
	if role.IsSystem == 1 && role.IsEnable == 1 {
		return nil, gerror.New("系统角色无法被停用")
	}
	isEnable := 0
	if role.IsEnable == 0 {
		isEnable = 1
	}
	_, err = dao.CmsRole.Ctx(ctx).Where(dao.CmsRole.Columns().Id, in.Id).Data(g.Map{
		dao.CmsRole.Columns().IsEnable: isEnable,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

// BackendApiRoleDelete 角色删除
func (s *sRole) BackendApiRoleDelete(ctx context.Context, in *backendApi.RoleDeleteReq) (out interface{}, err error) {
	var role *entity.CmsRole
	err = dao.CmsRole.Ctx(ctx).Where(dao.CmsRole.Columns().Id, in.Id).Scan(&role)
	if err != nil {
		return nil, err
	}
	if role.IsSystem == 1 {
		return nil, gerror.New("系统角色无法被删除")
	}
	_, err = dao.CmsRole.Ctx(ctx).Where(dao.CmsRole.Columns().Id, in.Id).Delete()
	if err != nil {
		return nil, err
	}
	//删除拥有该角色的用户映射
	_, err = dao.CmsRoleAccount.Ctx(ctx).Where(dao.CmsRoleAccount.Columns().RoleId, in.Id).Delete()
	if err != nil {
		return nil, err
	}
	//删除该角色的权限
	_, err = dao.CmsRulePermissions.Ctx(ctx).Where(dao.CmsRulePermissions.Columns().V0, in.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

// BackendApiRoleDeleteBatch 角色批量删除
func (s *sRole) BackendApiRoleDeleteBatch(ctx context.Context, in *backendApi.RoleDeleteBatchReq) (out interface{}, err error) {
	err = dao.CmsRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var err error
		for _, id := range in.Ids {
			var role *entity.CmsRole
			err := tx.Ctx(ctx).Model(entity.CmsRole{}).Where(dao.CmsRole.Columns().Id, id).Scan(&role)
			if err != nil {
				return err
			}
			if role.IsSystem == 1 {
				return gerror.New("删除失败，存在无法被删除的系统管理员：" + gvar.New(role.Id).String())
			}
			_, err = tx.Ctx(ctx).Model(entity.CmsRole{}).Where(dao.CmsRole.Columns().Id, id).Delete()
			if err != nil {
				return err
			}
			//删除拥有该角色的用户映射
			_, err = tx.Ctx(ctx).Model(entity.CmsRoleAccount{}).Where(dao.CmsRoleAccount.Columns().RoleId, id).Delete()
			if err != nil {
				return err
			}
			//删除该角色的权限
			_, err = tx.Ctx(ctx).Model(entity.CmsRulePermissions{}).Where(dao.CmsRulePermissions.Columns().V0, id).Delete()
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

// BackendApiRoleAdd 添加角色
func (s *sRole) BackendApiRoleAdd(ctx context.Context, in *backendApi.RoleAddReq) (out interface{}, err error) {
	var role *entity.CmsRole
	err = dao.CmsRole.Ctx(ctx).Where(dao.CmsRole.Columns().Title, in.Title).Scan(&role)
	if err != nil {
		return nil, err
	}
	if role != nil {
		return nil, gerror.New("角色名已存在")
	}
	//写入角色
	roleId, err := dao.CmsRole.Ctx(ctx).InsertAndGetId(g.Map{
		dao.CmsRole.Columns().Title:       in.Title,
		dao.CmsRole.Columns().Description: in.Description,
		dao.CmsRole.Columns().IsEnable:    in.Status,
		dao.CmsRole.Columns().Type:        "backend",
	})
	if err != nil {
		return nil, err
	}
	var rulePermissions []interface{}
	for _, rule := range in.Rules {
		var rulePermission = g.Map{}
		rulePermission["p_type"] = "p"
		rulePermission["v0"] = roleId
		rulePermission["v1"] = "backend"
		rulePermission["v2"] = rule
		rulePermissions = append(rulePermissions, rulePermission)
	}
	_, err = dao.CmsRulePermissions.Ctx(ctx).Insert(rulePermissions)
	if err != nil {
		return nil, err
	}
	return
}
