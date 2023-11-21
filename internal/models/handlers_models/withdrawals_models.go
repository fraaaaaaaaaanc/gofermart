package handlers_models

import (
	"errors"
)

var ErrTheAreNoWithdraw = errors.New("there are no withdraw for this user")

type RespWithdrawalsHistory struct {
	OrderNumber string  `json:"order"`
	SumWithdraw float64 `json:"sum"`
	ProcessedAt string  `json:"processed_at"`
}
