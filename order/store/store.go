package store

import (
	"database/sql"
	"fmt"

	"github.com/4lerman/e_com/order/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db,
	}
}

func (s *Store) ListOrders() ([]types.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := []types.Order{}
	for rows.Next() {
		order, err := scanRowIntoOrder(rows)
		if err != nil {
			return nil, err
		}

		orders = append(orders, *order)
	}

	return orders, nil
}

func (s *Store) CreateOrder(order types.Order) error {
	_, err := s.db.Exec("INSERT INTO orders (userId, total, status) VALUES ($1, $2, $3)",
		order.UserID, order.Total, order.Status)

	if err != nil {
		return err
	}

	return nil
}


func (s *Store) GetOrderById(orderId int) (*types.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders WHERE id = $1", orderId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	order := new(types.Order)
	for rows.Next() {
		order, err = scanRowIntoOrder(rows)
		if err != nil {
			return order, nil
		}
	}

	if order.ID == 0 {
		return nil, fmt.Errorf("order not found")
	}

	return order, nil
}

func (s *Store) GetOrdersByUserId(userId int) ([]types.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders WHERE userId = $1", userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := []types.Order{}
	for rows.Next() {
		order, err := scanRowIntoOrder(rows)
		if err != nil {
			return nil, err
		}

		orders = append(orders, *order)
	}


	return orders, nil
}

func (s *Store) GetOrdersByStatus(status string) ([]types.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders WHERE status = $1", status)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := []types.Order{}
	for rows.Next() {
		order, err := scanRowIntoOrder(rows)
		if err != nil {
			return nil, err
		}

		orders = append(orders, *order)
	}


	return orders, nil
}

func (s *Store) UpdateOrder(orderId int, order types.Order) error {
	_, err := s.db.Exec("UPDATE orders SET "+
		"userId = $1, total = $2, status = $3 WHERE id = $4", order.UserID, order.Total, order.Status, orderId)

	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}

	return nil
}

func (s *Store) DeleteOrder(orderId int) error {
	_, err := s.db.Exec("DELETE FROM orders WHERE id = $1", orderId)

	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}

	return nil
}


func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (orderid, productid, quantity, price) VALUES ($1, $2, $3, $4)",
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)

	if err != nil {
		return err
	}
	return nil
}

func scanRowIntoOrder(rows *sql.Rows) (*types.Order, error) {
	order := new(types.Order)

	err := rows.Scan(
		&order.ID,
		&order.UserID,
		&order.Total,
		&order.Status,
		&order.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}
