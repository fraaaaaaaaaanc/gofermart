package workwithapi

import (
	"gofermart/internal/storage"
	"time"
)

// FrequencyRequest constant response time of the request to the external API
const FrequencyRequest = time.Second * 5

type WorkAPI struct {
	accrualSystemAddress string
	ticker               *time.Ticker
	storage.StorageAPI
}

func NewWorkAPI(storage storage.StorageAPI, accrualSystemAddress string) *WorkAPI {
	ticker := time.NewTicker(FrequencyRequest)
	workAPI := &WorkAPI{
		accrualSystemAddress: accrualSystemAddress,
		ticker:               ticker,
		StorageAPI:           storage,
	}

	go workAPI.calculationOrdersAccrual()
	return workAPI
}
