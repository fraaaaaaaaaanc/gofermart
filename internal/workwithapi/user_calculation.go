package workwithapi

import (
	"context"
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	workwithapimodels "gofermart/internal/models/work_with_api_models"
)

func (w *WorkAPI) userCalculation() error {
	err := w.InTransaction(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		usersOrdersAccrualList, err := w.GetCalculatedUsers()
		if err != nil && !errors.Is(err, workwithapimodels.ErrNoUsers) {
			logger.Error("error when receiving users and points for their calculated orders", zap.Error(err))
			return err
		}
		if errors.Is(err, workwithapimodels.ErrNoUsers) {
			logger.Error("no users", zap.Error(err))
		}

		err = w.UpdateBalance(ctx, tx, usersOrdersAccrualList)
		if err != nil {
			logger.Error("error update balance", zap.Error(err))
		}

		err = w.DeleteOrders(ctx, tx)
		if err != nil {
			logger.Error("error delete orders", zap.Error(err))
		}

		return nil
	})

	return err
}
