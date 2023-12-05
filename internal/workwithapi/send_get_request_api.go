package workwithapi

import (
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	workwithapimodels "gofermart/internal/models/work_with_api_models"
	"time"
)

func (w *WorkAPI) getOrderAccrual(unAccrualOrdersList []string) (*workwithapimodels.ResGetOrderAccrual, error) {
	resGetOrdersAccrual, err := w.sendGetRequestAPI(unAccrualOrdersList)
	if err != nil &&
		!errors.Is(err, workwithapimodels.ErrRequestCount) &&
		!errors.Is(err, workwithapimodels.ErrNoRespAPI) {
		logger.Error("error when receiving order accrual statuses", zap.Error(err))
		return nil, err
	}
	if errors.Is(err, workwithapimodels.ErrNoRespAPI) {
		logger.Error("no responses", zap.Error(err))
		return nil, err
	}
	if errors.Is(err, workwithapimodels.ErrRequestCount) {
		w.ticker.Reset(time.Second * time.Duration(resGetOrdersAccrual.TimeRetryAfter))
		return nil, err
	}

	return resGetOrdersAccrual, nil
}
