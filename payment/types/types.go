package types

import (
	"time"
)

type PaymentStore interface {
	CreatePayment(Payment) error
	DeletePayment(int) error
	GetPaymentById(int) (*Payment, error)
	ListPayments() ([]Payment, error)
	UpdatePayment(int, Payment) error
	GetPaymentsByStatus(string) ([]Payment, error)
	GetPaymentsByUserId(int) ([]Payment, error)
	GetPaymentsByOrderId(int) ([]Payment, error)
}

type PaymentStatus string

const (
	Success PaymentStatus = "success"
	Failed  PaymentStatus = "failed"
)

type Payment struct {
	ID          int           `json:"id"`
	UserID      int           `json:"user_id"`
	OrderID     int           `json:"order_id"`
	Amount      float64       `json:"amount"`
	PaymentDate time.Time     `json:"payment_date"`
	Status      PaymentStatus `json:"status"`
}

type CreatePaymentPayload struct {
	UserID  int           `json:"user_id" validate:"required"`
	OrderID int           `json:"order_id" validate:"required"`
	Amount  float64       `json:"amount" validate:"required"`
	Status  PaymentStatus `json:"status"`
}

type UpdatePaymentPayload struct {
	UserID  int     `json:"user_id" validate:"required"`
	OrderID int     `json:"order_id" validate:"required"`
	Amount  float64 `json:"amount" validate:"required"`
}

type TokenResponse struct {
	Scope       string `json:"scope"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

type CryptogramRespone struct {
	Hpan       string `json:"hpan"`
	ExpDate    string `json:"expDate"`
	Cvc        string `json:"cvc"`
	TerminalId string `json:"terminalId"`
}

type PaymentResponse struct {
	Status    string  `json:"status"`
	Message   string  `json:"message"`
	PaymentID string  `json:"payment_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	InvoiceID string  `json:"invoice_id"`
}
