package app

import (
	"fmt"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"gofermart/internal/config"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
	"gofermart/internal/router"
	"gofermart/internal/storage"
	"gofermart/internal/workwithapi"
	"net/http"
)

type app struct {
	flagsConf *config.Flags
	logZap    *logger.ZapLogger
	hndlrs    allhandlers.Handlers
	router    chi.Router
	strgs     *storage.Storage
	workAPI   *workwithapi.WorkAPI
}

func NewApp() (*app, error) {
	flags := config.ParseConfFlags()
	log, err := logger.NewZapLogger(flags.LogFilePath, flags.ProjLvl)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	strg, err := storage.NewStorage(flags.DataBaseURI, log.Log)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	hndlr := allhandlers.NewHandlers(log.Log, strg, flags.SecretKeyJWTToken)
	workAPI := workwithapi.NewWorkAPI(log.Log, strg)
	rtr, err := router.NewRouter(hndlr, log.Log, flags.SecretKeyJWTToken)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	appObj := &app{
		flagsConf: flags,
		logZap:    log,
		hndlrs:    hndlr,
		router:    rtr,
		strgs:     strg,
		workAPI:   workAPI,
	}

	return appObj, nil
}

func Run() error {
	app, _ := NewApp()
	defer app.logZap.CloseFile()
	app.logZap.Log.Info("Server start", zap.String("Running server on", app.flagsConf.String()))
	err := http.ListenAndServe(app.flagsConf.String(), app.router)
	app.logZap.Log.Error("Error run server", zap.Error(err))
	return err
}
