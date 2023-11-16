package handlersmodels

import (
	"errors"
)

var ErrTheAreNoOrders = errors.New("there are no orders for this user")

type RespGetOrders struct {
	OrderNumber string `json:"number"`
	Status      string `json:"status"`
	Accrual     int    `json:"accrual,omitempty"`
	UploadedAt  string `json:"uploaded_at"`
}
