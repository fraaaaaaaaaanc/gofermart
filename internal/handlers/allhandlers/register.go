package allhandlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"gofermart/internal/logger"
	"gofermart/internal/models/handlers_models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var reqRegister handlersmodels.RequestRegister
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&reqRegister); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		logger.With(unLoggedUserID, err, r)
		return
	}

	if err := h.validator.Struct(reqRegister); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		logger.With(unLoggedUserID, err, r)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.With(unLoggedUserID, err, r)
		return
	}

	reqRegister.Password = string(hashedPassword)
	var userID int
	err = h.strg.InTransaction(r.Context(), func(ctx context.Context, tx *sql.Tx) error {
		userID, err = h.strg.AddNewUser(ctx, tx, &reqRegister)
		if err != nil {
			return err
		}
		err = h.strg.AddNewUserBalance(ctx, tx, userID)
		return err
	})

	if err != nil && !errors.Is(err, handlersmodels.ErrConflictLoginRegister) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.With(unLoggedUserID, err, r)
		return
	}

	if errors.Is(err, handlersmodels.ErrConflictLoginRegister) {
		http.Error(w, "login uniqueness error", http.StatusConflict)
		logger.With(unLoggedUserID, err, r)
		return
	}

	newCookie, err := h.cookie.NewUserCookie(userID)
	if err != nil {
		http.Error(w, "cookie_models creation error", http.StatusInternalServerError)
		logger.With(unLoggedUserID, err, r)
		return
	}

	http.SetCookie(w, newCookie)
	w.WriteHeader(http.StatusOK)
	logger.With(userID, nil, r)
}
