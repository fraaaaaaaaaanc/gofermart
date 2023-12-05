package storagegofermart

import (
	"context"
	"database/sql"
	"gofermart/internal/models/work_with_api_models"
)

func (s *Storage) UpdateBalance(ctx context.Context, tx *sql.Tx, usersOrdersAccrualList []workwithapimodels.UsersOrdersAccrual) error {
	for _, usersOrdersAccrual := range usersOrdersAccrualList {
		_, err := tx.ExecContext(ctx,
			"UPDATE balance "+
				"SET user_balance = user_balance + $1 "+
				"WHERE user_id = $2",
			usersOrdersAccrual.OrderAccrual, usersOrdersAccrual.UserID)

		if err != nil {
			return err
		}
	}

	return nil
}
