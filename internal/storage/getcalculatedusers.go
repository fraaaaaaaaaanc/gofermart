package storage

import (
	"gofermart/internal/models/orderstatuses"
	"gofermart/internal/models/workwithapimodels"
)

func (s *Storage) GetCalculatedUsers() ([]workwithapimodels.UsersOrdersAccrual, error) {
	rows, err := s.DB.Query("SELECT user_id, accrual FROM order_accrual "+
		"WHERE order_status_accrual = $1",
		orderstatuses.PROCESSED)
	if err != nil {
		return nil, err
	}

	var usersOrdersAccrualList []workwithapimodels.UsersOrdersAccrual
	for rows.Next() {
		var usersOrdersAccrual workwithapimodels.UsersOrdersAccrual
		if err = rows.Scan(&usersOrdersAccrual.UserID, &usersOrdersAccrual.OrderAccrual); err != nil {
			return nil, err
		}

		usersOrdersAccrualList = append(usersOrdersAccrualList, usersOrdersAccrual)
	}

	if usersOrdersAccrualList == nil {
		return nil, workwithapimodels.ErrNoUsers
	}

	return usersOrdersAccrualList, nil
}
