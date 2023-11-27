package storagedb

import (
	"context"
	"database/sql"
	"gofermart/internal/models/work_with_api_models"
	"time"
)

func (s *Storage) UpdateOrdersStatusAndAccrual(resGetOrdersAccrual *workwithapimodels.ResGetOrderAccrual) error {
	ctx, cansel := context.WithTimeout(context.Background(), time.Second*1)
	defer cansel()

	err := s.inTransaction(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		for _, respGetRequest := range resGetOrdersAccrual.RespGetRequestList {
			_, err := tx.ExecContext(ctx,
				"UPDATE orders SET order_status = $1, accrual = $2 "+
					"WHERE order_number = $3",
				respGetRequest.OrderStatus, respGetRequest.OrderAccrual, respGetRequest.OrderNumber)
			if err != nil {
				return err
			}

			_, err = tx.ExecContext(ctx,
				"UPDATE order_accrual SET order_status_accrual = $1, accrual = $2 "+
					"WHERE order_number = $3",
				respGetRequest.OrderStatus, respGetRequest.OrderAccrual, respGetRequest.OrderNumber)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
