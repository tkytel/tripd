package handler

type Peer struct {
	Measurable bool     `json:"measurable"`
	Identifier string   `json:"identifier"`
	Rtt        *float32 `json:"rtt"`
}

type About struct {
	OutboundAddress string `json:"outbound_address"`
	Timezone        string `json:"timezone"`
}
