package handlersmodels

import (
	"context"
	"errors"
)

var ErrMissingDataInTable = errors.New("this user does not exist")

type (
	RequestLogin struct {
		Login    string `json:"login" validate:"required"`
		Password string `json:"password" validate:"required"`
		Ctx      context.Context
	}

	ResultLogin struct {
		UserID   int
		Password string
	}
)
