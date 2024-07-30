package store

import (
	"database/sql"
	"fmt"

	"github.com/4lerman/e_com/product/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db,
	}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		product, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func (s *Store) GetProductByID(productId int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = $1", productId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	product := new(types.Product)
	for rows.Next() {
		product, err = scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

	}

	if product.ID == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return product, nil
}

func (s *Store) CreateProduct(product types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, price, quantity, category)"+
		"VALUES ($1, $2, $3, $4, $5)", product.Name, product.Description, product.Price, product.Quantity, product.Category)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateProduct(productId int, product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET "+
		"name = $1, description = $2, price = $3, quantity = $4, category = $5  WHERE id = $6",
		product.Name, product.Description, product.Price, product.Quantity, product.Category, productId)

	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}

func (s *Store) DeleteProduct(productId int) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id = $1", productId)

	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

func (s *Store) GetProductsByName(name string) ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE name ILIKE $1", "%"+name+"%")

	if err != nil {
		return nil, err
	}


	products := []types.Product{}
	for rows.Next() {
		product, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func (s *Store) GetProductsByCategory(category string) ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE category = $1", "%"+category+"%")

	if err != nil {
		return nil, err
	}


	products := []types.Product{}
	for rows.Next() {
		product, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Category,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
