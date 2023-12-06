package allhandlers

import (
	"encoding/json"
	"errors"
	"gofermart/internal/logger"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/utils"
	"net/http"
)

func (h *Handlers) WithDraw(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(cookiemodels.UserID).(int)

	var reqWithdraw handlersmodels.ReqWithdraw
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&reqWithdraw); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		logger.With(userID, err, r)
		return
	}

	if err := h.validator.Struct(reqWithdraw); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		logger.With(userID, err, r)
		return
	}

	if err := utils.IsLuhnValid(reqWithdraw.OrderNumber); err != nil {
		http.Error(w, "the number ordered did not pass verification", http.StatusUnprocessableEntity)
		logger.With(userID, err, r)
		return
	}

	reqWithdraw.UserID = userID
	err := h.strg.ProcessingDebitingFunds(r.Context(), reqWithdraw)
	if err != nil {
		switch true {
		case errors.Is(err, handlersmodels.ErrDuplicateOrderNumber):
			http.Error(w, "this order number already exists", http.StatusUnprocessableEntity)
			logger.With(userID, err, r)
			return
		case errors.Is(err, handlersmodels.ErrNegativeBalanceValue):
			http.Error(w, "there are not enough funds on the balance sheet to write off", http.StatusPaymentRequired)
			logger.With(userID, err, r)
			return
		case errors.Is(err, handlersmodels.ErrDuplicateOrderNumberHistoryBalance):
			http.Error(w, "funds have already been debited from the bonus account for this order number",
				http.StatusUnprocessableEntity)
			logger.With(userID, err, r)
			return
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			logger.With(userID, err, r)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	logger.With(userID, nil, r)
}
