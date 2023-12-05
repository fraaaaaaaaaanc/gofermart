package storagegofermart

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	handlersmodels "gofermart/internal/models/handlers_models"
)

func (s *Storage) AddNewUser(ctx context.Context, tx *sql.Tx, reqRegister *handlersmodels.RequestRegister) (int, error) {
	row := tx.QueryRowContext(ctx,
		"INSERT INTO users (user_name, password)"+
			"VALUES ($1, $2) RETURNING id",
		reqRegister.Login, reqRegister.Password)

	err := row.Err()
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.UniqueViolation == pgErr.Code {
			return 0, handlersmodels.ErrConflictLoginRegister
		} else {
			return 0, err
		}
	}

	var userID int
	if err = row.Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}
