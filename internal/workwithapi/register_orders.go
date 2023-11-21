package workwithapi

//
//import (
//	"errors"
//	"go.uber.org/zap"
//	"gofermart/internal/models/work_with_api_models"
//	"time"
//)
//
//func (w *WorkAPI) registerOrders() {
//	ticker := time.NewTicker(time.Second *
//		5)
//
//	for range ticker.C {
//		unRegisterOrdersList, err := w.strg.GetAllUnRegisterOrders()
//		if err != nil && !errors.Is(err, work_with_api_models.ErrNoOrdersForRegistration) {
//			w.log.Error("error when getting a list of unregistered order numbers", zap.Error(err))
//			break
//		}
//		if errors.Is(err, work_with_api_models.ErrNoOrdersForRegistration) {
//			w.log.Error("no orders", zap.Error(err))
//			break
//		}
//
//		w.RegisterOrderNumber(unRegisterOrdersList)
//	}
//
//	//var orderInfoList []*handlers_models.OrderInfo
//	//for {
//	//	select {
//	//	case orderInfo := <-w.OrderCh:
//	//		orderInfoList = append(orderInfoList, orderInfo)
//	//	case <-ticker.C:
//	//		if len(orderInfoList) == 0 {
//	//			continue
//	//		}
//	//		for _, orderInfo := range orderInfoList {
//	//			w.RegisterOrderNumber(orderInfo)
//	//		}
//	//		orderInfoList = nil
//	//	}
//	//}
//}
