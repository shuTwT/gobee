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

type PersonalAccessTokenCreateReq struct {
	// 令牌名称
	Name string `json:"name"`
	// 过期时间
	Expires LocalTime `json:"expires"`
	// 描述
	Description string `json:"description"`
}

type PersonalAccessTokenListResp struct {
	ID int `json:"id"`
	// 令牌名称
	Name string `json:"name"`
	// 过期时间
	Expires *LocalTime `json:"expires"`
	// 描述
	Description string `json:"description"`
}

type PersonalAccessTokenResp struct {
	ID int `json:"id"`
	// 令牌名称
	Name string `json:"name"`
	// 过期时间
	Expires *LocalTime `json:"expires"`
	// 描述
	Description string `json:"description"`
	// 令牌
	Token string `json:"token"`
}
