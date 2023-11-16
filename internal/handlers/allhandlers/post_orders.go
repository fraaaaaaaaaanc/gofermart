package allhandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	cookiemodels "gofermart/internal/models/cookie"
	"gofermart/internal/models/handlersmodels"
	"gofermart/internal/utils"
	"net/http"
)

func (h *Handlers) PostOrders(w http.ResponseWriter, r *http.Request) {
	var orderInfo handlersmodels.OrderInfo
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&orderInfo); err != nil {
		http.Error(w, "error reading the request body", http.StatusBadRequest)
		h.log.Error("invalid request format", zap.Error(err))
		return
	}
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, "error reading the request body", http.StatusBadRequest)
	//	h.log.Error("invalid request format", zap.Error(err))
	//	return
	//}
	defer r.Body.Close()
	fmt.Println(orderInfo)

	if err := utils.IsLuhnValid(orderInfo.OrderNumber); err != nil {
		http.Error(w, "the number ordered did not pass verification", http.StatusUnprocessableEntity)
		h.log.Error("invalid order number format", zap.Error(err))
		return
	}

	reqOrders := &handlersmodels.ReqOrder{
		OrderNumber: orderInfo.OrderNumber,
		UserID:      r.Context().Value(cookiemodels.UserID).(int),
		OrderStatus: "NEW",
		Ctx:         r.Context(),
	}

	err := h.strg.AddNewOrder(reqOrders)
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

	h.Ch <- &orderInfo
	w.WriteHeader(http.StatusAccepted)
}
