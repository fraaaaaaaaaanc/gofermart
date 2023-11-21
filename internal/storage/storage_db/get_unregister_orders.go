package storage_db

//
//import (
//	"gofermart/internal/models/orderstatuses"
//	"gofermart/internal/models/work_with_api_models"
//)
//
//func (s *Storage) GetAllUnRegisterOrders() ([]*work_with_api_models.UnRegisterOrders, error) {
//
//	rows, err := s.DB.Query("SELECT order_number, description, price "+
//		"FROM order_accrual oa "+
//		"JOIN orders_info oi ON oa.order_id = oi.order_id "+
//		"WHERE oa.order_status_accrual = $1;",
//		orderstatuses.NEW)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	ordersMap := make(map[string]*work_with_api_models.UnRegisterOrders)
//
//	for rows.Next() {
//		var orderNumber, description string
//		var price float64
//		if err = rows.Scan(&orderNumber, &description, &price); err != nil {
//			return nil, err
//		}
//		order, ok := ordersMap[orderNumber]
//		if !ok {
//			order = &work_with_api_models.UnRegisterOrders{
//				OrderNumber: orderNumber,
//				OrderInfo:   make([]work_with_api_models.OrderInfo, 0),
//			}
//			ordersMap[orderNumber] = order
//		}
//
//		orderInfo := work_with_api_models.OrderInfo{
//			Description: description,
//			Price:       price,
//		}
//		order.OrderInfo = append(order.OrderInfo, orderInfo)
//	}
//
//	var unRegisterOrdersList []*work_with_api_models.UnRegisterOrders
//	for _, order := range ordersMap {
//		unRegisterOrdersList = append(unRegisterOrdersList, order)
//	}
//
//	if rows.Err() != nil {
//		return nil, err
//	}
//
//	if unRegisterOrdersList == nil {
//		return nil, work_with_api_models.ErrNoOrdersForRegistration
//	}
//
//	return unRegisterOrdersList, nil
//}
