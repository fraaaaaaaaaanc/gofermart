package router

import (
	"github.com/go-chi/chi"
	"gofermart/internal/handlers/allhandlers"
)

func NewRouter(hndlr allhandlers.Handlers) (chi.Router, error) {
	r := chi.NewRouter()

	r.Route("/api/user/", func(r chi.Router) {
		r.Route("/orders", func(r chi.Router) {
			r.Get("/", hndlr.GetOrders)
			r.Post("/", hndlr.PostOrders)
		})

		r.Get("/balance", hndlr.Balance)
		r.Get("/withdrawals", hndlr.GetOrders)

		r.Post("/register", hndlr.Register)
		r.Post("/login", hndlr.Login)
		r.Post("/balance/withdraw", hndlr.Balance)
	})

	return r, nil
}
