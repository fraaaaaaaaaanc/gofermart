package storagegofermart

import (
	"context"
	"database/sql"
	"errors"
	handlersmodels "gofermart/internal/models/handlers_models"
)

func (s *Storage) CheckUserLoginData(reqLogin *handlersmodels.RequestLogin) (*handlersmodels.ResultLogin, error) {
	ctx, cansel := context.WithTimeout(reqLogin.Ctx, durationWorkCtx)
	defer cansel()

	row := s.db.QueryRowContext(ctx,
		"SELECT id, password FROM users WHERE user_name = $1",
		reqLogin.Login)

	var userID int
	var password string
	err := row.Scan(&userID, &password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, handlersmodels.ErrMissingDataInTable
		} else {
			return nil, err
		}
	}
	return &handlersmodels.ResultLogin{
		UserID:   userID,
		Password: password,
	}, nil
}
