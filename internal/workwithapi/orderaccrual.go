package workwithapi

import (
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/models/workwithapimodels"
	"time"
)

func (w *WorkAPI) getOrdersAccrual() {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for range ticker.C {
		ticker.Reset(time.Second * 5)
		unAccrualOrdersList, err := w.strg.GetAllUnAccrualOrders()
		if err != nil && !errors.Is(err, workwithapimodels.ErrNoOrdersForAcrrual) {
			w.log.Error("error when working with the database at the time of receiving the list of "+
				"outstanding orders", zap.Error(err))
			continue
		}
		if errors.Is(err, workwithapimodels.ErrNoOrdersForAcrrual) {
			w.log.Info("no orders", zap.Error(err))
			continue
		}
		resGetOrdersAccrual, err := w.getOrderAccrual(unAccrualOrdersList)
		if err != nil &&
			!errors.Is(err, workwithapimodels.ErrRequestCount) &&
			!errors.Is(err, workwithapimodels.ErrNoRespAPI) {
			w.log.Error("error when receiving order accrual statuses", zap.Error(err))
			continue
		}
		if errors.Is(err, workwithapimodels.ErrNoRespAPI) {
			w.log.Error("no responses", zap.Error(err))
			continue
		}
		if errors.Is(err, workwithapimodels.ErrRequestCount) {
			ticker.Reset(time.Second * time.Duration(resGetOrdersAccrual.TimeRetryAfter))
		}
		if err = w.strg.UpdateOrdersStatusAndAccrual(resGetOrdersAccrual); err != nil {
			w.log.Error("error when changing data in the data base table orders/order_accrual", zap.Error(err))
			continue
		}
		usersOrdersAccrualList, err := w.strg.GetCalculatedUsers()
		if err != nil && !errors.Is(err, workwithapimodels.ErrNoUsers) {
			w.log.Error("error when receiving users and points for their calculated orders", zap.Error(err))
			continue
		}
		if errors.Is(err, workwithapimodels.ErrNoUsers) {
			w.log.Error("no users", zap.Error(err))
			continue
		}
		err = w.strg.UpdateBalance(usersOrdersAccrualList)
		if err != nil {
			w.log.Error("error update balance or delete orders", zap.Error(err))
		}
	}
}
