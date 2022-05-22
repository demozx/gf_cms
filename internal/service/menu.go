package service

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	insMenu = sMenu{}
)

//菜单
type sMenu struct{}

func Menu() *sMenu {
	return &insMenu
}

type MenuGroups struct {
	Title    string         `yaml:"title"`
	Children []MenuChildren `yaml:"children"`
}
type MenuChildren struct {
	Title      string `yaml:"title"`
	Route      string `yaml:"route"`
	Permission string `yaml:"permission"`
}
type Menus struct {
	Title  string       `yaml:"title"`
	Groups []MenuGroups `yaml:"groups"`
}
type MenuConfig struct {
	Backend Menus `yaml:"backend"`
	Web     Menus `yaml:"web"`
}

func (*sMenu) readYamlConfig(path string) (*MenuConfig, error) {
	conf := &MenuConfig{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	//fmt.Println("conf: ", conf)
	return conf, nil
}

func (*sMenu) readYaml() *MenuConfig {
	var SystemRoot = Util().GetConfig("server.systemRoot")
	conf, err := Menu().readYamlConfig(SystemRoot + "/manifest/config/menu.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

// BackendAll Backend 获取全部后台菜单
func (*sMenu) BackendAll() []MenuGroups {
	backendAll := Menu().readYaml().Backend.Groups
	return backendAll
}

func (*sMenu) BackendMy() {

}
