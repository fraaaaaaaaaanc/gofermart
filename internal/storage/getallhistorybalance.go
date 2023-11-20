package storage

import (
	"context"
	cookiemodels "gofermart/internal/models/cookie"
	"gofermart/internal/models/handlersmodels"
	"time"
)

func (s *Storage) GetAllHistoryBalance(ctx context.Context) ([]handlersmodels.RespWithdrawalsHistory, error) {
	userID := ctx.Value(cookiemodels.UserID).(int)
	rows, err := s.DB.Query("SELECT order_number_unregister, withdrawn_sum, withdrawn_datetime FROM history_balance "+
		"WHERE user_id = $1 ORDER BY withdrawn_datetime ASC",
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var respWithdrawalsHistory []handlersmodels.RespWithdrawalsHistory
	for rows.Next() {
		var withdrawal handlersmodels.RespWithdrawalsHistory
		var processedAt string
		if err = rows.Scan(&withdrawal.OrderNumber,
			&withdrawal.SumWithdraw,
			&processedAt); err != nil {
			return nil, err
		}
		withdrawal.ProcessedAt, err = time.Parse(time.RFC3339, processedAt)
		if err != nil {
			return nil, err
		}

		respWithdrawalsHistory = append(respWithdrawalsHistory, withdrawal)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if respWithdrawalsHistory == nil {
		return nil, handlersmodels.ErrTheAreNoWithdraw
	}
	return respWithdrawalsHistory, nil
}
