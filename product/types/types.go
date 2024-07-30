package types

import "time"

type ProductStore interface {
	GetProducts() ([]Product, error)
	CreateProduct(Product) error
	GetProductByID(int) (*Product, error)
	UpdateProduct(int, Product) error
	DeleteProduct(int) error
	GetProductsByName(string) ([]Product, error)
	GetProductsByCategory(string) ([]Product, error)
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"omitempty"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	Category    string  `json:"category" validate:"required"`
}

type UpdateProductPayload struct {
	Name        string  `json:"name" validate:"omitempty"`
	Description string  `json:"description" validate:"omitempty"`
	Price       float64 `json:"price" validate:"omitempty"`
	Quantity    int     `json:"quantity" validate:"omitempty"`
	Category    string  `json:"category" validate:"omitempty"`
}
