package router

import (
	"gofermart/internal/compress"
	"gofermart/internal/cookie"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func NewRouter(hndlr allhandlers.Handlers, cookies cookie.Cookie) (chi.Router, error) {
	r := chi.NewRouter()

    // Настройка CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:5173"},
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"}, 
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true, 
        MaxAge: 86400, 
    })

    r.Use(c.Handler)

    r.Use(compress.MiddlewareCompress(), logger.MiddlewareHandlerLog())

    r.Route("/api/user/", func(r chi.Router) {
        r.Route("/orders", func(r chi.Router) {
            r.With(cookies.MiddlewareCheckCookie()).Get("/", hndlr.GetOrders)
            r.With(cookies.MiddlewareCheckCookie()).Post("/", hndlr.PostOrders)
        })

        r.With(cookies.MiddlewareCheckCookie()).Get("/balance", hndlr.Balance)
        r.With(cookies.MiddlewareCheckCookie()).Get("/withdrawals", hndlr.Withdrawals)

        r.Post("/register", hndlr.Register)
        r.Post("/login", hndlr.Login)
        r.With(cookies.MiddlewareCheckCookie()).Post("/balance/withdraw", hndlr.WithDraw)
    })

    return r, nil
}
