package app

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"gofermart/internal/config"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
	"gofermart/internal/models/handlersmodels"
	"gofermart/internal/router"
	"gofermart/internal/storage"
	"gofermart/internal/workwithapi"
	"net/http"
	"time"
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
		panic(err)
	}
	strg, err := storage.NewStorage(flags.DataBaseURI)
	if err != nil {
		panic(err)
	}
	hndlr := allhandlers.NewHandlers(log.Log, strg)
	workAPI := workwithapi.NewWorkAPI(log.Log, strg)
	rtr, err := router.NewRouter(hndlr, log.Log)
	if err != nil {
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

	go appObj.listenChanel()
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

func (a *app) listenChanel() {
	ticker := time.NewTicker(1 * time.Second)

	var orderInfoList []*handlersmodels.OrderInfo
	for {
		select {
		case orderInfo := <-a.hndlrs.Ch:
			orderInfoList = append(orderInfoList, orderInfo)
		case <-ticker.C:
			if len(orderInfoList) == 0 {
				continue
			}
			for _, orderInfo := range orderInfoList {
				a.workAPI.RegisterOrderNumber(orderInfo)
			}
			orderInfoList = nil
		}
	}
}
