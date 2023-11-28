package storagedb

import (
	"context"
	"database/sql"
	"gofermart/internal/models/orderstatuses"
	"time"
)

func (s *Storage) UpdateOrderStatus(orderNumber string) error {
	newCtx, cansel := context.WithTimeout(context.Background(), time.Second*1)
	defer cansel()

	err := s.inTransaction(newCtx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(newCtx,
			"UPDATE orders "+
				"Set order_status = $1 "+
				"WHERE order_number = $2",
			orderstatuses.PROCESSING, orderNumber)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(newCtx,
			"UPDATE order_accrual "+
				"Set order_status_accrual = $1 "+
				"WHERE order_number = $2",
			orderstatuses.REGISTERED, orderNumber)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
