package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/mantela"
)

func HandlePeers(c *fiber.Ctx) error {
	cfg := config.Get()

	res, err := mantela.FetchMantela(cfg.Mantela.Url)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch peers",
		})
	}

	peers := make([]Peer, 0, len(res.Providers))

	for _, v := range res.Providers {
		if !strings.Contains(v.Identifier, "XXX") {
			isMeasurable := false
			sipServer, _ := mantela.FetchMantela(v.Mantela)

			if sipServer.AboutMe.SipServer != "" {
				isMeasurable = true
			}

			peers = append(peers, Peer{
				Measurable: isMeasurable,
				Identifier: v.Identifier,
				Rtt:        nil,
			})
		}
	}

	return c.JSON(peers)
}
