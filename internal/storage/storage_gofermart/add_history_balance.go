package storagegofermart

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	handlersmodels "gofermart/internal/models/handlers_models"
)

func (s *Storage) AddHistoryBalance(ctx context.Context, tx *sql.Tx, reqWithdraw handlersmodels.ReqWithdraw) error {
	_, err := tx.ExecContext(ctx,
		"INSERT INTO history_balance (order_number_unregister, user_id, withdrawn_sum) "+
			"VALUES ($1, $2, $3)",
		reqWithdraw.OrderNumber, reqWithdraw.UserID, reqWithdraw.SumWithdraw)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			err = handlersmodels.ErrDuplicateOrderNumberHistoryBalance
		}
		return err
	}
	return nil
}
