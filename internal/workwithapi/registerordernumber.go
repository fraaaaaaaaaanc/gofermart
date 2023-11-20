package workwithapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"gofermart/internal/models/workwithapimodels"
	"net/http"
)

func (w *WorkAPI) RegisterOrderNumber(ordersInfo []*workwithapimodels.UnRegisterOrders) {
	for _, orderInfo := range ordersInfo {
		orderInfoJSON, err := json.Marshal(orderInfo)
		if err != nil {
			w.log.Error("error creating the request body", zap.Error(err))
			continue
		}
		resp, err := http.Post("http://localhost:8080/api/orders",
			"application/json",
			bytes.NewBuffer(orderInfoJSON))
		if err != nil {
			w.log.Error("error sending a POST http://localhost:8080/api/orders request to the bonus points "+
				"calculation server API", zap.Error(err))
			continue
		}
		switch resp.StatusCode {
		case http.StatusAccepted:
			w.log.Info("order sent by POST http://localhost:8080/api/orders successfully accepted for processing")
			fmt.Println(orderInfo.OrderNumber)
			err := w.strg.UpdateOrderStatus(orderInfo.OrderNumber)
			if err != nil {
				w.log.Error("error changing the order status in the database table after its registration in " +
					"the external API")
			}
			w.log.Info("the order status has been successfully changed")
		case http.StatusBadRequest:
			w.log.Error("invalid POST request format http://localhost:8080/api/orders")
		case http.StatusConflict:
			w.log.Error("order sent via POST http://localhost:8080/api/orders already accepted for processing")
		case http.StatusInternalServerError:
			w.log.Error("internal server error when sending a POST request http://localhost:8080/api/orders")
		}
	}
}
