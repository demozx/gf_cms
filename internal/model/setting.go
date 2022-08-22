package model

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
	BackendView Settings `yaml:"backend_view"`
	Web         Settings `yaml:"web"`
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
