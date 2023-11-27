package workwithapimodels

import (
	"errors"
	"github.com/shopspring/decimal"
)

var ErrRequestCount = errors.New("exceeded the number of requests to the service")

var ErrNoOrdersForRegistration = errors.New("there are no orders for registration")

var ErrNoOrdersForAcrrual = errors.New("there are no orders for accrual")

var ErrNoRespAPI = errors.New("the external server did not give positive results")

var ErrNoUsers = errors.New("")

type RespGetRequest struct {
	OrderNumber  string          `json:"order"`
	OrderStatus  string          `json:"status"`
	OrderAccrual decimal.Decimal `json:"accrual"`
}

type ResGetOrderAccrual struct {
	TimeRetryAfter     int
	RespGetRequestList []RespGetRequest
}

type UsersOrdersAccrual struct {
	UserID       int
	OrderAccrual decimal.Decimal
}

//type OrdersInfo struct {
//	OrdersNumbers  []string
//	OrdersStatuses []string
//	OrdersAccruals []decimal.Decimal
//}
