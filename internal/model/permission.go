package model

type PermissionGroups struct {
	Slug        string                  `yaml:"slug"`
	Title       string                  `yaml:"title"`
	Permissions []PermissionPermissions `yaml:"permissions"`
}
type PermissionPermissions struct {
	Title         string `yaml:"title"`
	Slug          string `yaml:"slug"`
	HasPermission bool   `yaml:"has_permission"`
}
type Permissions struct {
	Title  string             `yaml:"title"`
	Groups []PermissionGroups `yaml:"groups"`
}
type PermissionConfig struct {
	Backend Permissions `yaml:"backend"`
	Web     Permissions `yaml:"web"`
}
