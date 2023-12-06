package storagegofermart

import (
	"context"
	"database/sql"
	handlersmodels "gofermart/internal/models/handlers_models"
)

func (s *Storage) ProcessingDebitingFunds(ctxRequest context.Context, reqWithdraw handlersmodels.ReqWithdraw) error {
	err := s.InTransaction(ctxRequest, func(ctx context.Context, tx *sql.Tx) error {
		if err := s.CheckOrderNumber(ctx, tx, reqWithdraw.OrderNumber); err != nil {
			return err
		}
		if err := s.WithdrawBalance(ctx, tx, reqWithdraw); err != nil {
			return err
		}

		return s.AddHistoryBalance(ctx, tx, reqWithdraw)
	})

	return err
}
