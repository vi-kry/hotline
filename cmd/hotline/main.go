package main

import (
	"context"
	"fmt"
	"hotline/internal/config"
	"hotline/internal/handler"
	"hotline/internal/models"
	"hotline/internal/server"
	"hotline/internal/service"
	"hotline/internal/transport"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	transports := transport.NewTransport(cfg)
	fmt.Println(transports)

	services := service.NewService(transports)
	fmt.Println(services)

	handlers := handler.NewHandler(services)
	fmt.Println(handlers)

	srv := new(server.Server)

	go func() {
		cfgHTTP := models.HTTPServer{
			Port: cfg.HTTPServer.Port,
			Path: cfg.HTTPServer.Path,
		}
		log.Printf("start http server on port %s", cfgHTTP.Port)
		err := srv.Run(cfgHTTP.Port, handlers.InitRoutes(cfgHTTP.Path, cfg.Env))
		if err != http.ErrServerClosed {
			log.Panicf("error occurred while running http server: %s", err.Error())
		}
	}()

	signalListner := make(chan os.Signal, 1)
	signal.Notify(signalListner,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	stop := <-signalListner
	fmt.Println(cfg)
	log.Printf("Shutting Down app: %s", stop)

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occurred on server shutting down: %s", err.Error())
	}
}
