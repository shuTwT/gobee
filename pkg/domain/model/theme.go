package model

import (
	"mime/multipart"
	"time"
)

type ThemeConfig struct {
	Type            string   `yaml:"type" validate:"required"`
	Name            string   `yaml:"name" validate:"required"`
	DisplayName      string   `yaml:"display-name" validate:"required"`
	Description      string   `yaml:"description,omitempty"`
	Author          *ThemeAuthor `yaml:"author,omitempty"`
	Logo            string   `yaml:"logo,omitempty"`
	Homepage         string   `yaml:"homepage,omitempty"`
	Repo            string   `yaml:"repo,omitempty"`
	Issue            string   `yaml:"issue,omitempty"`
	SettingName      string   `yaml:"setting-name,omitempty"`
	ConfigMapName    string   `yaml:"config-map-name,omitempty"`
	Version          string   `yaml:"version" validate:"required"`
	Require          string   `yaml:"require,omitempty"`
	License          string   `yaml:"license,omitempty"`
}

type ThemeAuthor struct {
	Name  string `yaml:"name,omitempty"`
	Email string `yaml:"email,omitempty"`
}

type CreateThemeReq struct {
	File *multipart.FileHeader `form:"file" validate:"required"`
}

type ThemeResp struct {
	ID              int        `json:"id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Name            string     `json:"name"`
	DisplayName      string     `json:"display_name"`
	Description      string     `json:"description"`
	AuthorName      string     `json:"author_name"`
	AuthorEmail     string     `json:"author_email"`
	Logo            string     `json:"logo"`
	Homepage         string     `json:"homepage"`
	Repo            string     `json:"repo"`
	Issue            string     `json:"issue"`
	SettingName      string     `json:"setting_name"`
	ConfigMapName    string     `json:"config_map_name"`
	Version          string     `json:"version"`
	Require          string     `json:"require"`
	License          string     `json:"license"`
	Path            string     `json:"path"`
	Enabled          bool       `json:"enabled"`
}
