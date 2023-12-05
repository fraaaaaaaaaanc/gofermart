package storagegofermart

import (
	"context"
	handlersmodels "gofermart/internal/models/handlers_models"
	"time"
)

func (s *Storage) GetAllUserOrders(ctxRequest context.Context, userID int) ([]handlersmodels.RespGetOrders, error) {
	ctx, cansel := context.WithTimeout(ctxRequest, durationWorkCtx)
	defer cansel()

	rows, err := s.db.QueryContext(ctx, "SELECT order_number, order_status, accrual, order_datetime FROM orders"+
		" WHERE user_id = $1 ORDER BY order_datetime ASC",
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var respGetOrders []handlersmodels.RespGetOrders
	for rows.Next() {
		var orderInfo handlersmodels.RespGetOrders
		var uploadedAt time.Time
		if err = rows.Scan(&orderInfo.OrderNumber,
			&orderInfo.Status,
			&orderInfo.Accrual,
			&uploadedAt); err != nil {
			return nil, err
		}
		orderInfo.UploadedAt = uploadedAt.Format(time.RFC3339)
		respGetOrders = append(respGetOrders, orderInfo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if respGetOrders == nil {
		return nil, handlersmodels.ErrTheAreNoOrders
	}

	return respGetOrders, nil
}
