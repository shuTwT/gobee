package model

type MemberCreateReq struct {
	UserID      int    `json:"user_id" validate:"required"`
	MemberLevel int    `json:"member_level" validate:"required"`
	MemberNo    string `json:"member_no" validate:"required"`
	ExpireTime  string `json:"expire_time,omitempty"`
	Remark      string `json:"remark,omitempty"`
}

type MemberUpdateReq struct {
	MemberLevel *int    `json:"member_level,omitempty"`
	ExpireTime  *string `json:"expire_time,omitempty"`
	Points      *int    `json:"points,omitempty"`
	TotalSpent  *int    `json:"total_spent,omitempty"`
	OrderCount  *int    `json:"order_count,omitempty"`
	Active      *bool   `json:"active,omitempty"`
	Remark      *string `json:"remark,omitempty"`
}

type MemberResp struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	MemberLevel int    `json:"member_level"`
	MemberNo    string `json:"member_no"`
	JoinTime    string `json:"join_time"`
	ExpireTime  string `json:"expire_time"`
	Points      int    `json:"points"`
	TotalSpent  int    `json:"total_spent"`
	OrderCount  int    `json:"order_count"`
	Active      bool   `json:"active"`
	Remark      string `json:"remark,omitempty"`
}
