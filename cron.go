package main

import (
	"time"

	"github.com/tkytel/tripd/utils"
)

func startBackgroundTasks() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			utils.RetrievePeers()
		}
	}
}
