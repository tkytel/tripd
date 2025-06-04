package handler

type Peer struct {
	Identifier string   `json:"identifier"`
	Rtt        *float32 `json:"rtt"`
}

type About struct {
	OutboundAddress string `json:"outbound_address"`
}
