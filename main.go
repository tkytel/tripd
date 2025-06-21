package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/handler"
	"github.com/tkytel/tripd/utils"
	"github.com/tkytel/tripd/web"
)

func main() {
	rand.Seed(time.Now().UnixNano())
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
