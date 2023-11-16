package cookie

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"gofermart/internal/models/cookie"
	"net/http"
	"os"
	"time"
)

const timeLiveToken = time.Hour * 6

func BuildJwtString(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &cookiemodels.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(timeLiveToken)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_FOR_COOKIE_TOKEN")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewCookie(userID int) (*http.Cookie, error) {
	tokenString, err := BuildJwtString(userID)
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

func getUserIDCookie(tokenString string) (int, error) {
	valid := validator.New()
	claims := &cookiemodels.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY_FOR_COOKIE_TOKEN")), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, cookiemodels.CookieIsNotValid
	}
	if err = valid.Struct(claims); err != nil {
		return 0, cookiemodels.AbsentUserID
	}
	return claims.UserID, nil
}

func MiddlewareCheckCookie(log *zap.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := r.Cookie("Authorization")
			if err != nil {
				http.Error(w, "there is no authorization cookie", http.StatusUnauthorized)
				log.Error("the r.cookie(\"Authorization\") parameter is missing", zap.Error(err))
				return
			}
			userID, err := getUserIDCookie(tokenString.Value)
			if err != nil {
				http.Error(w, "error working with the authorization token", http.StatusUnauthorized)
				log.Error("an error occurred while working with the authorization token", zap.Error(err))
				return
			}
			ctx := context.WithValue(r.Context(), cookiemodels.UserID, userID)
			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
	}
}
