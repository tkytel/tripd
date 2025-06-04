package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tkytel/tripd/mantela"
)

func HandlePeers(c *fiber.Ctx) error {
	res, err := mantela.FetchMantela()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch peers",
		})
	}

	peers := make([]Peer, 0, len(res.Providers))

	for _, v := range res.Providers {
		if !strings.Contains(v.Identifier, "XXX") {
			peers = append(peers, Peer{
				Identifier: v.Identifier,
				Rtt:        nil,
			})
		}
	}

	return c.JSON(peers)
}
