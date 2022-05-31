package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service/internal/dao"
	"gf_cms/internal/service/internal/do"
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

func Admin() *sAdmin {
	return &insAdmin
}

// LoginVerify 登录验证
func (s *sAdmin) LoginVerify(ctx context.Context, in model.AdminLoginInput) (admin *entity.CmsAdmin, err error) {
	// 验证验证码
	if !Captcha().Verify(in.CaptchaId, in.CaptchaStr) {
		return admin, gerror.New(`验证码错误`)
	}
	md5Password := passMd5(in.Password)
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

	return admin, nil
}

func (s *sAdmin) GetUserByUserNamePassword(ctx context.Context, in model.AdminLoginInput) g.Map {
	var admin *entity.CmsAdmin
	md5Password := passMd5(in.Password)
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
func passMd5(password string) string {
	bytePassword := []byte(password)
	md5Password := fmt.Sprintf("%x", md5.Sum(bytePassword))
	return md5Password
}

// GetRoleIdsByAccountId 获取用户的所有角色id
func GetRoleIdsByAccountId(accountId string) []gdb.Value {
	roleIds, err := dao.CmsRoleAccount.Ctx(Ctx).Where("account_id", accountId).Fields("role_id").Array()
	if err != nil {
		panic(err)
	}
	return roleIds
}
