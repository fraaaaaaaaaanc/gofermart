package workwithapi

import (
	"go.uber.org/zap"
	"gofermart/internal/storage"
)

type WorkAPI struct {
	log  *zap.Logger
	strg *storage.Storage
}

func NewWorkAPI(log *zap.Logger, strg *storage.Storage) *WorkAPI {
	workAPI := &WorkAPI{
		log:  log,
		strg: strg,
	}

	//go workAPI.registerOrders()
	go workAPI.getOrdersAccrual()
	return workAPI
}
