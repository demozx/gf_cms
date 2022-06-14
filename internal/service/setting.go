package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/yaml.v3"
	"log"
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
	Title       string `yaml:"title"`
	Type        string `yaml:"type"`
	Name        string `yaml:"name"`
	Default     string `yaml:"default"`
	Tip         string `yaml:"tip"`
	Description string `yaml:"description"`
}
type SettingConfig struct {
	Backend Settings `yaml:"backend"`
	Web     Settings `yaml:"web"`
}
type Settings struct {
	Title  string          `yaml:"title"`
	Groups []SettingGroups `yaml:"groups"`
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
