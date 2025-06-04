package mantela

type Mantela struct {
	Schema     string             `json:"$schema"`
	Version    string             `json:"version"`
	AboutMe    MantelaAboutMe     `json:"aboutMe"`
	Extensions []MantelaExtension `json:"extensions"`
	Providers  []MantelaProvider  `json:"providers"`
}

type MantelaAboutMe struct {
	Identifier      string   `json:"identifier"`
	Name            string   `json:"name"`
	PreferredPrefix []string `json:"preferredPrefix"`
	SipUsername     string   `json:"sipUsername"`
	SipPassword     string   `json:"sipPassword"`
	SipServer       string   `json:"sipServer"`
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
