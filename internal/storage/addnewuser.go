package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gofermart/internal/models/handlersmodels"
	"time"
)

func (s *Storage) AddNewUser(reqChanelRegister *handlersmodels.RequestRegister) (int, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return 0, err
	}

	ctx, cansel := context.WithTimeout(reqChanelRegister.Ctx, time.Second*1)
	defer cansel()

	row := tx.QueryRowContext(ctx,
		"INSERT INTO users (user_name, password)"+
			"VALUES ($1, $2) RETURNING id",
		reqChanelRegister.Login, reqChanelRegister.Password)

	err = row.Err()
	if err != nil {
		tx.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.UniqueViolation == pgErr.Code {
			return 0, handlersmodels.ErrConflictLoginRegister
		} else {
			return 0, err
		}
	}

	var userID int
	if err = row.Scan(&userID); err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.ExecContext(ctx,
		"INSERT INTO balance (user_id)"+
			"VALUES ($1)",
		userID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return userID, nil
}
