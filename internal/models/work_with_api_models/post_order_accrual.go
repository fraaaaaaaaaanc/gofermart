package work_with_api_models

type UnRegisterOrders struct {
	OrderNumber string      `json:"order"`
	OrderInfo   []OrderInfo `json:"goods"`
}

type OrderInfo struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
