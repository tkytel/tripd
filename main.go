package main

import (
	"log"

	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/utils"
	"github.com/tkytel/tripd/web"
)

func main() {
	config.Init()

	log.Println("Initializing peers directory")
	utils.RetrievePeers()

	go startBackgroundTasks()

	web.Init()
}
