package model

import (
	"mime/multipart"
	"time"
)

type PluginConfig struct {
	Name             string   `yaml:"name" validate:"required"`
	Key              string   `yaml:"key" validate:"required"`
	Version          string   `yaml:"version" validate:"required"`
	Description      string   `yaml:"description,omitempty"`
	ProtocolVersion  string   `yaml:"protocol_version,omitempty"`
	MagicCookieKey   string   `yaml:"magic_cookie_key,omitempty"`
	MagicCookieValue string   `yaml:"magic_cookie_value" validate:"required"`
	Config           string   `yaml:"config,omitempty"`
	Dependencies     []string `yaml:"dependencies,omitempty"`
	AutoStart        bool     `yaml:"auto_start,omitempty"`
}

type CreatePluginReq struct {
	File *multipart.FileHeader `form:"file" validate:"required"`
}

type PluginResp struct {
	ID               int        `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	Key              string     `json:"key"`
	Name             string     `json:"name"`
	Version          string     `json:"version"`
	Description      string     `json:"description"`
	BinPath          string     `json:"bin_path"`
	ProtocolVersion  string     `json:"protocol_version"`
	MagicCookieKey   string     `json:"magic_cookie_key"`
	MagicCookieValue string     `json:"magic_cookie_value"`
	Dependencies     []string   `json:"dependencies"`
	Config           string     `json:"config"`
	Enabled          bool       `json:"enabled"`
	AutoStart        bool       `json:"auto_start"`
	Status           string     `json:"status"`
	LastError        string     `json:"last_error"`
	LastStartedAt    *time.Time `json:"last_started_at"`
	LastStoppedAt    *time.Time `json:"last_stopped_at"`
}
