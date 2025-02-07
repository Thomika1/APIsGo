package main

import (
	"github.com/Thomika1/APIsGo.git/config"
	"github.com/Thomika1/APIsGo.git/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("->main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config init error: %v", err)
		return
	}

	router.Initialize()

} // Main
