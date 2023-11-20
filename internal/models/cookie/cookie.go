package cookiemodels

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

var ErrCookieIsNotValid = errors.New("token is not valid")

var ErrAbsentUserID = errors.New("the token does not have a field userID")

type ContextKey string

var UserID ContextKey = "UserID"

type Claims struct {
	jwt.RegisteredClaims
	UserID int `validate:"required"`
}
