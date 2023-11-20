package allhandlers

import (
	"github.com/go-playground/validator"
	"go.uber.org/zap"
	"gofermart/internal/storage"
)

type Handlers struct {
	log               *zap.Logger
	validator         *validator.Validate
	strg              *storage.Storage
	secretKeyJWTToken string
}

func NewHandlers(logger *zap.Logger, storage *storage.Storage, secretKeyJWTToken string) Handlers {
	valid := validator.New()
	return Handlers{
		log:               logger,
		strg:              storage,
		validator:         valid,
		secretKeyJWTToken: secretKeyJWTToken,
	}
}
