package handlersmodels

import (
	"errors"
	"time"
)

var ErrTheAreNoWithdraw = errors.New("there are no withdraw for this user")

type RespWithdrawalsHistory struct {
	OrderNumber string    `json:"order"`
	SumWithdraw float64   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}
