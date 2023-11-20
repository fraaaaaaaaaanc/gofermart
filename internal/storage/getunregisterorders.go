package storage

import (
	"gofermart/internal/models/orderstatuses"
	"gofermart/internal/models/workwithapimodels"
)

func (s *Storage) GetAllUnRegisterOrders() ([]*workwithapimodels.UnRegisterOrders, error) {

	rows, err := s.DB.Query("SELECT order_number, description, price "+
		"FROM order_accrual oa "+
		"JOIN orders_info oi ON oa.order_id = oi.order_id "+
		"WHERE oa.order_status_accrual = $1;",
		orderstatuses.NEW)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[string]*workwithapimodels.UnRegisterOrders)

	for rows.Next() {
		var orderNumber, description string
		var price float64
		if err = rows.Scan(&orderNumber, &description, &price); err != nil {
			return nil, err
		}
		order, ok := ordersMap[orderNumber]
		if !ok {
			order = &workwithapimodels.UnRegisterOrders{
				OrderNumber: orderNumber,
				OrderInfo:   make([]workwithapimodels.OrderInfo, 0),
			}
			ordersMap[orderNumber] = order
		}

		orderInfo := workwithapimodels.OrderInfo{
			Description: description,
			Price:       price,
		}
		order.OrderInfo = append(order.OrderInfo, orderInfo)
	}

	var unRegisterOrdersList []*workwithapimodels.UnRegisterOrders
	for _, order := range ordersMap {
		unRegisterOrdersList = append(unRegisterOrdersList, order)
	}

	if unRegisterOrdersList == nil {
		return nil, workwithapimodels.ErrNoOrdersForRegistration
	}

	return unRegisterOrdersList, nil
}
