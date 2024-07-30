package types

import "time"

type OrderStore interface {
	CreateOrder(Order) error
	CreateOrderItem(OrderItem) error
	DeleteOrder(int) error
	GetOrderById(int) (*Order, error)
	ListOrders() ([]Order, error)
	UpdateOrder(int, Order) error
	GetOrdersByStatus(string) ([]Order, error)
	GetOrdersByUserId(int) ([]Order, error)
}

type OrderStatus string

const (
	New        OrderStatus = "new"
	In_Process OrderStatus = "in_process"
	Done       OrderStatus = "done"
)

type Order struct {
	ID        int         `json:"id"`
	UserID    int         `json:"user_id"`
	Total     float64     `json:"total"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderI_D"`
	ProductID int       `json:"productID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateOrderPayload struct {
	UserID int         `json:"user_id" validate:"required"`
	Total  float64     `json:"total" validate:"required"`
	Status OrderStatus `json:"status" validate:"required"`
}

type UpdateOrderPayload struct {
	UserID int         `json:"user_id" validate:"omitempty"`
	Total  float64     `json:"total" validate:"omitempty"`
	Status OrderStatus `json:"status" validate:"omitempty"`
}

type CreateOrderItemPayload struct {
	ProductID int       `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required"`
}
