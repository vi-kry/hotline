package main

import (
	"fmt"
	"hotline/internal/config"
	"hotline/internal/transport"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	transports := transport.NewTransport(cfg)
	fmt.Println(transports)
}
