package storagegofermart

import (
	"context"
	"database/sql"
	"gofermart/internal/models/work_with_api_models"
)

func (s *Storage) UpdateOrdersStatus(ctx context.Context, tx *sql.Tx, respGetRequestList workwithapimodels.RespGetRequest) error {
	_, err := tx.ExecContext(ctx,
		"UPDATE orders SET order_status = $1, accrual = $2 "+
			"WHERE order_number = $3",
		respGetRequestList.OrderStatus, respGetRequestList.OrderAccrual, respGetRequestList.OrderNumber)

	return err
}
