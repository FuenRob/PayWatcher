package model

import "time"

type Payment struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	CategoryID  uint      `json:"category_id"`
	UserID      uint      `json:"user_id"`
	NetAmount   float64   `json:"net_amount"`
	GrossAmount float64   `json:"gross_amount"`
	Deductible  float64   `json:"deductible"`
	ChargeDate  time.Time `json:"charge_date"`
	Recurrent   bool      `json:"recurrent"`
	PaymentType string    `json:"payment_type"`
	Paid        bool      `json:"paid"`
}

type UpdateOrCreatePayment struct {
	Name        string  `json:"name"`
	CategoryID  uint    `json:"category_id"`
	NetAmount   float64 `json:"net_amount"`
	GrossAmount float64 `json:"gross_amount"`
	Deductible  float64 `json:"deductible"`
	ChargeDate  string  `json:"charge_date"`
	Recurrent   bool    `json:"recurrent"`
	PaymentType string  `json:"payment_type"`
	Paid        bool    `json:"paid"`
}
