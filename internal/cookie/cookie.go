package cookie

import (
	"context"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	"gofermart/internal/models/cookie_models"
	"net/http"
	"time"
)

const timeLiveToken = time.Hour * 6

func (c Cookie) buildJwtString(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &cookiemodels.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(timeLiveToken)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(c.secretKeyJWTToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (c Cookie) NewUserCookie(userID int) (*http.Cookie, error) {
	tokenString, err := c.buildJwtString(userID)
	if err != nil {
		return nil, err
	}
	return &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   7200,
		HttpOnly: true,
	}, nil
}

func (c Cookie) getUserIDCookie(tokenString string) (int, error) {
	valid := validator.New()
	claims := &cookiemodels.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.secretKeyJWTToken), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, cookiemodels.ErrCookieIsNotValid
	}
	if err = valid.Struct(claims); err != nil {
		return 0, cookiemodels.ErrAbsentUserID
	}
	return claims.UserID, nil
}

func (c Cookie) MiddlewareCheckCookie() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := r.Cookie("Authorization")
			if err != nil {
				http.Error(w, "there is no authorization cookie_models", http.StatusUnauthorized)
				logger.Error("the r.cookie_models(\"Authorization\") parameter is missing", zap.Error(err))
				return
			}
			userID, err := c.getUserIDCookie(tokenString.Value)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "error working with the authorization token", http.StatusUnauthorized)
				logger.Error("an error occurred while working with the authorization token", zap.Error(err))
				return
			}
			ctx := context.WithValue(r.Context(), cookiemodels.UserID, userID)
			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
	}
}
