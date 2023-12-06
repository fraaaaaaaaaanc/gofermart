package storagegofermart

import (
	"context"
	"database/sql"
	handlersmodels "gofermart/internal/models/handlers_models"
)

func (s *Storage) AddNewOrderAndAccrual(ctxRequest context.Context, reqOrder *handlersmodels.ReqOrder) error {
	err := s.InTransaction(ctxRequest, func(ctx context.Context, tx *sql.Tx) error {
		reqOrder, err := s.AddNewOrder(ctx, tx, reqOrder)
		if err != nil {
			return err
		}
		err = s.AddNewOrderAccrual(ctx, tx, reqOrder)
		return err
	})

	return err
}
