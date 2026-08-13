package model

// LoginRequest 登录请求结构
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResp struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	Expires      int64    `json:"expires"`
	Username     string   `json:"username"`
	Roles        []string `json:"roles"`
}

// RefreshTokenRequest 刷新令牌请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

// RefreshTokenResp 刷新令牌响应，字段与 LoginResp 对齐
type RefreshTokenResp struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	Expires      int64    `json:"expires"`
	Username     string   `json:"username"`
	Roles        []string `json:"roles"`
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

type LoginUser struct {
	ID       int      `json:"id"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}
