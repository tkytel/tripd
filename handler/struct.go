package handler

import "time"

type About struct {
	Identifier      string    `json:"identifier" desc:"identifier of this PBX system on Mantela"`
	OutboundAddress string    `json:"outbound_address" desc:"which address will be used in order to connect other PBXs"`
	Timezone        string    `json:"timezone" desc:"timezone of this PBX system"`
	HopEnabled      bool      `json:"hop_enabled" desc:"whether if this PBX system supports hopping"`
	LastMeasured    time.Time `json:"last_measured" desc:"timestamp of last measure happened"`
}
