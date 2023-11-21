package allhandlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	"net/http"
)

func (h *Handlers) Balance(w http.ResponseWriter, r *http.Request) {
	respUserBalance, err := h.strg.GetUserBalance(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.Error("an error occurred while working with the database and generating the response", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dec := json.NewEncoder(w)
	if err = dec.Encode(respUserBalance); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.Error("error forming the response", zap.Error(err))
		return
	}
}
