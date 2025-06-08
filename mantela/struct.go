package mantela

import (
	"encoding/json"
	"fmt"
)

type Mantela struct {
	Schema     string             `json:"$schema"`
	Version    string             `json:"version"`
	AboutMe    MantelaAboutMe     `json:"aboutMe"`
	Extensions []MantelaExtension `json:"extensions"`
	Providers  []MantelaProvider  `json:"providers"`
}

type PreferredPrefixType []string

func (p *PreferredPrefixType) UnmarshalJSON(data []byte) error {
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*p = []string{single}
		return nil
	}

	var list []string
	if err := json.Unmarshal(data, &list); err == nil {
		*p = list
		return nil
	}

	return fmt.Errorf("PreferredPrefix: invalid format")
}

type MantelaAboutMe struct {
	Identifier      string              `json:"identifier"`
	Name            string              `json:"name"`
	PreferredPrefix PreferredPrefixType `json:"preferredPrefix"`
	SipUsername     string              `json:"sipUsername"`
	SipPassword     string              `json:"sipPassword"`
	SipServer       string              `json:"sipServer"`
	SipUri          []string            `json:"sipUri"`
	TripUri         []string            `json:"tripUri"`
}

type MantelaExtension struct {
	Name       string   `json:"name"`
	Extension  string   `json:"extension"`
	Type       string   `json:"type"`
	Identifier string   `json:"identifier"`
	TransferTo []string `json:"transferTo,omitempty"`
	Model      string   `json:"model,omitempty"`
}

type MantelaProvider struct {
	Identifier string `json:"identifier"`
	Mantela    string `json:"mantela"`
	Name       string `json:"name"`
	Prefix     string `json:"prefix"`
}
