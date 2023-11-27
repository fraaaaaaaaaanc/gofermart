package workwithapi

import (
	"go.uber.org/zap"
	"gofermart/internal/logger"
	"net/http"
)

func (w *WorkAPI) GetRequestOrderAccrual(orderNumber string) *http.Response {
	resp, err := http.Get(w.accrualSystemAddress + "/api/orders/" + orderNumber)
	if err != nil {
		logger.Error("error when interacting with an external API while sending a GET request "+
			"\"http://localhost:8080/api/orders/{order}\"", zap.Error(err))
		return nil
	}
	return resp
}
