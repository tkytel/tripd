package main

import (
	"log"

	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/handler"
	"github.com/tkytel/tripd/utils"
	"github.com/tkytel/tripd/web"
)

func main() {
	handler.Ready = false
	config.Init()

	log.Println("Initializing peers directory")
	go func() {
		utils.RetrievePeers()
		handler.Ready = true
	}()

	go startBackgroundTasks()

	web.Init()
}
