package storagedb

import (
	"context"
	"database/sql"
	"gofermart/internal/models/work_with_api_models"
	"time"
)

func (s *Storage) UpdateBalance(usersOrdersAccrualList []workwithapimodels.UsersOrdersAccrual) error {
	ctx, cansel := context.WithTimeout(context.Background(), time.Second*1)
	defer cansel()

	err := s.inTransaction(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
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
		err := s.DeleteOrders(tx)
		if err != nil {
			return err
		}
		return err
	})
	return err
}
