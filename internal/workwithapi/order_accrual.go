package workwithapi

import (
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	"gofermart/internal/models/work_with_api_models"
	"time"
)

func (w *WorkAPI) getOrdersAccrual() {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for range ticker.C {
		ticker.Reset(time.Second * 5)
		unAccrualOrdersList, err := w.strg.GetAllUnAccrualOrders()
		if err != nil && !errors.Is(err, work_with_api_models.ErrNoOrdersForAcrrual) {
			logger.Error("error when working with the database at the time of receiving the list of "+
				"outstanding orders", zap.Error(err))
			continue
		}
		if errors.Is(err, work_with_api_models.ErrNoOrdersForAcrrual) {
			logger.Info("no orders", zap.Error(err))
			continue
		}
		resGetOrdersAccrual, err := w.getOrderAccrual(unAccrualOrdersList)
		if err != nil &&
			!errors.Is(err, work_with_api_models.ErrRequestCount) &&
			!errors.Is(err, work_with_api_models.ErrNoRespAPI) {
			logger.Error("error when receiving order accrual statuses", zap.Error(err))
			continue
		}
		if errors.Is(err, work_with_api_models.ErrNoRespAPI) {
			logger.Error("no responses", zap.Error(err))
			continue
		}
		if errors.Is(err, work_with_api_models.ErrRequestCount) {
			ticker.Reset(time.Second * time.Duration(resGetOrdersAccrual.TimeRetryAfter))
		}
		if err = w.strg.UpdateOrdersStatusAndAccrual(resGetOrdersAccrual); err != nil {
			logger.Error("error when changing data in the data base table orders/order_accrual", zap.Error(err))
			continue
		}
		usersOrdersAccrualList, err := w.strg.GetCalculatedUsers()
		if err != nil && !errors.Is(err, work_with_api_models.ErrNoUsers) {
			logger.Error("error when receiving users and points for their calculated orders", zap.Error(err))
			continue
		}
		if errors.Is(err, work_with_api_models.ErrNoUsers) {
			logger.Error("no users", zap.Error(err))
			continue
		}
		err = w.strg.UpdateBalance(usersOrdersAccrualList)
		if err != nil {
			logger.Error("error update balance or delete orders", zap.Error(err))
		}
	}
}
