package storage_db

import (
	"context"
	"database/sql"
	"gofermart/internal/models/orderstatuses"
	"time"
)

func (s *Storage) DeleteOrders(tx *sql.Tx) error {
	ctx, cansel := context.WithTimeout(context.Background(), time.Second*1)
	defer cansel()

	_, err := tx.ExecContext(ctx,
		"DELETE FROM order_accrual "+
			"WHERE order_status_accrual IN ($1, $2)",
		orderstatuses.INVALID, orderstatuses.PROCESSED)
	if err != nil {
		return err
	}

	return err
}
