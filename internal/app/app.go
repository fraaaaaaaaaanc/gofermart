package app

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"gofermart/internal/config"
	"gofermart/internal/cookie"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
	"gofermart/internal/router"
	"gofermart/internal/storage"
	"gofermart/internal/storage/storage_gofermart"
	"gofermart/internal/workwithapi"
	"net/http"
)

type app struct {
	flagsConf *config.Flags
	hndlrs    allhandlers.Handlers
	router    chi.Router
	strgs     storage.StorageGofermart
	workAPI   *workwithapi.WorkAPI
	cookie    cookie.Cookie
}

func NewApp() (*app, error) {
	flags := config.ParseConfFlags()

	err := logger.NewZapLogger(flags.LogFilePath, flags.ProjLvl)
	if err != nil {
		return nil, err
	}

	strg, err := storagegofermart.NewStorage(flags.DataBaseURI)
	if err != nil {
		logger.Error("error creating the storage object", zap.Error(err))
		return nil, err
	}

	cookies := cookie.NewCookie(flags.SecretKeyJWTToken)

	hndlr := allhandlers.NewHandlers(strg, cookies)

	workAPI := workwithapi.NewWorkAPI(strg, flags.AccrualSystemAddress)

	rtr, err := router.NewRouter(hndlr, cookies)
	if err != nil {
		logger.Error("error creating the Router object", zap.Error(err))
		return nil, err
	}

	appObj := &app{
		flagsConf: flags,
		hndlrs:    hndlr,
		router:    rtr,
		strgs:     strg,
		workAPI:   workAPI,
		cookie:    cookies,
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
