package main

import (
	"fmt"
	"hotline/internal/config"
	"hotline/internal/handler"
	"hotline/internal/service"
	"hotline/internal/transport"
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
}
