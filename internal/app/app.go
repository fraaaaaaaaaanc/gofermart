package app

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"gofermart/internal/config"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
	"gofermart/internal/router"
	"gofermart/internal/storage"
	"gofermart/internal/storage/storage_db"
	"gofermart/internal/workwithapi"
	"net/http"
)

type app struct {
	flagsConf *config.Flags
	hndlrs    allhandlers.Handlers
	router    chi.Router
	strgs     storage.StorageMock
	workAPI   *workwithapi.WorkAPI
}

func NewApp() (*app, error) {
	flags := config.ParseConfFlags()
	err := logger.NewZapLogger(flags.LogFilePath, flags.ProjLvl)
	if err != nil {
		panic(err)
	}
	strg, err := storage_db.NewStorage(flags.DataBaseURI)
	if err != nil {
		panic(err)
	}
	hndlr := allhandlers.NewHandlers(strg, flags.SecretKeyJWTToken)
	workAPI := workwithapi.NewWorkAPI(strg, flags.AccrualSystemAddress)
	rtr, err := router.NewRouter(hndlr, flags.SecretKeyJWTToken)
	if err != nil {
		panic(err)
	}

	appObj := &app{
		flagsConf: flags,
		hndlrs:    hndlr,
		router:    rtr,
		strgs:     strg,
		workAPI:   workAPI,
	}

	return appObj, nil
}

func Run() error {
	app, _ := NewApp()
	defer logger.CloseFile()
	defer app.strgs.CloseDB()
	logger.Info("Server start", zap.String("Running server on", app.flagsConf.String()))
	err := http.ListenAndServe(app.flagsConf.String(), app.router)
	logger.Error("Error run server", zap.Error(err))
	return err
}
