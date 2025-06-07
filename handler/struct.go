package handler

type About struct {
	Identifier      string `json:"identifier"`
	OutboundAddress string `json:"outbound_address"`
	Timezone        string `json:"timezone"`
	HopEnabled      bool   `json:"hop_enabled"`
}
