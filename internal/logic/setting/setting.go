package setting

import (
	"context"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	insSetting = sSetting{}
)

// 设置
type sSetting struct{}

func init() {
	service.RegisterSetting(New())
}

func New() *sSetting {
	return &sSetting{}
}

func Setting() *sSetting {
	return &insSetting
}

func (*sSetting) readYaml(ctx context.Context) (conf *model.SettingConfig, err error) {
	data, err := g.Cfg("setting").Data(ctx)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(data, &conf)
	if err != nil {
		return nil, err
	}
	return
}

// BackendViewAll 获取所有后台菜单
func (*sSetting) BackendViewAll() (backendAll []model.SettingGroups, err error) {
	cacheKey := util.PublicCachePreFix + ":settings:backend_all"
	result, err := g.Redis().Do(util.Ctx, "GET", cacheKey)
	if err != nil {
		return nil, err
	}
	if !result.IsEmpty() {
		var settingGroups []model.SettingGroups
		if err = result.Structs(&settingGroups); err != nil {
			return nil, err
		}
		return settingGroups, nil
	}
	conf, _ := Setting().readYaml(util.Ctx)
	backendAll = conf.Backend.Groups
	_, err = g.Redis().Do(util.Ctx, "SET", cacheKey, backendAll)
	if err != nil {
		return nil, err
	}
	return backendAll, nil
}

// Save 保存设置
func (*sSetting) Save(forms map[string]interface{}) (res bool, err error) {
	group := "backend"
	settingModel := "system_setting"
	names := g.Slice{}
	for name, value := range forms {
		names = append(names, name)
		one, err := g.Model(settingModel).Where("group", group).Where("name", name).One()
		if err != nil {
			return false, err
		}
		if one.IsEmpty() {
			_, err := g.Model(settingModel).Data(g.Map{
				"group": group,
				"name":  name,
				"value": value,
			}).Insert()
			if err != nil {
				return false, err
			}
		} else {
			_, err = g.Model(settingModel).Data(g.Map{
				"name":  name,
				"value": value,
			}).Where("group", group).Where("name", name).Update()
			if err != nil {
				return false, err
			}
		}
	}
	_, err = g.Model(settingModel).Where("group", group).WhereNotIn("name", names).Delete()
	if err != nil {
		return false, err
	}
	_, err = util.Util().ClearSystemSettingCache()
	if err != nil {
		return false, err
	}
	return true, nil
}
