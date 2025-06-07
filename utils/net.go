package utils

import (
	"net"
	"strings"
)

func GetOutboundIP() (string, error) {
	ifaces, err := net.Interfaces()
	res := ""

	if err != nil {
		return res, err
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
			res = ip.String()

			// prioritize tailscale interface (#1)
			if strings.Contains(i.Name, "tailscale") {
				break
			}
		}
	}

	return res, nil
}
