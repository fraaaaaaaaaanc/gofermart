package allhandlers

import (
	"context"
	"database/sql"
	"errors"
	"gofermart/internal/logger"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/utils"
	"io"
	"net/http"
)

func (h *Handlers) PostOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(cookiemodels.UserID).(int)

	orderNumber, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading the request body", http.StatusBadRequest)
		logger.With(userID, err, r)
		return
	}
	defer r.Body.Close()

	if err = utils.IsLuhnValid(string(orderNumber)); err != nil {
		http.Error(w, "the number ordered did not pass verification", http.StatusUnprocessableEntity)
		logger.With(userID, err, r)
		return
	}

	reqOrder := &handlersmodels.ReqOrder{
		OrderNumber: string(orderNumber),
		UserID:      userID,
	}
	err = h.strg.InTransaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		reqOrder, err := h.strg.AddNewOrder(ctx, tx, reqOrder)
		if err != nil {
			return err
		}
		err = h.strg.AddNewOrderAccrual(ctx, tx, reqOrder)
		return err
	})
	if err != nil &&
		!errors.Is(err, handlersmodels.ErrConflictOrderNumberAnotherUser) &&
		!errors.Is(err, handlersmodels.ErrConflictOrderNumberSameUser) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.With(userID, err, r)
		return
	}

	if errors.Is(err, handlersmodels.ErrConflictOrderNumberAnotherUser) {
		http.Error(w, "order_number uniqueness error", http.StatusConflict)
		logger.With(userID, err, r)
		return
	}

	if errors.Is(err, handlersmodels.ErrConflictOrderNumberSameUser) {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	logger.With(userID, nil, r)
}
