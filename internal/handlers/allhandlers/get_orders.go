package allhandlers

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/models/handlersmodels"
	"net/http"
)

func (h *Handlers) GetOrders(w http.ResponseWriter, r *http.Request) {
	respGetOrders, err := h.strg.GetAllUserOrders(r.Context())
	if err != nil && !errors.Is(err, handlersmodels.ErrTheAreNoOrders) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		h.log.Error("error when working with the database", zap.Error(err))
		return
	}

	if errors.Is(err, handlersmodels.ErrTheAreNoOrders) {
		http.Error(w, "orders for this user have not been placed yet", http.StatusPaymentRequired)
		h.log.Error("the list of bonus points deductions for this user is empty", zap.Error(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	if err = enc.Encode(respGetOrders); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		h.log.Error("error forming the response", zap.Error(err))
		return
	}
}
