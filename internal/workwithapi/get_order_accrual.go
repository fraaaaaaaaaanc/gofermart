package workwithapi

import (
	"encoding/json"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	"gofermart/internal/models/work_with_api_models"
	"net/http"
	"strconv"
)

func (w *WorkAPI) sendGetRequestAPI(ordersNumber []string) (*workwithapimodels.ResGetOrderAccrual, error) {
	var respGetOrderAccrual workwithapimodels.ResGetOrderAccrual
	for _, orderNumber := range ordersNumber {
		resp := w.GetRequestOrderAccrual(orderNumber)
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusOK:
			logger.Info("GET request /api/orders/{order} status success")
			var respGetRequest workwithapimodels.RespGetRequest
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
			return &respGetOrderAccrual, workwithapimodels.ErrRequestCount
		case http.StatusInternalServerError:
			logger.Error("internal server error when sending GET request " +
				"\"http://localhost:8080/api/orders/{number}\"")
		}
	}
	if respGetOrderAccrual.RespGetRequestList == nil {
		return nil, workwithapimodels.ErrNoRespAPI
	}
	return &respGetOrderAccrual, nil
}
