package storage

import (
	"context"
	"database/sql"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/models/work_with_api_models"
)

type StorageMock interface {
	AddNewOrder(reqOrder *handlers_models.ReqOrder) error
	AddNewUser(reqChanelRegister *handlers_models.RequestRegister) (int, error)
	CheckOrderNumber(ctx context.Context, orderNumber string) error
	CheckUserLoginData(reqLogin *handlers_models.RequestLogin) (*handlers_models.ResultLogin, error)
	DeleteOrders(tx *sql.Tx) error
	GetAllUnAccrualOrders() ([]string, error)
	GetAllHistoryBalance(userID int) ([]handlers_models.RespWithdrawalsHistory, error)
	GetAllUserOrders(userID int) ([]handlers_models.RespGetOrders, error)
	GetCalculatedUsers() ([]work_with_api_models.UsersOrdersAccrual, error)
	GetUserBalance(ctx context.Context) (*handlers_models.RespUserBalance, error)
	UpdateBalance(usersOrdersAccrualList []work_with_api_models.UsersOrdersAccrual) error
	UpdateOrdersStatusAndAccrual(resGetOrdersAccrual *work_with_api_models.ResGetOrderAccrual) error
	UpdateOrderStatus(orderNumber string) error
	WithdrawBalance(reqWithdraw handlers_models.ReqWithdraw) error
	CloseDB()
}
