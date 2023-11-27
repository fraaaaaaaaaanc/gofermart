package storage_db

import (
	"context"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"time"
)

func (s *Storage) GetUserBalance(ctx context.Context) (*handlers_models.RespUserBalance, error) {
	userID := ctx.Value(cookiemodels.UserID).(int)

	newCtx, cansel := context.WithTimeout(ctx, time.Second*1)
	defer cansel()

	row := s.db.QueryRowContext(newCtx,
		"SELECT user_balance, withdrawn_balance FROM balance "+
			"WHERE user_id = $1",
		userID)

	respUserBalance := &handlers_models.RespUserBalance{}
	if err := row.Scan(&respUserBalance.UserBalance, &respUserBalance.WithdrawnBalance); err != nil {
		return nil, err
	}
	return respUserBalance, nil
}
