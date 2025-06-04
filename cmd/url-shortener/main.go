package main

import (
	"fmt"

	"github.com/Lacky1234union/UrlShorter/internal/config"
)

func main() {
	cfg := config.MustLoad() // TODO: init config: cleanenv

	fmt.Println(cfg)
	//
	//TODO: init logger: slog (logger)
	//
	//TODO: init roouter: chi, "chi render"
	//
	//TODO: run server
}
