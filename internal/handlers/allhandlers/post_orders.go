package allhandlers

import (
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/models/orderstatuses"
	"gofermart/internal/utils"
	"io"
	"net/http"
)

func (h *Handlers) PostOrders(w http.ResponseWriter, r *http.Request) {
	//var orderInfo handlers_models.OrderInfo
	//dec := json.NewDecoder(r.Body)
	//if err := dec.Decode(&orderInfo); err != nil {
	//	http.Error(w, "error reading the request body", http.StatusBadRequest)
	//	h.log.Error("invalid request format", zap.Error(err))
	//	return
	//}
	orderNumber, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading the request body", http.StatusBadRequest)
		logger.Error("invalid request format", zap.Error(err))
		return
	}
	defer r.Body.Close()

	if err = utils.IsLuhnValid(string(orderNumber)); err != nil {
		http.Error(w, "the number ordered did not pass verification", http.StatusUnprocessableEntity)
		logger.Error("invalid order number format", zap.Error(err))
		return
	}

	//userID := r.Context().Value(cookiemodels.UserID).(int)
	//orderInfo.UserID = userID
	//reqOrder := &handlers_models.ReqOrder{
	//	OrderStatus: orderstatuses.NEW,
	//	Ctx:         r.Context(),
	//	OrderInfo:   orderInfo,
	//}

	userID := r.Context().Value(cookiemodels.UserID).(int)
	reqOrder := &handlers_models.ReqOrder{
		OrderStatus: orderstatuses.NEW,
		OrderNumber: string(orderNumber),
		UserID:      userID,
		Ctx:         r.Context(),
	}
	err = h.strg.AddNewOrder(reqOrder)
	if err != nil &&
		!errors.Is(err, handlers_models.ErrConflictOrderNumberAnotherUser) &&
		!errors.Is(err, handlers_models.ErrConflictOrderNumberSameUser) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.Error("error when adding the user's user_id and order_number to the database", zap.Error(err))
		return
	}

	if errors.Is(err, handlers_models.ErrConflictOrderNumberAnotherUser) {
		http.Error(w, "order_number uniqueness error", http.StatusConflict)
		logger.Error("the order_number sent by the user already exists in the database", zap.Error(err))
		return
	}

	if errors.Is(err, handlers_models.ErrConflictOrderNumberSameUser) {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
