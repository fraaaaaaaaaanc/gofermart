package storagegofermart

import (
	"context"
	"database/sql"
	handlersmodels "gofermart/internal/models/handlers_models"
)

func (s *Storage) CheckOrderNumber(ctx context.Context, tx *sql.Tx, orderNumber string) error {
	row := tx.QueryRowContext(ctx,
		"SELECT EXISTS (SELECT 1 FROM orders WHERE order_number = $1);",
		orderNumber)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return err
	}
	if exists {
		return handlersmodels.ErrDuplicateOrderNumber
	}
	return nil
}
