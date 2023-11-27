package allhandlers

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"net/http"
)

func (h *Handlers) GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(cookiemodels.UserID).(int)
	respGetOrders, err := h.strg.GetAllUserOrders(userID)
	if err != nil && !errors.Is(err, handlersmodels.ErrTheAreNoOrders) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.Error("error when working with the database", zap.Error(err))
		return
	}

	if errors.Is(err, handlersmodels.ErrTheAreNoOrders) {
		http.Error(w, "orders for this user have not been placed yet", http.StatusNoContent)
		logger.Error("the list of bonus points deductions for this user is empty", zap.Error(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	if err = enc.Encode(respGetOrders); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.Error("error forming the response", zap.Error(err))
		return
	}
}
