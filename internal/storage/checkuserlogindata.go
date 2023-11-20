package storage

import (
	"context"
	"database/sql"
	"errors"
	"gofermart/internal/models/handlersmodels"
	"time"
)

func (s *Storage) CheckUserLoginData(reqLogin *handlersmodels.RequestLogin) (*handlersmodels.ResultLogin, error) {
	ctx, cansel := context.WithTimeout(reqLogin.Ctx, time.Second*1)
	defer cansel()

	row := s.DB.QueryRowContext(ctx,
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
