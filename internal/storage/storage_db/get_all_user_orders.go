package storage_db

import (
	"gofermart/internal/models/handlers_models"
	"time"
)

func (s *Storage) GetAllUserOrders(userID int) ([]handlers_models.RespGetOrders, error) {

	rows, err := s.db.Query("SELECT order_number, order_status, accrual, order_datetime FROM orders"+
		" WHERE user_id = $1 ORDER BY order_datetime ASC",
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var respGetOrders []handlers_models.RespGetOrders
	for rows.Next() {
		var orderInfo handlers_models.RespGetOrders
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
		return nil, handlers_models.ErrTheAreNoOrders
	}

	return respGetOrders, nil
}
