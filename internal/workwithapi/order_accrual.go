package workwithapi

import (
	"go.uber.org/zap"
	"gofermart/internal/logger"
)

func (w *WorkAPI) calculationOrdersAccrual() {
	for range w.ticker.C {
		err := w.processingUnAccrualOrders()
		if err != nil {
			logger.Error("error in calculating order bonuses", zap.Error(err))
		}
	}
}
