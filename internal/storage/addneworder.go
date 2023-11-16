package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gofermart/internal/models/handlersmodels"
	"time"
)

func (s *Storage) AddNewOrder(reqOrder *handlersmodels.ReqOrder) error {
	ctx, cansel := context.WithTimeout(reqOrder.Ctx, time.Second*1)
	defer cansel()

	row := s.DB.QueryRowContext(ctx,
		"INSERT INTO orders (user_id, order_number, order_status)"+
			"VALUES ($1, $2, $3)",
		reqOrder.UserID, reqOrder.OrderNumber, reqOrder.OrderStatus)

	err := row.Err()
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.UniqueViolation == pgErr.Code {
			row = s.DB.QueryRowContext(ctx,
				"SELECT user_id FROM orders WHERE order_number = $1",
				reqOrder.OrderNumber)
		} else {
			return err
		}
		var userID int
		if err = row.Scan(&userID); err != nil {
			return err
		}
		if userID != reqOrder.UserID {
			return handlersmodels.ErrConflictOrderNumberAnotherUser
		} else {
			return handlersmodels.ErrConflictOrderNumberSameUser
		}
	}
	return nil
}
