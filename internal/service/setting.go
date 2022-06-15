package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/yaml.v3"
	"log"
	"net/url"
	"os"
)

var (
	insSetting = sSetting{}
)

//设置
type sSetting struct{}

func Setting() *sSetting {
	return &insSetting
}

type SettingGroups struct {
	Title    string            `yaml:"title"`
	Children []SettingChildren `yaml:"children"`
}
type SettingChildren struct {
	Title       string           `yaml:"title"`
	Type        string           `yaml:"type"`
	Name        string           `yaml:"name"`
	Default     string           `yaml:"default"`
	Tip         string           `yaml:"tip"`
	Description string           `yaml:"description"`
	Options     []SettingOptions `yaml:"options"`
}
type SettingConfig struct {
	Backend Settings `yaml:"backend"`
	Web     Settings `yaml:"web"`
}
type Settings struct {
	Title  string          `yaml:"title"`
	Groups []SettingGroups `yaml:"groups"`
}

type SettingOptions struct {
	Title string `yaml:"title"`
	Value string `yaml:"value"`
}

type SettingNames struct {
}

func (*sSetting) readYamlConfig(path string) (*SettingConfig, error) {
	conf := &SettingConfig{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	//fmt.Println("SettingConfig: ", conf)
	return conf, nil
}

func (*sSetting) readYaml() *SettingConfig {
	conf, err := Setting().readYamlConfig(Util().SystemRoot() + "/manifest/config/setting.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

// BackendAll 获取所有后台菜单
func (*sSetting) BackendAll() []SettingGroups {
	cacheKey := PublicCachePreFix + ":settings:backend_all"
	result, err := g.Redis().Do(Ctx, "GET", cacheKey)
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		var settingGroups []SettingGroups
		if err = result.Structs(&settingGroups); err != nil {
			panic(err)
		}
		return settingGroups
	}
	//g.Log().Debug(Ctx, "Setting().readYaml()", Setting().readYaml())
	backendAll := Setting().readYaml().Backend.Groups
	_, err = g.Redis().Do(Ctx, "SET", cacheKey, backendAll)
	if err != nil {
		panic(err)
	}
	return backendAll
}

// Save 保存设置
func (*sSetting) Save(forms url.Values) bool {
	group := "backend"
	model := "system_setting"
	names := g.Slice{}
	for name, value := range forms {
		names = append(names, name)
		one, err := g.Model(model).Where("group", group).Where("name", name).One()
		if err != nil {
			panic(err)
		}
		if one.IsEmpty() {
			go g.Model(model).Data(g.Map{
				"group": group,
				"name":  name,
				"value": value[0],
			}).Insert()
		} else {
			go g.Model(model).Data(g.Map{
				"name":  name,
				"value": value[0],
			}).Where("group", group).Where("name", name).Update()
		}
	}
	go g.Model(model).Where("group", group).WhereNotIn("name", names).Delete()
	go Util().ClearSystemSettingCache()
	return true
}
