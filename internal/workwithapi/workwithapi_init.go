package workwithapi

import (
	"go.uber.org/zap"
	"gofermart/internal/storage"
)

type WorkAPI struct {
	accrualSystemAddress string
	log                  *zap.Logger
	strg                 *storage.Storage
}

func NewWorkAPI(log *zap.Logger, strg *storage.Storage, accrualSystemAddress string) *WorkAPI {
	workAPI := &WorkAPI{
		accrualSystemAddress: accrualSystemAddress,
		log:                  log,
		strg:                 strg,
	}

	//go workAPI.registerOrders()
	go workAPI.getOrdersAccrual()
	return workAPI
}
