package storagegofermart

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gofermart/internal/models/handlers_models"
)

func (s *Storage) AddNewOrder(ctx context.Context, tx *sql.Tx, reqOrder *handlersmodels.ReqOrder) (*handlersmodels.ReqOrder, error) {
	row := tx.QueryRowContext(ctx,
		"INSERT INTO orders (user_id, order_number, order_status)"+
			"VALUES ($1, $2, $3) RETURNING id",
		reqOrder.UserID, reqOrder.OrderNumber, orderStatusNew)

	err := row.Err()
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.UniqueViolation == pgErr.Code {
			row = s.db.QueryRowContext(ctx,
				"SELECT user_id FROM orders WHERE order_number = $1",
				reqOrder.OrderNumber)
		} else {
			return nil, err
		}
		var userID int
		if err = row.Scan(&userID); err != nil {
			return nil, err
		}
		if userID != reqOrder.UserID {
			return nil, handlersmodels.ErrConflictOrderNumberAnotherUser
		} else {
			return nil, handlersmodels.ErrConflictOrderNumberSameUser
		}
	}

	if err = row.Scan(&reqOrder.OrderID); err != nil {
		return nil, err
	}

	return reqOrder, nil
}
