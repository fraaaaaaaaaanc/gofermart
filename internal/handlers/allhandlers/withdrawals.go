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

func (h *Handlers) Withdrawals(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(cookiemodels.UserID).(int)
	respWithdrawalsHistory, err := h.strg.GetAllHistoryBalance(userID)
	if err != nil && !errors.Is(err, handlers_models.ErrTheAreNoWithdraw) {
		http.Error(w, "bonus points have not been debited from this account before", http.StatusNoContent)
		logger.Error("error forming the response", zap.Error(err))
		return
	}

	if errors.Is(err, handlers_models.ErrTheAreNoWithdraw) {
		http.Error(w, "bonus points have not been debited from this account before", http.StatusNoContent)
		logger.Error("error forming the response", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	if err = enc.Encode(respWithdrawalsHistory); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.Error("error forming the response", zap.Error(err))
		return
	}
}
