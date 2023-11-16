package router

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"gofermart/internal/compress"
	"gofermart/internal/cookie"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
)

func NewRouter(hndlr allhandlers.Handlers, log *zap.Logger) (chi.Router, error) {
	r := chi.NewRouter()

	r.Use(compress.MiddlewareCompress(log),
		logger.MiddlewareHandlerLog(log))

	r.Route("/api/user/", func(r chi.Router) {
		r.Route("/orders", func(r chi.Router) {
			r.With(cookie.MiddlewareCheckCookie(log)).Get("/", hndlr.GetOrders)
			r.With(cookie.MiddlewareCheckCookie(log)).Post("/", hndlr.PostOrders)
		})

		r.With(cookie.MiddlewareCheckCookie(log)).Get("/balance", hndlr.Balance)
		r.With(cookie.MiddlewareCheckCookie(log)).Get("/withdrawals", hndlr.Withdrawals)

		r.Post("/register", hndlr.Register)
		r.Post("/login", hndlr.Login)
		r.With(cookie.MiddlewareCheckCookie(log)).Post("/balance/withdraw", hndlr.WithDraw)
	})

	return r, nil
}
