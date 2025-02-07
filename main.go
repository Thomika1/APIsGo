package main

import (
	"fmt"

	"github.com/Thomika1/APIsGo.git/config"
	"github.com/Thomika1/APIsGo.git/router"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()

} // Main
