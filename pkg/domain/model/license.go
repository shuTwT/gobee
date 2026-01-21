package model

import "time"

type LicenseCreateReq struct {
	MachineCode   string    `json:"machine_code" validate:"required"`
	LicenseKey    string    `json:"license_key" validate:"required"`
	CustomerName  string    `json:"customer_name"`
	ExpireDate    time.Time `json:"expire_date" validate:"required"`
}

type LicenseUpdateReq struct {
	MachineCode  string    `json:"machine_code,omitempty"`
	LicenseKey   string    `json:"license_key,omitempty"`
	CustomerName string    `json:"customer_name,omitempty"`
	ExpireDate   time.Time `json:"expire_date,omitempty"`
	Status       int       `json:"status,omitempty"`
}

type LicenseResp struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	MachineCode  string    `json:"machine_code"`
	LicenseKey   string    `json:"license_key"`
	CustomerName string    `json:"customer_name"`
	ExpireDate   time.Time `json:"expire_date"`
	Status       int       `json:"status"`
}

type LicenseVerifyReq struct {
	MachineCode string `json:"machine_code" validate:"required"`
}

type LicenseVerifyResp struct {
	Valid        bool   `json:"valid"`
	CustomerName string `json:"customer_name"`
	ExpireDate   string `json:"expire_date"`
	Message      string `json:"message"`
}
