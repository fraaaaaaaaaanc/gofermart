package workwithapi

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	"gofermart/internal/models/work_with_api_models"
	"net/http"
	"strconv"
)

func (w *WorkAPI) getOrderAccrual(ordersNumber []string) (*work_with_api_models.ResGetOrderAccrual, error) {
	var respGetOrderAccrual work_with_api_models.ResGetOrderAccrual
	for _, orderNumber := range ordersNumber {
		resp := w.GetRequestOrderAccrual(orderNumber)
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusOK:
			logger.Info("GET request /api/orders/{order} status success")
			var respGetRequest work_with_api_models.RespGetRequest
			dec := json.NewDecoder(resp.Body)
			if err := dec.Decode(&respGetRequest); err != nil {
				logger.Error("error forming the Get Request response "+
					"\"http://localhost:8080/api/orders/{number}\"", zap.Error(err))
				break
			}
			respGetOrderAccrual.RespGetRequestList = append(respGetOrderAccrual.RespGetRequestList, respGetRequest)
		case http.StatusNoContent:
			logger.Error("order number sent by Get Request \"http://localhost:8080/api/orders/{number}\" " +
				"not registered in the calculation system")
		case http.StatusTooManyRequests:
			timeRetryAfterString := resp.Header.Get("Retry-After")
			timeRetryAfter, err := strconv.Atoi(timeRetryAfterString)
			if err != nil {
				logger.Error("error converting timeRetryAfterString to int type", zap.Error(err))
				return &respGetOrderAccrual, err
			}
			respGetOrderAccrual.TimeRetryAfter = timeRetryAfter
			return &respGetOrderAccrual, work_with_api_models.ErrRequestCount
		case http.StatusInternalServerError:
			logger.Error("internal server error when sending GET request " +
				"\"http://localhost:8080/api/orders/{number}\"")
		}
	}
	fmt.Println(&respGetOrderAccrual.RespGetRequestList)
	if respGetOrderAccrual.RespGetRequestList == nil {
		return nil, work_with_api_models.ErrNoRespAPI
	}
	return &respGetOrderAccrual, nil
}
