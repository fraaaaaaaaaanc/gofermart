package allhandlers

import (
	"encoding/json"
	"errors"
	"gofermart/internal/logger"
	"gofermart/internal/models/handlers_models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin *handlersmodels.RequestLogin
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&reqLogin); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		logger.With(unLoggedUserID, err, r)
		return
	}

	if err := h.validator.Struct(reqLogin); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		logger.With(unLoggedUserID, err, r)
		return
	}

	reqLogin.Ctx = r.Context()
	resLogin, err := h.strg.CheckUserLoginData(reqLogin)

	if err != nil && !errors.Is(err, handlersmodels.ErrMissingDataInTable) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.With(unLoggedUserID, err, r)
		return
	}

	if errors.Is(err, handlersmodels.ErrMissingDataInTable) {
		http.Error(w, "authentication error: login not found", http.StatusUnauthorized)
		logger.With(unLoggedUserID, err, r)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(resLogin.Password), []byte(reqLogin.Password)); err != nil {
		http.Error(w, "authentication error: invalid password", http.StatusUnauthorized)
		logger.With(unLoggedUserID, err, r)
		return
	}

	newCookie, err := h.cookie.NewUserCookie(resLogin.UserID)
	if err != nil {
		http.Error(w, "cookie_models creation error", http.StatusInternalServerError)
		logger.With(unLoggedUserID, err, r)
		return
	}

	http.SetCookie(w, newCookie)
	w.WriteHeader(http.StatusOK)
	logger.With(resLogin.UserID, nil, r)
}
