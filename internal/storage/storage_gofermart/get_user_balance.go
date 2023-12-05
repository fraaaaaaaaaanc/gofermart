package storagegofermart

import (
	"context"
	cookiemodels "gofermart/internal/models/cookie_models"
	handlersmodels "gofermart/internal/models/handlers_models"
)

func (s *Storage) GetUserBalance(ctx context.Context) (*handlersmodels.RespUserBalance, error) {
	userID := ctx.Value(cookiemodels.UserID).(int)

	newCtx, cansel := context.WithTimeout(ctx, durationWorkCtx)
	defer cansel()

	row := s.db.QueryRowContext(newCtx,
		"SELECT user_balance, withdrawn_balance FROM balance "+
			"WHERE user_id = $1",
		userID)

	respUserBalance := &handlersmodels.RespUserBalance{}
	if err := row.Scan(&respUserBalance.UserBalance, &respUserBalance.WithdrawnBalance); err != nil {
		return nil, err
	}
	return respUserBalance, nil
}
