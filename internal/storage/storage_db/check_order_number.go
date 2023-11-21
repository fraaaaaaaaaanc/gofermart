package storage_db

import (
	"context"
	"gofermart/internal/models/handlers_models"
)

func (s *Storage) CheckOrderNumber(ctx context.Context, orderNumber string) error {
	row := s.db.QueryRowContext(ctx,
		"SELECT EXISTS (SELECT 1 FROM orders WHERE order_number = $1);",
		orderNumber)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return err
	}
	if exists {
		return handlers_models.ErrDuplicateOrderNumber
	}
	return nil
}
