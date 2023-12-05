package workwithapi

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	workwithapimodels "gofermart/internal/models/work_with_api_models"
)

func (w *WorkAPI) updateOrdersStatusAndAccrual(resGetOrderAccrual *workwithapimodels.ResGetOrderAccrual) error {
	err := w.InTransaction(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		for _, respGetRequest := range resGetOrderAccrual.RespGetRequestList {
			err := w.UpdateOrdersStatus(ctx, tx, respGetRequest)
			if err != nil {
				logger.Error("error when changing data in the data base table orders", zap.Error(err))
				return err
			}

			err = w.UpdateOrdersAccrual(ctx, tx, respGetRequest)
			if err != nil {
				logger.Error("error when changing data in the data base table order_accrual", zap.Error(err))
				return err
			}
		}
		return nil
	})

	return err
}
