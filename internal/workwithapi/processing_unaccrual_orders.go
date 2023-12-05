package workwithapi

func (w *WorkAPI) processingUnAccrualOrders() error {
	unAccrualOrdersNumberList, err := w.getAllUnAccrualOrders()
	if err != nil {
		return err
	}

	resGetOrderAccrual, err := w.getOrderAccrual(unAccrualOrdersNumberList)
	if err != nil {
		return err
	}

	err = w.updateOrdersStatusAndAccrual(resGetOrderAccrual)
	if err != nil {
		return err
	}

	err = w.userCalculation()

	return err
}
