package storagegofermart

import (
	"context"
	"database/sql"
)

func (s *Storage) DeleteOrders(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx,
		"DELETE FROM order_accrual "+
			"WHERE order_status_accrual IN ($1, $2)",
		orderStatusInvalid, orderStatusProcessed)

	return err
}
