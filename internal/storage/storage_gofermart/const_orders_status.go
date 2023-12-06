package storagegofermart

import "time"

type OrderStatus string

const (
	orderStatusNew        OrderStatus = "NEW"
	orderStatusProcessing OrderStatus = "PROCESSING"
	orderStatusInvalid    OrderStatus = "INVALID"
	orderStatusProcessed  OrderStatus = "PROCESSED"
	orderStatusRegistered OrderStatus = "REGISTERED"
)

// constant running time of the context for database queries
const durationWorkCtx = time.Second * 3
