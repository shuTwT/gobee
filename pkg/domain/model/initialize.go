package model

type PreInitResp struct {
	DBType string `json:"dbType"`
}

// InitializeRequest 定义了初始化请求的结构
type InitializeRequest struct {
	AdminUsername   string `json:"admin_username" validate:"required"`
	AdminPassword   string `json:"admin_password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6"`
	AdminEmail      string `json:"admin_email" validate:"required,email"`
}
