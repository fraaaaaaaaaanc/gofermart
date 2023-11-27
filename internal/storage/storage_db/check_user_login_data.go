package storage_db

import (
	"context"
	"database/sql"
	"errors"
	"gofermart/internal/models/handlers_models"
	"time"
)

func (s *Storage) CheckUserLoginData(reqLogin *handlers_models.RequestLogin) (*handlers_models.ResultLogin, error) {
	ctx, cansel := context.WithTimeout(reqLogin.Ctx, time.Second*1)
	defer cansel()

	row := s.db.QueryRowContext(ctx,
		"SELECT id, password FROM users WHERE user_name = $1",
		reqLogin.Login)

	var userID int
	var password string
	err := row.Scan(&userID, &password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, handlers_models.ErrMissingDataInTable
		} else {
			return nil, err
		}
	}
	return &handlers_models.ResultLogin{
		UserID:   userID,
		Password: password,
	}, nil
}
