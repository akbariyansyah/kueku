package cmd

import (
	"context"
	"fmt"
	"kueku/config"
	"kueku/internal/api"
	"kueku/internal/api/cake"
	"kueku/internal/container/founder"
	"kueku/internal/pkg/logo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/labstack/echo/v4"
)

var (
	path = "./config/config.yml"
)

func Run() {
	app := founder.NewApp(path)
	cfg := app.Config()
	db := app.DB()
	router := api.BuildEcho(cfg)
	placeholder := squirrel.Question

	query := founder.NewQueries(placeholder)
	repo := founder.NewRepository(db, query)
	usecases := founder.NewUsecases(repo, query)

	cake.CakeRoute(router, usecases.CakeUsecase())

	log.Println(logo.GetLogo())
	srv := &http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%d", cfg.Port),
		ReadTimeout:  cfg.Server.ReadTimeOut,
		WriteTimeout: cfg.Server.WriteTimeOut}

	go func() {
		if err := router.StartServer(srv); err != nil {
			panic(err)
		}

	}()

	gracefullShutdown(cfg, router)
}

func gracefullShutdown(conf *config.Config, router *echo.Echo) {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("shutdown the server in %v ... ", conf.Server.ShutdownDelay)
	time.Sleep(conf.Server.ShutdownDelay)
	if err := router.Shutdown(ctx); err != nil {
		log.Fatalln("failed shutdown the server : ", err)
	}

	log.Println("good bye !")
}
