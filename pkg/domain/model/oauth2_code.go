package model

import "time"

// Oauth2Code represents the OAuth2 authorization code domain model.
type Oauth2Code struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
	Code        string    `json:"code"`
	ExpireAt    time.Time `json:"expire_at"`
	ClientID    string    `json:"client_id"`
	RedirectURI string    `json:"redirect_uri"`
	Scope       string    `json:"scope"`
}

// Oauth2CodeCreateReq represents the request body for creating an OAuth2 authorization code.

type Oauth2CodeCreateReq struct {
	UserID      int       `json:"user_id" validate:"required"`
	Code        string    `json:"code" validate:"required"`
	ExpireAt    time.Time `json:"expire_at" validate:"required"`
	ClientID    string    `json:"client_id" validate:"required"`
	RedirectURI string    `json:"redirect_uri" validate:"required,url"`
	Scope       string    `json:"scope" validate:"required"`
}

// Oauth2CodeUpdateReq represents the request body for updating an OAuth2 authorization code.

type Oauth2CodeUpdateReq struct {
	ExpireAt *time.Time `json:"expire_at,omitempty"`
	Scope    *string    `json:"scope,omitempty"`
}

// Oauth2CodeResp represents the response body for an OAuth2 authorization code.

type Oauth2CodeResp struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int       `json:"user_id"`
	Code        string    `json:"code"`
	ExpireAt    time.Time `json:"expire_at"`
	ClientID    string    `json:"client_id"`
	RedirectURI string    `json:"redirect_uri"`
	Scope       string    `json:"scope"`
}