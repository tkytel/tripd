package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkytel/tripd/utils"
)

func HandlePeers(c *fiber.Ctx) error {
	if Ready {
		return c.JSON(utils.Peers)
	} else {
		return c.SendStatus(503)
	}
}
