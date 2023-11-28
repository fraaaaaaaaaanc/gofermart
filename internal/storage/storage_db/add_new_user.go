package storagedb

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	handlersmodels "gofermart/internal/models/handlers_models"
	"time"
)

func (s *Storage) AddNewUser(reqChanelRegister *handlersmodels.RequestRegister) (int, error) {
	ctx, cansel := context.WithTimeout(reqChanelRegister.Ctx, time.Second*1)
	defer cansel()

	var userID int
	err := s.inTransaction(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx,
			"INSERT INTO users (user_name, password)"+
				"VALUES ($1, $2) RETURNING id",
			reqChanelRegister.Login, reqChanelRegister.Password)

		err := row.Err()
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgerrcode.UniqueViolation == pgErr.Code {
				return handlersmodels.ErrConflictLoginRegister
			} else {
				return err
			}
		}

		if err = row.Scan(&userID); err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx,
			"INSERT INTO balance (user_id)"+
				"VALUES ($1)",
			userID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return userID, nil
}
