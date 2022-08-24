package model

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
	Backend    Menus `yaml:"backend"`
	BackendApi Menus `yaml:"backend_api"`
	Web        Menus `yaml:"web"`
}
