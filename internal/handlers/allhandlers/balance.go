package allhandlers

import (
	"encoding/json"
	"gofermart/internal/logger"
	cookiemodels "gofermart/internal/models/cookie_models"
	"net/http"
)

func (h *Handlers) Balance(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(cookiemodels.UserID).(int)
	respUserBalance, err := h.strg.GetUserBalance(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.With(userID, err, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dec := json.NewEncoder(w)
	if err = dec.Encode(respUserBalance); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.With(userID, err, r)
		return
	}
	logger.With(userID, nil, r)
}
