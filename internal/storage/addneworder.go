package storage

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gofermart/internal/models/handlersmodels"
	"time"
)

func (s *Storage) AddNewOrder(reqOrder *handlersmodels.ReqOrder) error {
	ctx, cansel := context.WithTimeout(reqOrder.Ctx, time.Second*1)
	defer cansel()

	err := s.inTransaction(ctx, s.DB, func(ctx context.Context, tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx,
			"INSERT INTO orders (user_id, order_number, order_status)"+
				"VALUES ($1, $2, $3) RETURNING id",
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

		var orderID int
		if err = row.Scan(&orderID); err != nil {
			return err
		}

		//goods := reqOrder.Goods
		//for _, product := range goods {
		//	_, err = tx.ExecContext(ctx,
		//		"INSERT INTO orders_info (order_id, description, price) "+
		//			"VALUES ($1, $2, $3)",
		//		orderID, product.Description, product.Price)
		//	if err != nil {
		//		return nil, err
		//	}
		//}

		_, err = tx.ExecContext(ctx,
			"INSERT INTO order_accrual (order_id, user_id, order_number, order_status_accrual) "+
				"VALUES ($1, $2, $3, $4)",
			orderID, reqOrder.UserID, reqOrder.OrderNumber, reqOrder.OrderStatus)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
