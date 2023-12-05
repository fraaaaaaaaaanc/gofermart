package allhandlers

import (
	"encoding/json"
	"errors"
	"gofermart/internal/logger"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"net/http"
)

func (h *Handlers) GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(cookiemodels.UserID).(int)
	respGetOrders, err := h.strg.GetAllUserOrders(r.Context(), userID)
	if err != nil && !errors.Is(err, handlersmodels.ErrTheAreNoOrders) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.With(userID, err, r)
		return
	}

	if errors.Is(err, handlersmodels.ErrTheAreNoOrders) {
		http.Error(w, "orders for this user have not been placed yet", http.StatusNoContent)
		logger.With(userID, err, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	if err = enc.Encode(respGetOrders); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.With(userID, err, r)
		return
	}
	logger.With(userID, nil, r)
}
