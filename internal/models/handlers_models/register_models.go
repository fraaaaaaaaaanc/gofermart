package handlers_models

import (
	"context"
	"errors"
)

var ErrConflictLoginRegister = errors.New("data conflict, the login sent by the user already exists in the " +
	"repository")

type RequestRegister struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
	Ctx      context.Context
}
