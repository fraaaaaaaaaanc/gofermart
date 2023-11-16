package storage

import (
	"context"
	"fmt"
	"time"
)

func (s *Storage) UpdateOrderStatus(orderNumber, status string) error {
	newCtx, cansel := context.WithTimeout(context.Background(), time.Second*1)
	defer cansel()

	_, err := s.DB.ExecContext(newCtx,
		"UPDATE orders "+
			"Set order_status = $1 "+
			"WHERE order_number = $2",
		status, orderNumber)

	fmt.Println(err)
	return err
}
