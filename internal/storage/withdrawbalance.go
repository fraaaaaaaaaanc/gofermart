package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	cookiemodels "gofermart/internal/models/cookie"
	"gofermart/internal/models/handlersmodels"
	"time"
)

func (s *Storage) WithdrawBalance(reqWithdraw handlersmodels.ReqWithdraw) error {
	ctx, cansel := context.WithTimeout(reqWithdraw.Ctx, time.Second*1)
	defer cansel()

	tx, err := s.DB.Begin()
	if err != nil {
		return err

	}
	userID := reqWithdraw.Ctx.Value(cookiemodels.UserID).(int)
	_, err = tx.ExecContext(ctx,
		"UPDATE balance "+
			"Set user_balance = user_balance - $1, withdrawn_balance = withdrawn_balance + $1 "+
			"WHERE user_id = $2",
		reqWithdraw.SumWithdraw, userID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) && pgErr.Code == "CHECK_VIOLATION" {
			err = handlersmodels.ErrNegativeBalanceValue
		}
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx,
		"INSERT INTO history_balance (order_number, user_id, withdrawn_sum) "+
			"VALUES ($1, $2, $3)",
		reqWithdraw.OrderNumber, userID, reqWithdraw.SumWithdraw)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			err = handlersmodels.ErrDuplicateOrderNumber
		}
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
