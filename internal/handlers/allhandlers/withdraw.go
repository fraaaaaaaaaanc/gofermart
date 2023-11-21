package allhandlers

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/utils"
	"net/http"
)

func (h *Handlers) WithDraw(w http.ResponseWriter, r *http.Request) {
	var reqWithdraw handlers_models.ReqWithdraw
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&reqWithdraw); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		logger.Error(
			"invalid request format, error when transferring data to the structure handlers_models.ReqWithdraw",
			zap.Error(err))
		return
	}

	if err := h.validator.Struct(reqWithdraw); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		logger.Error(
			"invalid request format, not all fields of the structure were filled in handlers_models.ReqWithdraw",
			zap.Error(err))
		return
	}

	if err := utils.IsLuhnValid(reqWithdraw.OrderNumber); err != nil {
		http.Error(w, "the number ordered did not pass verification", http.StatusUnprocessableEntity)
		logger.Error("invalid order number format", zap.Error(err))
		return
	}

	err := h.strg.CheckOrderNumber(r.Context(), reqWithdraw.OrderNumber)
	if err != nil && !errors.Is(err, handlers_models.ErrDuplicateOrderNumber) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.Error("error when working with the database", zap.Error(err))
		return
	}

	if errors.Is(err, handlers_models.ErrDuplicateOrderNumber) {
		http.Error(w, "this order number already exists", http.StatusUnprocessableEntity)
		logger.Error("the order number sent by the user is already in the orders table", zap.Error(err))
		return
	}

	reqWithdraw.Ctx = r.Context()
	err = h.strg.WithdrawBalance(reqWithdraw)
	if err != nil &&
		!errors.Is(err, handlers_models.ErrNegativeBalanceValue) &&
		!errors.Is(err, handlers_models.ErrDuplicateOrderNumberHistoryBalance) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.Error("error when working with the database", zap.Error(err))
		return
	}

	if errors.Is(err, handlers_models.ErrDuplicateOrderNumberHistoryBalance) {
		http.Error(w, "funds have already been debited from the bonus account for this order number",
			http.StatusUnprocessableEntity)
		logger.Error("an error occurred when adding the order number to the history_balance table, "+
			"this order already exists",
			zap.Error(err))
		return
	}

	if errors.Is(err, handlers_models.ErrNegativeBalanceValue) {
		http.Error(w, "there are not enough funds on the balance sheet to write off", http.StatusPaymentRequired)
		logger.Error("error when debiting funds from the bonus account, insufficient funds", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
