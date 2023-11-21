package workwithapi

import (
	"gofermart/internal/storage"
)

type WorkAPI struct {
	accrualSystemAddress string
	strg                 storage.StorageMock
}

func NewWorkAPI(storage storage.StorageMock, accrualSystemAddress string) *WorkAPI {
	workAPI := &WorkAPI{
		accrualSystemAddress: accrualSystemAddress,
		strg:                 storage,
	}

	//go workAPI.registerOrders()
	go workAPI.getOrdersAccrual()
	return workAPI
}
