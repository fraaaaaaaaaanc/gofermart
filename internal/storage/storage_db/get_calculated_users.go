package storagedb

import (
	"gofermart/internal/models/orderstatuses"
	workwithapimodels "gofermart/internal/models/work_with_api_models"
)

func (s *Storage) GetCalculatedUsers() ([]workwithapimodels.UsersOrdersAccrual, error) {
	rows, err := s.db.Query("SELECT user_id, accrual FROM order_accrual "+
		"WHERE order_status_accrual = $1",
		orderstatuses.PROCESSED)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersOrdersAccrualList []workwithapimodels.UsersOrdersAccrual
	for rows.Next() {
		var usersOrdersAccrual workwithapimodels.UsersOrdersAccrual
		if err = rows.Scan(&usersOrdersAccrual.UserID, &usersOrdersAccrual.OrderAccrual); err != nil {
			return nil, err
		}

		usersOrdersAccrualList = append(usersOrdersAccrualList, usersOrdersAccrual)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if usersOrdersAccrualList == nil {
		return nil, workwithapimodels.ErrNoUsers
	}

	return usersOrdersAccrualList, nil
}
