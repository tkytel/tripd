package main

import (
	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/web"
)

func main() {
	config.Init()
	web.Init()
}
