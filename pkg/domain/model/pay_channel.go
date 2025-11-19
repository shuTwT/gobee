package model

import "time"

// PayChannel represents the payment channel domain model.
type PayChannel struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Type      string    `json:"type"`
	Config    string    `json:"config"`
}

// PayChannelCreateReq represents the request body for creating a payment channel.

type PayChannelCreateReq struct {
	Name   string `json:"name" validate:"required"`
	Code   string `json:"code" validate:"required"`
	Type   string `json:"type" validate:"required"`
	Config string `json:"config" validate:"required"`
}

// PayChannelUpdateReq represents the request body for updating a payment channel.

type PayChannelUpdateReq struct {
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Type   string `json:"type,omitempty"`
	Config string `json:"config,omitempty"`
}

// PayChannelResp represents the response body for a payment channel.

type PayChannelResp struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Type      string    `json:"type"`
	Config    string    `json:"config"`
}
