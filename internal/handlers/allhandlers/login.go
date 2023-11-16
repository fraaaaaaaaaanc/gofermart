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

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin *handlersmodels.RequestLogin
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&reqLogin); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		h.log.Error(
			"invalid request format, error when transferring data to the structure models.RequestLogin",
			zap.Error(err))
		return
	}

	if err := h.validator.Struct(reqLogin); err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		h.log.Error(
			"invalid request format, not all fields of the structure were filled in models.Login",
			zap.Error(err))
		return
	}

	reqLogin.Ctx = r.Context()
	resLogin, err := h.strg.CheckUserLogin(reqLogin)

	if err != nil && !errors.Is(err, handlersmodels.ErrMissingDataInTable) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		h.log.Error("error when searching for data in the table", zap.Error(err))
		return
	}

	if errors.Is(err, handlersmodels.ErrMissingDataInTable) {
		http.Error(w, "authentication error: login not found", http.StatusUnauthorized)
		h.log.Error("invalid login/password pair", zap.Error(handlersmodels.ErrMissingDataInTable))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(resLogin.Password), []byte(reqLogin.Password)); err != nil {
		http.Error(w, "authentication error: invalid password", http.StatusUnauthorized)
		h.log.Error("invalid login/password pair", zap.Error(err))
		return
	}

	newCookie, err := cookie.NewCookie(resLogin.UserID)
	if err != nil {
		http.Error(w, "cookie creation error", http.StatusInternalServerError)
		h.log.Error("an error occurred when creating a new cookie", zap.Error(err))
		return
	}

	http.SetCookie(w, newCookie)
	w.WriteHeader(http.StatusOK)
}
