package model

// InitializeRequest 定义了初始化请求的结构
type InitializeRequest struct {
	DBType          string `json:"dbType" validate:"required,oneof=sqlite mysql postgres"`
	DBHost          string `json:"dbHost"`
	DBPort          int    `json:"dbPort"`
	DBUser          string `json:"dbUser"`
	DBPassword      string `json:"dbPassword"`
	DBName          string `json:"dbName"`
	AdminUsername   string `json:"adminUsername" validate:"required"`
	AdminPassword   string `json:"adminPassword" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6"`
	AdminEmail      string `json:"adminEmail" validate:"required,email"`
}
