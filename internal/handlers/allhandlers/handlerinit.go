package allhandlers

import (
	"github.com/go-playground/validator"
	"go.uber.org/zap"
	"gofermart/internal/models/handlersmodels"
	"gofermart/internal/storage"
)

type Handlers struct {
	log       *zap.Logger
	validator *validator.Validate
	strg      *storage.Storage
	Ch        chan *handlersmodels.OrderInfo
}

func NewHandlers(logger *zap.Logger, storage *storage.Storage) Handlers {
	valid := validator.New()
	return Handlers{
		log:       logger,
		strg:      storage,
		validator: valid,
		Ch:        make(chan *handlersmodels.OrderInfo, 100),
	}
}
