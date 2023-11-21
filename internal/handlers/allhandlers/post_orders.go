package allhandlers

import (
	"errors"
	"go.uber.org/zap"
	cookiemodels "gofermart/internal/models/cookie"
	"gofermart/internal/models/handlersmodels"
	"gofermart/internal/models/orderstatuses"
	"gofermart/internal/utils"
	"io"
	"net/http"
)

func (h *Handlers) PostOrders(w http.ResponseWriter, r *http.Request) {
	orderNumber, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading the request body", http.StatusBadRequest)
		h.log.Error("invalid request format", zap.Error(err))
		return
	}
	defer r.Body.Close()

	if err = utils.IsLuhnValid(string(orderNumber)); err != nil {
		http.Error(w, "the number ordered did not pass verification", http.StatusUnprocessableEntity)
		h.log.Error("invalid order number format", zap.Error(err))
		return
	}

	userID := r.Context().Value(cookiemodels.UserID).(int)
	reqOrder := &handlersmodels.ReqOrder{
		OrderStatus: orderstatuses.NEW,
		OrderNumber: string(orderNumber),
		UserID:      userID,
		Ctx:         r.Context(),
	}
	err = h.strg.AddNewOrder(reqOrder)
	if err != nil &&
		!errors.Is(err, handlersmodels.ErrConflictOrderNumberAnotherUser) &&
		!errors.Is(err, handlersmodels.ErrConflictOrderNumberSameUser) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		h.log.Error("error when adding the user's user_id and order_number to the database", zap.Error(err))
		return
	}

	if errors.Is(err, handlersmodels.ErrConflictOrderNumberAnotherUser) {
		http.Error(w, "order_number uniqueness error", http.StatusConflict)
		h.log.Error("the order_number sent by the user already exists in the database", zap.Error(err))
		return
	}

	if errors.Is(err, handlersmodels.ErrConflictOrderNumberSameUser) {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
