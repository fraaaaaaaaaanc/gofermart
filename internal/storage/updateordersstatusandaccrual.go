package storage

import (
	"context"
	"database/sql"
	"gofermart/internal/models/workwithapimodels"
	"time"
)

func (s *Storage) UpdateOrdersStatusAndAccrual(resGetOrdersAccrual *workwithapimodels.ResGetOrderAccrual) error {
	ctx, cansel := context.WithTimeout(context.Background(), time.Second*1)
	defer cansel()

	err := s.inTransaction(ctx, s.DB, func(ctx context.Context, tx *sql.Tx) error {
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
	//_, err = tx.ExecContext(ctx,
	//	"UPDATE orders SET order_status = ANY ($1), accrual = ANY ($2) "+
	//		"WHERE order_number = ANY ($3)",
	//	resGetOrdersAccrual.OrdersStatuses, resGetOrdersAccrual.OrdersAccruals, resGetOrdersAccrual.OrdersNumbers)
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//_, err = tx.ExecContext(ctx,
	//	"UPDATE order_accrual SET order_status_accrual = ANY ($1) "+
	//		"WHERE order_number = ANY ($2)",
	//	resGetOrdersAccrual.OrdersStatuses, resGetOrdersAccrual.OrdersNumbers)
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}

	return err
}
