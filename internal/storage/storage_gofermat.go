package storage

import (
	"context"
	"database/sql"
	"gofermart/internal/models/handlers_models"
)

//go:generate mockgen -source=storage_gofermat.go -destination=mock/mock.go -package=mock

type StorageGofermart interface {
	AddNewOrderAndAccrual(ctxRequest context.Context, reqOrder *handlersmodels.ReqOrder) error
	AddNewUserAndBalance(ctxRequest context.Context, reqRegister handlersmodels.RequestRegister) (int, error)
	CheckUserLoginData(reqLogin *handlersmodels.RequestLogin) (*handlersmodels.ResultLogin, error)
	GetAllHistoryBalance(ctx context.Context, userID int) ([]handlersmodels.RespWithdrawalsHistory, error)
	GetAllUserOrders(ctx context.Context, userID int) ([]handlersmodels.RespGetOrders, error)
	GetUserBalance(ctx context.Context) (*handlersmodels.RespUserBalance, error)
	ProcessingDebitingFunds(ctxRequest context.Context, reqWithdraw handlersmodels.ReqWithdraw) error
	InTransaction(parentsCtx context.Context, f func(ctx context.Context, tx *sql.Tx) error) error
	CloseDB()
}
