package storagegofermart

import (
	"context"
	"database/sql"
	workwithapimodels "gofermart/internal/models/work_with_api_models"
)

func (s *Storage) UpdateOrdersAccrual(ctx context.Context, tx *sql.Tx, respGetRequestList workwithapimodels.RespGetRequest) error {
	_, err := tx.ExecContext(ctx,
		"UPDATE order_accrual SET order_status_accrual = $1, accrual = $2 "+
			"WHERE order_number = $3",
		respGetRequestList.OrderStatus, respGetRequestList.OrderAccrual, respGetRequestList.OrderNumber)

	return err
}
