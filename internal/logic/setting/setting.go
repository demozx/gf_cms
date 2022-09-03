package setting

import (
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"log"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/yaml.v3"
)

var (
	insSetting = sSetting{}
)

//设置
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

func (*sSetting) readYamlConfig(path string) (*model.SettingConfig, error) {
	conf := &model.SettingConfig{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	//fmt.Println("SettingConfig: ", conf)
	return conf, nil
}

func (*sSetting) readYaml() *model.SettingConfig {
	conf, err := Setting().readYamlConfig(util.Util().SystemRoot() + "/manifest/config/setting.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

// BackendViewAll 获取所有后台菜单
func (*sSetting) BackendViewAll() []model.SettingGroups {
	cacheKey := util.PublicCachePreFix + ":settings:backend_all"
	result, err := g.Redis().Do(util.Ctx, "GET", cacheKey)
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		var settingGroups []model.SettingGroups
		if err = result.Structs(&settingGroups); err != nil {
			panic(err)
		}
		return settingGroups
	}
	//g.Log().Debug(Ctx, "Setting().readYaml()", Setting().readYaml())
	backendAll := Setting().readYaml().Backend.Groups
	_, err = g.Redis().Do(util.Ctx, "SET", cacheKey, backendAll)
	if err != nil {
		panic(err)
	}
	return backendAll
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
			panic(err)
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
