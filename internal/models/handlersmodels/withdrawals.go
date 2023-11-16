package handlersmodels

import (
	"errors"
	"github.com/shopspring/decimal"
	"time"
)

var ErrTheAreNoWithdraw = errors.New("there are no withdraw for this user")

type RespWithdrawalsHistory struct {
	OrderNumber string          `json:"number"`
	SumWithdraw decimal.Decimal `json:"sum"`
	ProcessedAt time.Time       `json:"processed_at"`
}
