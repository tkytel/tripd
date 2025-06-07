package utils

type Peer struct {
	Measurable bool     `json:"measurable"`
	Identifier string   `json:"identifier"`
	Rtt        *float32 `json:"rtt"`
}
