package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tkytel/tripd/utils"
)

func HandleAbout(c *fiber.Ctx) error {
	outboundAddress, err := utils.GetOutboundIP()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Could not determine outbound address: %v", err),
		})
	}

	res := About{
		OutboundAddress: outboundAddress,
		Timezone:        utils.GetTimezone(),
	}

	return c.JSON(res)
}
