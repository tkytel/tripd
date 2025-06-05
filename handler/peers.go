package handler

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/mantela"
)

var Peers []Peer

func HandlePeers(c *fiber.Ctx) error {
	return c.JSON(Peers)
}

func RetrievePeers() {
	cfg := config.Get()

	res, err := mantela.FetchMantela(cfg.Mantela.Url)
	if err != nil {
		log.Println("Failed to fetch mantela:", err)
	}

	p := make([]Peer, 0, len(res.Providers))

	for _, v := range res.Providers {
		if !strings.Contains(v.Identifier, "XXX") {
			isMeasurable := false
			sipServer, _ := mantela.FetchMantela(v.Mantela)

			if sipServer.AboutMe.SipServer != "" {
				isMeasurable = true
			}

			p = append(p, Peer{
				Measurable: isMeasurable,
				Identifier: v.Identifier,
				Rtt:        nil,
			})
		}
	}

	Peers = p

	log.Println("Updated peers with", len(Peers), "entries")
}
