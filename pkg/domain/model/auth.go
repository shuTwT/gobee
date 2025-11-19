package model

// LoginRequest 登录请求结构
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResp struct {
	AccessToken string   `json:"accessToken"`
	Expires     int64    `json:"expires"`
	Username    string   `json:"username"`
	Roles       []string `json:"roles"`
}
