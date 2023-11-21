package storage_db

import (
	"gofermart/internal/models/handlers_models"
	"time"
)

func (s *Storage) GetAllHistoryBalance(userID int) ([]handlers_models.RespWithdrawalsHistory, error) {
	rows, err := s.db.Query("SELECT order_number_unregister, withdrawn_sum, withdrawn_datetime FROM history_balance "+
		"WHERE user_id = $1 ORDER BY withdrawn_datetime ASC",
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var respWithdrawalsHistory []handlers_models.RespWithdrawalsHistory
	for rows.Next() {
		var withdrawal handlers_models.RespWithdrawalsHistory
		var processedAt time.Time
		if err = rows.Scan(&withdrawal.OrderNumber,
			&withdrawal.SumWithdraw,
			&processedAt); err != nil {
			return nil, err
		}
		withdrawal.ProcessedAt = processedAt.Format(time.RFC3339)
		if err != nil {
			return nil, err
		}

		respWithdrawalsHistory = append(respWithdrawalsHistory, withdrawal)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if respWithdrawalsHistory == nil {
		return nil, handlers_models.ErrTheAreNoWithdraw
	}
	return respWithdrawalsHistory, nil
}
