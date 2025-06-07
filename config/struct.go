package config

type Config struct {
	Mantela Mantela `toml:"mantela"`
	Ping    Ping    `toml:"ping"`
}

type Mantela struct {
	Url string `toml:"url"`
}

type Ping struct {
	Count int `toml:"count"`
}
