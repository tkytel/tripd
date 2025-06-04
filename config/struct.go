package config

type Config struct {
	Mantela Mantela `toml:"mantela"`
}

type Mantela struct {
	Url string `toml:"url"`
}
