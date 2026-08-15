package model

type CouponUsageCreateReq struct {
	CouponCode     string    `json:"coupon_code" validate:"required"`
	UserID         int       `json:"user_id" validate:"required"`
	OrderID        *int      `json:"order_id,omitempty"`
	DiscountAmount int       `json:"discount_amount" validate:"required"`
	ExpireAt       LocalTime `json:"expire_at" validate:"required"`
	Remark         *string   `json:"remark,omitempty"`
}

type CouponUsageUpdateReq struct {
	OrderID        *int       `json:"order_id,omitempty"`
	Status         *int       `json:"status,omitempty"`
	UsedAt         *LocalTime `json:"used_at,omitempty"`
	DiscountAmount *int       `json:"discount_amount,omitempty"`
	ExpireAt       *LocalTime `json:"expire_at,omitempty"`
	Remark         *string    `json:"remark,omitempty"`
}

type CouponUsageBatchUpdateReq struct {
	IDs    []int `json:"ids" validate:"required"`
	Status *int  `json:"status,omitempty"`
}

type CouponUsageBatchDeleteReq struct {
	IDs []int `json:"ids" validate:"required"`
}

type CouponUsageSearchReq struct {
	CouponCode *string `json:"coupon_code" query:"coupon_code" form:"coupon_code"`
	UserID     *int    `json:"user_id" query:"user_id" form:"user_id"`
	Status     *int    `json:"status" query:"status" form:"status"`
	Page       int     `json:"page" query:"page" form:"page" validate:"required,min=1"`
	Size       int     `json:"page_size" query:"page_size" form:"page_size" validate:"required,min=1,max=100"`
}

type CouponUsageResp struct {
	ID             int        `json:"id"`
	CreatedAt      LocalTime  `json:"created_at"`
	UpdatedAt      LocalTime  `json:"updated_at"`
	CouponCode     string     `json:"coupon_code"`
	UserID         int        `json:"user_id"`
	OrderID        int        `json:"order_id,omitempty"`
	Status         int        `json:"status"`
	UsedAt         *LocalTime `json:"used_at,omitempty"`
	DiscountAmount int        `json:"discount_amount"`
	ExpireAt       LocalTime  `json:"expire_at"`
	Remark         string     `json:"remark,omitempty"`
}
