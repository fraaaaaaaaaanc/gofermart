package storage

import (
	"context"
	"database/sql"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/models/work_with_api_models"
)

type StorageMock interface {
	AddNewOrder(reqOrder *handlersmodels.ReqOrder) error
	AddNewUser(reqChanelRegister *handlersmodels.RequestRegister) (int, error)
	CheckOrderNumber(ctx context.Context, orderNumber string) error
	CheckUserLoginData(reqLogin *handlersmodels.RequestLogin) (*handlersmodels.ResultLogin, error)
	DeleteOrders(tx *sql.Tx) error
	GetAllUnAccrualOrders() ([]string, error)
	GetAllHistoryBalance(userID int) ([]handlersmodels.RespWithdrawalsHistory, error)
	GetAllUserOrders(userID int) ([]handlersmodels.RespGetOrders, error)
	GetCalculatedUsers() ([]workwithapimodels.UsersOrdersAccrual, error)
	GetUserBalance(ctx context.Context) (*handlersmodels.RespUserBalance, error)
	UpdateBalance(usersOrdersAccrualList []workwithapimodels.UsersOrdersAccrual) error
	UpdateOrdersStatusAndAccrual(resGetOrdersAccrual *workwithapimodels.ResGetOrderAccrual) error
	UpdateOrderStatus(orderNumber string) error
	WithdrawBalance(reqWithdraw handlersmodels.ReqWithdraw) error
	CloseDB()
}
