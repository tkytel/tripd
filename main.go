package main

import (
	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/handler"
	"github.com/tkytel/tripd/web"
)

func main() {
	config.Init()

	handler.RetrievePeers()
	go startBackgroundTasks()

	web.Init()
}
