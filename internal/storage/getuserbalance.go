package storage

import (
	"context"
	cookiemodels "gofermart/internal/models/cookie"
	"gofermart/internal/models/handlersmodels"
	"time"
)

func (s *Storage) GetUserBalance(ctx context.Context) (*handlersmodels.RespUserBalance, error) {
	userID := ctx.Value(cookiemodels.UserID).(int)

	newCtx, cansel := context.WithTimeout(ctx, time.Second*1)
	defer cansel()

	row := s.DB.QueryRowContext(newCtx,
		"SELECT user_balance, withdrawn_balance FROM balance "+
			"WHERE user_id = $1",
		userID)

	respUserBalance := &handlersmodels.RespUserBalance{}
	if err := row.Scan(&respUserBalance.UserBalance, &respUserBalance.WithdrawnBalance); err != nil {
		return nil, err
	}
	return respUserBalance, nil
}