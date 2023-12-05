package storage

import (
	"context"
	"database/sql"
	"gofermart/internal/models/handlers_models"
)

//go:generate mockgen -source=storage_gofermat.go -destination=mock/mock.go -package=mock

type StorageGofermart interface {
	AddHistoryBalance(ctx context.Context, tx *sql.Tx, reqWithdraw handlersmodels.ReqWithdraw) error
	AddNewOrder(ctx context.Context, tx *sql.Tx, reqOrder *handlersmodels.ReqOrder) (*handlersmodels.ReqOrder, error)
	AddNewOrderAccrual(ctx context.Context, tx *sql.Tx, reqOrder *handlersmodels.ReqOrder) error
	AddNewUser(ctx context.Context, tx *sql.Tx, reqRegister *handlersmodels.RequestRegister) (int, error)
	AddNewUserBalance(ctx context.Context, tx *sql.Tx, userID int) error
	CheckOrderNumber(ctx context.Context, tx *sql.Tx, orderNumber string) error
	CheckUserLoginData(reqLogin *handlersmodels.RequestLogin) (*handlersmodels.ResultLogin, error)
	GetAllHistoryBalance(ctx context.Context, userID int) ([]handlersmodels.RespWithdrawalsHistory, error)
	GetAllUserOrders(ctx context.Context, userID int) ([]handlersmodels.RespGetOrders, error)
	GetUserBalance(ctx context.Context) (*handlersmodels.RespUserBalance, error)
	WithdrawBalance(ctx context.Context, tx *sql.Tx, reqWithdraw handlersmodels.ReqWithdraw) error
	InTransaction(parentsCtx context.Context, f func(ctx context.Context, tx *sql.Tx) error) error
	CloseDB()
}
