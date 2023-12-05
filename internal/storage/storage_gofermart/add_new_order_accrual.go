package storagegofermart

import (
	"context"
	"database/sql"
	handlersmodels "gofermart/internal/models/handlers_models"
)

func (s *Storage) AddNewOrderAccrual(ctx context.Context, tx *sql.Tx, reqOrder *handlersmodels.ReqOrder) error {
	_, err := tx.ExecContext(ctx,
		"INSERT INTO order_accrual (order_id, user_id, order_number, order_status_accrual) "+
			"VALUES ($1, $2, $3, $4)",
		reqOrder.OrderID, reqOrder.UserID, reqOrder.OrderNumber, orderStatusNew)
	if err != nil {
		return err
	}

	return nil
}
