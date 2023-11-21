package router

import (
	"github.com/go-chi/chi"
	"gofermart/internal/compress"
	"gofermart/internal/cookie"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
)

func NewRouter(hndlr allhandlers.Handlers, secretKey string) (chi.Router, error) {
	r := chi.NewRouter()

	r.Use(compress.MiddlewareCompress(),
		logger.MiddlewareHandlerLog())

	r.Route("/api/user/", func(r chi.Router) {
		r.Route("/orders", func(r chi.Router) {
			r.With(cookie.MiddlewareCheckCookie(secretKey)).Get("/", hndlr.GetOrders)
			r.With(cookie.MiddlewareCheckCookie(secretKey)).Post("/", hndlr.PostOrders)
		})

		r.With(cookie.MiddlewareCheckCookie(secretKey)).Get("/balance", hndlr.Balance)
		r.With(cookie.MiddlewareCheckCookie(secretKey)).Get("/withdrawals", hndlr.Withdrawals)

		r.Post("/register", hndlr.Register)
		r.Post("/login", hndlr.Login)
		r.With(cookie.MiddlewareCheckCookie(secretKey)).Post("/balance/withdraw", hndlr.WithDraw)
	})

	return r, nil
}
