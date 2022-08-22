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
	BackendView Permissions `yaml:"backend_view"`
	BackendApi  Permissions `yaml:"backend_api"`
	WebView     Permissions `yaml:"web_view"`
	WebApi      Permissions `yaml:"web_api"`
}

type PermissionAllItem struct {
	Slug                   string                  `yaml:"slug"`
	Title                  string                  `yaml:"title"`
	BackendViewPermissions []PermissionPermissions `yaml:"permissions"`
	BackendApiPermissions  []PermissionPermissions `yaml:"permissions"`
}
