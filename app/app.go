package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/DeniesKresna/danatest/config"
	"github.com/DeniesKresna/danatest/service/delivery/gate"
	"github.com/DeniesKresna/danatest/service/repository/mysql"
	"github.com/DeniesKresna/danatest/service/usecase"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Application struct {
	Conf    *config.Config
	AppGate *gate.Gate
}

func InitApp(conf *config.Config) *Application {
	mysqlStockRepo := mysql.InitMysqlStockRepository(conf.DB)
	stockUsecase := usecase.InitStockUsecase(conf.DB, mysqlStockRepo)
	appGate := gate.InitGate(stockUsecase)

	return &Application{
		Conf:    conf,
		AppGate: appGate,
	}
}

func (app *Application) IsListPathExisted(path string, listPath []string) (res bool) {
	for _, p := range listPath {
		fp := strings.ToLower(p)
		if fp == path {
			utlog.Errorf("path: %s has been registered", path)
			return true
		}
	}
	return false
}

func (app *Application) GateOpen() (err error) {
	r := mux.NewRouter().StrictSlash(true)

	listRoutesPath := []string{}

	if len(app.AppGate.ListRoutes) <= 0 {
		err = errors.New("no routes found")
		return
	}

	for _, route := range app.AppGate.ListRoutes {
		if app.IsListPathExisted(route.Path, listRoutesPath) {
			err = errors.New("route has been registered")
			return
		}

		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)

		listRoutesPath = append(listRoutesPath, route.Path)
	}

	//this changes for cors for browser
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"Get", "Post", "Delete", "Put", http.MethodOptions},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	router := c.Handler(r)

	addr := fmt.Sprintf("%s:%s", app.Conf.Service.Host, app.Conf.Service.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utlog.Error("open service gate to the world")
		}
	}()

	utlog.Infof("service started on %s", addr)

	<-done
	utlog.Info("closing the service gate...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		utlog.Info("gate closed failed")
	}
	utlog.Info("gate closed properly")

	return nil
}
