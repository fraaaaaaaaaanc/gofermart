package allhandlers

import (
	"encoding/json"
	"errors"
	"gofermart/internal/logger"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"net/http"
)

func (h *Handlers) Withdrawals(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(cookiemodels.UserID).(int)

	respWithdrawalsHistory, err := h.strg.GetAllHistoryBalance(r.Context(), userID)
	if err != nil && !errors.Is(err, handlersmodels.ErrTheAreNoWithdraw) {
		http.Error(w, "bonus points have not been debited from this account before", http.StatusNoContent)
		logger.With(userID, err, r)
		return
	}

	if errors.Is(err, handlersmodels.ErrTheAreNoWithdraw) {
		http.Error(w, "bonus points have not been debited from this account before", http.StatusNoContent)
		logger.With(userID, err, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	if err = enc.Encode(respWithdrawalsHistory); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.With(userID, err, r)
		return
	}
	logger.With(userID, nil, r)
}
