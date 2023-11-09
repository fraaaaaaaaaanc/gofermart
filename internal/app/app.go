package app

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"gofermart/internal/config"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
	"gofermart/internal/router"
	"gofermart/internal/storage"
	"net/http"
)

type app struct {
	flagsConf *config.Flags
	logZap    *logger.ZapLogger
	hndlrs    allhandlers.Handlers
	router    chi.Router
	strgs     *storage.Storage
}

func NewApp() (*app, error) {
	flags := config.ParseConfFlags()
	log, err := logger.NewZapLogger(flags.LogFilePath, flags.ProjLvl)
	if err != nil {
		panic(err)
	}
	strg, err := storage.NewStorage(flags.DataBaseURI)
	if err != nil {
		panic(err)
	}
	hndlr := allhandlers.NewHandlers()
	rtr, err := router.NewRouter(hndlr)
	if err != nil {
		panic(err)
	}

	return &app{
		flagsConf: flags,
		logZap:    log,
		hndlrs:    hndlr,
		router:    rtr,
		strgs:     strg,
	}, nil
}

func Run() error {
	app, _ := NewApp()
	defer app.logZap.CloseFile()

	app.logZap.Info("Server start", zap.String("Running server on", app.flagsConf.String()))
	err := http.ListenAndServe(app.flagsConf.String(), app.router)
	app.logZap.Error("Error run server", zap.Error(err))
	return err
}
