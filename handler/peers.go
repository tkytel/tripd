package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkytel/tripd/utils"
)

func HandlePeers(c *fiber.Ctx) error {
	return c.JSON(utils.Peers)
}
