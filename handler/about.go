package handler

import (
	"net"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HandleAbout(c *fiber.Ctx) error {
	res := About{
		OutboundAddress: "",
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return c.JSON(fiber.Map{
			"error": "Failed to fetch network interfaces",
		})
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.To4() == nil {
				// IPv6 or invalid address. skipping
				continue
			}
			res.OutboundAddress = ip.String()

			// prioritize tailscale interface (#1)
			if strings.Contains(i.Name, "tailscale") {
				break
			}
		}
	}

	if res.OutboundAddress != "" {
		return c.JSON(res)
	}

	return c.JSON(fiber.Map{
		"error": "Could not find effective outbound address",
	})
}
