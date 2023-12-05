package storagegofermart

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gofermart/internal/models/handlers_models"
)

func (s *Storage) WithdrawBalance(ctx context.Context, tx *sql.Tx, reqWithdraw handlersmodels.ReqWithdraw) error {
	_, err := tx.ExecContext(ctx,
		"UPDATE balance "+
			"Set user_balance = user_balance - $1, withdrawn_balance = withdrawn_balance + $1 "+
			"WHERE user_id = $2",
		reqWithdraw.SumWithdraw, reqWithdraw.UserID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.CheckViolation {
			err = handlersmodels.ErrNegativeBalanceValue
		}
		return err
	}

	return nil
}
