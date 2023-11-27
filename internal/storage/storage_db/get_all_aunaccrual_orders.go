package storagedb

import (
	"gofermart/internal/models/orderstatuses"
	workwithapimodels "gofermart/internal/models/work_with_api_models"
)

func (s *Storage) GetAllUnAccrualOrders() ([]string, error) {
	rows, err := s.db.Query("SELECT order_number FROM order_accrual "+
		"WHERE order_status_accrual IN ($1, $2, $3)",
		orderstatuses.NEW, orderstatuses.REGISTERED, orderstatuses.PROCESSING)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordersList []string
	for rows.Next() {
		var order string
		if err = rows.Scan(&order); err != nil {
			return nil, err
		}
		ordersList = append(ordersList, order)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if ordersList == nil {
		return nil, workwithapimodels.ErrNoOrdersForAcrrual
	}

	return ordersList, nil
}
