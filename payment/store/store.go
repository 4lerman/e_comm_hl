package store

import (
	"database/sql"
	"fmt"

	"github.com/4lerman/e_com/payment/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) ListPayments() ([]types.Payment, error) {
	rows, err := s.db.Query("SELECT * FROM payments")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payments := []types.Payment{}
	for rows.Next() {
		payment, err := ScanRowIntoPayment(rows)
		if err != nil {
			return nil, err
		}

		payments = append(payments, *payment)
	}

	return payments, nil
}

func (s *Store) CreatePayment(payment types.Payment) error {
	_, err := s.db.Exec("INSERT INTO payments (userId, orderId, amount, status)"+
		"VALUES ($1, $2, $3, $4)", payment.UserID, payment.OrderID, payment.Amount, payment.Status)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetPaymentById(paymentId int) (*types.Payment, error) {
	rows, err := s.db.Query("SELECT * FROM payments WHERE id = $1", paymentId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payment := new(types.Payment)
	for rows.Next() {
		payment, err = ScanRowIntoPayment(rows)
		if err != nil {
			return payment, nil
		}
	}

	if payment.ID == 0 {
		return nil, fmt.Errorf("payment not found")
	}

	return payment, nil
}

func (s *Store) GetPaymentsByStatus(status string) ([]types.Payment, error) {
	rows, err := s.db.Query("SELECT * FROM payments WHERE status = $1", status)

	if err != nil {
		return nil, err
	}

	payments := []types.Payment{}
	for rows.Next() {
		payment, err := ScanRowIntoPayment(rows)
		if err != nil {
			return nil, err
		}

		payments = append(payments, *payment)
	}

	return payments, nil
}

func (s *Store) GetPaymentsByUserId(userId int) ([]types.Payment, error) {
	rows, err := s.db.Query("SELECT * FROM payments WHERE userId = $1", userId)

	if err != nil {
		return nil, err
	}

	payments := []types.Payment{}
	for rows.Next() {
		payment, err := ScanRowIntoPayment(rows)
		if err != nil {
			return nil, err
		}

		payments = append(payments, *payment)
	}

	return payments, nil
}

func (s *Store) GetPaymentsByOrderId(orderId int) ([]types.Payment, error) {
	rows, err := s.db.Query("SELECT * FROM payments WHERE orderId = $1", orderId)

	if err != nil {
		return nil, err
	}

	payments := []types.Payment{}
	for rows.Next() {
		payment, err := ScanRowIntoPayment(rows)
		if err != nil {
			return nil, err
		}

		payments = append(payments, *payment)
	}

	return payments, nil
}


func (s *Store) UpdatePayment(paymentId int, payment types.Payment) error {
	_, err := s.db.Exec("UPDATE payments SET "+
		"userId = $1, orderId = $2, amount = $3 WHERE id = $4", payment.UserID, payment.OrderID, payment.Amount, paymentId)

	if err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}

	return nil
}

func (s *Store) DeletePayment(paymentId int) error {
	_, err := s.db.Exec("DELETE FROM payments WHERE id = $1", paymentId)

	if err != nil {
		return fmt.Errorf("failed to delete payment: %w", err)
	}

	return nil
}

func ScanRowIntoPayment(rows *sql.Rows) (*types.Payment, error) {
	payment := new(types.Payment)

	err := rows.Scan(
		&payment.ID,
		&payment.UserID,
		&payment.OrderID,
		&payment.Amount,
		&payment.PaymentDate,
		&payment.Status,
	)

	if err != nil {
		return nil, err
	}

	return payment, nil
}
