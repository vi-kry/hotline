package main

import (
	"fmt"
	"hotline/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
