package allhandlers

import (
	"github.com/go-playground/validator"
	"gofermart/internal/cookie"
	"gofermart/internal/storage"
)

const unLoggedUserID = 0

type Handlers struct {
	validator *validator.Validate
	strg      storage.StorageGofermart
	cookie    cookie.Cookie
}

func NewHandlers(storage storage.StorageGofermart, cookie cookie.Cookie) Handlers {
	valid := validator.New()
	return Handlers{
		strg:      storage,
		validator: valid,
		cookie:    cookie,
	}
}
