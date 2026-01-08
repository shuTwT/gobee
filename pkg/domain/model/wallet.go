package model

type WalletCreateReq struct {
	UserID int `json:"user_id" validate:"required"`
}

type WalletUpdateReq struct {
	Password string `json:"password,omitempty" validate:"min=6"`
	Remark   string `json:"remark,omitempty"`
	Active   *bool  `json:"active,omitempty"`
}

type WalletResp struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	Balance      int    `json:"balance"`
	FrozenAmount int    `json:"frozen_amount"`
	TotalIncome  int    `json:"total_income"`
	TotalExpense int    `json:"total_expense"`
	Active       bool   `json:"active"`
	Remark       string `json:"remark,omitempty"`
}
