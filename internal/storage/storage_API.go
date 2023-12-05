package storage

import (
	"context"
	"database/sql"
	workwithapimodels "gofermart/internal/models/work_with_api_models"
)

type StorageAPI interface {
	GetAllUnAccrualOrders() ([]string, error)
	GetCalculatedUsers() ([]workwithapimodels.UsersOrdersAccrual, error)
	UpdateOrdersStatus(ctx context.Context, tx *sql.Tx, respGetRequestList workwithapimodels.RespGetRequest) error
	UpdateOrdersAccrual(ctx context.Context, tx *sql.Tx, respGetRequestList workwithapimodels.RespGetRequest) error
	UpdateBalance(ctx context.Context, tx *sql.Tx, usersOrdersAccrualList []workwithapimodels.UsersOrdersAccrual) error
	DeleteOrders(ctx context.Context, tx *sql.Tx) error
	InTransaction(parentsCtx context.Context, f func(ctx context.Context, tx *sql.Tx) error) error
}
