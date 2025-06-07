package utils

import (
	"fmt"
	"time"
)

func GetTimezone() string {
	timezone, offset := time.Now().Zone()

	return fmt.Sprintf("%s (UTC%+d)", timezone, offset/3600)
}
