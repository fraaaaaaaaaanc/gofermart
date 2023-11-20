package workwithapi

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"gofermart/internal/models/workwithapimodels"
	"net/http"
	"strconv"
)

func (w *WorkAPI) getOrderAccrual(ordersNumber []string) (*workwithapimodels.ResGetOrderAccrual, error) {
	var respGetOrderAccrual workwithapimodels.ResGetOrderAccrual
	for _, orderNumber := range ordersNumber {
		resp := w.GetRequestOrderAccrual(orderNumber)
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusOK:
			w.log.Info("GET request http://localhost:8080/api/orders/{order} status success")
			var respGetRequest workwithapimodels.RespGetRequest
			dec := json.NewDecoder(resp.Body)
			if err := dec.Decode(&respGetRequest); err != nil {
				w.log.Error("error forming the Get Request response "+
					"\"http://localhost:8080/api/orders/{number}\"", zap.Error(err))
				break
			}
			respGetOrderAccrual.RespGetRequestList = append(respGetOrderAccrual.RespGetRequestList, respGetRequest)
			//ordersInfo.OrdersNumbers = append(ordersInfo.OrdersNumbers, respGetRequest.OrderNumber)
			//ordersInfo.OrdersStatuses = append(ordersInfo.OrdersStatuses, respGetRequest.OrderStatus)
			//ordersInfo.OrdersAccruals = append(ordersInfo.OrdersAccruals, respGetRequest.OrderAccrual)
		case http.StatusNoContent:
			w.log.Error("order number sent by Get Request \"http://localhost:8080/api/orders/{number}\" " +
				"not registered in the calculation system")
		case http.StatusTooManyRequests:
			timeRetryAfterString := resp.Header.Get("Retry-After")
			timeRetryAfter, err := strconv.Atoi(timeRetryAfterString)
			if err != nil {
				w.log.Error("error converting timeRetryAfterString to int type", zap.Error(err))
				return &respGetOrderAccrual, err
			}
			respGetOrderAccrual.TimeRetryAfter = timeRetryAfter
			return &respGetOrderAccrual, workwithapimodels.ErrRequestCount
		case http.StatusInternalServerError:
			w.log.Error("internal server error when sending GET request " +
				"\"http://localhost:8080/api/orders/{number}\"")
		}
	}
	fmt.Println(&respGetOrderAccrual.RespGetRequestList)
	if respGetOrderAccrual.RespGetRequestList == nil {
		return nil, workwithapimodels.ErrNoRespAPI
	}
	return &respGetOrderAccrual, nil
}
