package allhandlers

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/models/handlersmodels"
	"net/http"
)

func (h *Handlers) Withdrawals(w http.ResponseWriter, r *http.Request) {
	respWithdrawalsHistory, err := h.strg.GetAllHistoryBalance(r.Context())
	if err != nil && !errors.Is(err, handlersmodels.ErrTheAreNoWithdraw) {
		http.Error(w, "bonus points have not been debited from this account before", http.StatusNoContent)
		h.log.Error("error forming the response", zap.Error(err))
		return
	}

	if errors.Is(err, handlersmodels.ErrTheAreNoWithdraw) {
		http.Error(w, "bonus points have not been debited from this account before", http.StatusNoContent)
		h.log.Error("error forming the response", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	if err = enc.Encode(respWithdrawalsHistory); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		h.log.Error("error forming the response", zap.Error(err))
		return
	}
}
