package storagegofermart

import (
	"context"
	"gofermart/internal/models/handlers_models"
	"time"
)

func (s *Storage) GetAllHistoryBalance(ctxRequest context.Context, userID int) ([]handlersmodels.RespWithdrawalsHistory, error) {
	ctx, cansel := context.WithTimeout(ctxRequest, durationWorkCtx)
	defer cansel()

	rows, err := s.db.QueryContext(ctx, "SELECT order_number_unregister, withdrawn_sum, withdrawn_datetime FROM history_balance "+
		"WHERE user_id = $1 ORDER BY withdrawn_datetime ASC",
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var respWithdrawalsHistory []handlersmodels.RespWithdrawalsHistory
	for rows.Next() {
		var withdrawal handlersmodels.RespWithdrawalsHistory
		var processedAt time.Time
		if err = rows.Scan(&withdrawal.OrderNumber,
			&withdrawal.SumWithdraw,
			&processedAt); err != nil {
			return nil, err
		}
		withdrawal.ProcessedAt = processedAt.Format(time.RFC3339)

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
