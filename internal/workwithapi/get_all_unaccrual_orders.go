package workwithapi

import (
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	workwithapimodels "gofermart/internal/models/work_with_api_models"
)

func (w *WorkAPI) getAllUnAccrualOrders() ([]string, error) {
	unAccrualOrdersList, err := w.GetAllUnAccrualOrders()
	if err != nil && !errors.Is(err, workwithapimodels.ErrNoOrdersForAcrrual) {
		logger.Error("error when working with the database at the time of receiving the list of "+
			"outstanding orders", zap.Error(err))
		return nil, err
	}
	if errors.Is(err, workwithapimodels.ErrNoOrdersForAcrrual) {
		logger.Info("no orders", zap.Error(err))
		return nil, err
	}
	return unAccrualOrdersList, err
}
