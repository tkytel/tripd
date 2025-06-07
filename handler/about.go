package handler

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/tkytel/tripd/config"
	"github.com/tkytel/tripd/mantela"
	"github.com/tkytel/tripd/utils"
)

func HandleAbout(c *fiber.Ctx) error {
	res, err := GenerateAbout()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(res)
}

func GenerateAbout() (*About, error) {
	outboundAddress, err := utils.GetOutboundIP()
	if err != nil {
		return nil, fmt.Errorf("could not determine outbound address: %v", err)
	}

	cfg := config.Get()

	identifier := ""
	hopEnabled := false
	var wg sync.WaitGroup

	res, err := mantela.FetchMantela(cfg.Mantela.Url)
	if err != nil {
		goto End
	}

	identifier = res.AboutMe.Identifier

	for _, v := range res.Providers {
		v := v
		wg.Add(1)

		go func() {
			defer wg.Done()

			if v.Identifier == res.AboutMe.Identifier {
				hopEnabled = true
				return
			}
		}()

		wg.Wait()
	}

End:
	resp := About{
		Identifier:      identifier,
		OutboundAddress: outboundAddress,
		Timezone:        utils.GetTimezone(),
		HopEnabled:      hopEnabled,
		LastMeasured:    utils.LastMeasured,
	}

	return &resp, nil
}
