package allhandlers

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gofermart/internal/cookie"
	"gofermart/internal/models/handlersmodels"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var reqRegister handlersmodels.RequestRegister
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&reqRegister); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		h.log.Error(
			"invalid request format, error when transferring data to the structure handlersmodels.RequestRegister",
			zap.Error(err))
		return
	}

	if err := h.validator.Struct(reqRegister); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		h.log.Error(
			"invalid request format, not all fields of the structure were filled in models.RequestRegister",
			zap.Error(err))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		h.log.Error("internal server error when hashing the password", zap.Error(err))
		return
	}

	reqRegister.Password = string(hashedPassword)
	reqRegister.Ctx = r.Context()
	userID, err := h.strg.AddNewUser(&reqRegister)

	if err != nil && !errors.Is(err, handlersmodels.ErrConflictLoginRegister) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		h.log.Error("error when adding the user's username and password to the database", zap.Error(err))
		return
	}

	if errors.Is(err, handlersmodels.ErrConflictLoginRegister) {
		http.Error(w, "login uniqueness error", http.StatusConflict)
		h.log.Error("the login sent by the user already exists in the database", zap.Error(err))
		return
	}

	newCookie, err := cookie.NewCookie(userID)
	if err != nil {
		http.Error(w, "cookie creation error", http.StatusInternalServerError)
		h.log.Error("an error occurred when creating a new cookie", zap.Error(err))
		return
	}

	http.SetCookie(w, newCookie)
	w.WriteHeader(http.StatusOK)
}
