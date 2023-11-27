package allhandlers

import (
	"github.com/go-playground/validator"
	"gofermart/internal/storage"
)

type Handlers struct {
	validator         *validator.Validate
	strg              storage.StorageMock
	secretKeyJWTToken string
}

func NewHandlers(storage storage.StorageMock, secretKeyJWTToken string) Handlers {
	valid := validator.New()
	return Handlers{
		strg:              storage,
		validator:         valid,
		secretKeyJWTToken: secretKeyJWTToken,
	}
}
