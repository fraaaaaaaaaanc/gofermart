package storage

import (
	"context"
	"gofermart/internal/models/handlersmodels"
)

func (s *Storage) CheckOrderNumber(ctx context.Context, orderNumber string) error {
	row := s.DB.QueryRowContext(ctx,
		"SELECT EXISTS (SELECT 1 FROM orders WHERE order_number = $1);",
		orderNumber)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return err
	}
	if exists {
		return handlersmodels.ErrDuplicateOrderNumberOrders
	}
	return nil
}
