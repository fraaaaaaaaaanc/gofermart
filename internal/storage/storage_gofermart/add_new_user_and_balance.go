package storagegofermart

import (
	"context"
	"database/sql"
	handlersmodels "gofermart/internal/models/handlers_models"
)

func (s *Storage) AddNewUserAndBalance(ctxRequest context.Context, reqRegister handlersmodels.RequestRegister) (int, error) {
	var userID int
	err := s.InTransaction(ctxRequest, func(ctx context.Context, tx *sql.Tx) error {
		var err error
		userID, err = s.AddNewUser(ctx, tx, &reqRegister)
		if err != nil {
			return err
		}
		err = s.AddNewUserBalance(ctx, tx, userID)
		return err
	})

	return userID, err
}
