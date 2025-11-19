package model

import "time"

// PayOrder represents the payment order domain model.
type PayOrder struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ChannelID  string    `json:"channel_id"`
	OrderID    string    `json:"order_id"`
	OutTradeNo string    `json:"out_trade_no"`
	TotalFee   string    `json:"total_fee"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	NotifyURL  string    `json:"notify_url"`
	ReturnURL  string    `json:"return_url"`
	Extra      string    `json:"extra"`
	PayURL     string    `json:"pay_url,omitempty"`
	State      string    `json:"state"`
	ErrorMsg   string    `json:"error_msg,omitempty"`
	Raw        string    `json:"raw,omitempty"`
}

// PayOrderCreateReq represents the request body for creating a payment order.
type PayOrderCreateReq struct {
	ChannelID  string `json:"channel_id" validate:"required"`
	OrderID    string `json:"order_id"`
	OutTradeNo string `json:"out_trade_no" validate:"required"`
	TotalFee   string `json:"total_fee" validate:"required"`
	Subject    string `json:"subject" validate:"required"`
	Body       string `json:"body" validate:"required"`
	NotifyURL  string `json:"notify_url" validate:"required,url"`
	ReturnURL  string `json:"return_url" validate:"required,url"`
	Extra      string `json:"extra"`
	PayURL     string `json:"pay_url,omitempty"`
	State      string `json:"state"`
	ErrorMsg   string `json:"error_msg,omitempty"`
	Raw        string `json:"raw,omitempty"`
}

// PayOrderUpdateReq represents the request body for updating a payment order.
type PayOrderUpdateReq struct {
	ChannelID  string `json:"channel_id"`
	OrderID    string `json:"order_id"`
	OutTradeNo string `json:"out_trade_no"`
	TotalFee   string `json:"total_fee"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	NotifyURL  string `json:"notify_url"`
	ReturnURL  string `json:"return_url"`
	Extra      string `json:"extra"`
	PayURL     string `json:"pay_url,omitempty"`
	State      string `json:"state"`
	ErrorMsg   string `json:"error_msg,omitempty"`
	Raw        string `json:"raw,omitempty"`
}

// PayOrderResp represents the response body for a payment order.
type PayOrderResp struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ChannelID  string    `json:"channel_id"`
	OrderID    string    `json:"order_id"`
	OutTradeNo string    `json:"out_trade_no"`
	TotalFee   string    `json:"total_fee"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	NotifyURL  string    `json:"notify_url"`
	ReturnURL  string    `json:"return_url"`
	Extra      string    `json:"extra"`
	PayURL     *string   `json:"pay_url,omitempty"`
	State      string    `json:"state"`
	ErrorMsg   *string   `json:"error_msg,omitempty"`
}
