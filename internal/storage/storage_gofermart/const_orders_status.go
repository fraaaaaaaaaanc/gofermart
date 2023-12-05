package storagegofermart

import "time"

type OrderStatus string

const (
	orderStatusNew        OrderStatus = "NEW"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusInvalid    OrderStatus = "INVALID"
	OrderStatusProcessed  OrderStatus = "PROCESSED"
	OrderStatusRegistered OrderStatus = "REGISTERED"
)

// constant running time of the context for database queries
const durationWorkCtx = time.Second * 3
