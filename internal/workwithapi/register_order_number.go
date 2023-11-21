package workwithapi

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	"gofermart/internal/models/work_with_api_models"
	"net/http"
)

func (w *WorkAPI) RegisterOrderNumber(ordersInfo []*work_with_api_models.UnRegisterOrders) {
	for _, orderInfo := range ordersInfo {
		orderInfoJSON, err := json.Marshal(orderInfo)
		if err != nil {
			logger.Error("error creating the request body", zap.Error(err))
			continue
		}
		resp, err := http.Post("http://localhost:8080/api/orders",
			"application/json",
			bytes.NewBuffer(orderInfoJSON))
		defer func() {
			if err = resp.Body.Close(); err != nil {
				logger.Error("request body closing error", zap.Error(err))
			}
		}()

		if err != nil {
			logger.Error("error sending a POST http://localhost:8080/api/orders request to the bonus points "+
				"calculation server API", zap.Error(err))
			continue
		}
		switch resp.StatusCode {
		case http.StatusAccepted:
			logger.Info("order sent by POST http://localhost:8080/api/orders successfully accepted for processing")
			err := w.strg.UpdateOrderStatus(orderInfo.OrderNumber)
			if err != nil {
				logger.Error("error changing the order status in the database table after its registration in " +
					"the external API")
			}
			logger.Info("the order status has been successfully changed")
		case http.StatusBadRequest:
			logger.Error("invalid POST request format http://localhost:8080/api/orders")
		case http.StatusConflict:
			logger.Error("order sent via POST http://localhost:8080/api/orders already accepted for processing")
		case http.StatusInternalServerError:
			logger.Error("internal server error when sending a POST request http://localhost:8080/api/orders")
		}
	}
}
