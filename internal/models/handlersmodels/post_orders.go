package handlersmodels

import (
	"context"
	"errors"
)

var ErrConflictOrderNumberAnotherUser = errors.New("data conflict, the order number sent by the user already exists in " +
	"the storage (added by another user)")

var ErrConflictOrderNumberSameUser = errors.New("data conflict, the order number sent by the user already exists in " +
	"the storage (added by same user)")

type ReqOrder struct {
	OrderNumber string
	UserID      int
	OrderStatus string
	Ctx         context.Context
}

type OrderInfo struct {
	OrderNumber string  `json:"order"`
	Goods       []Goods `json:"goods"`
}

type Goods struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
