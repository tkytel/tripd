package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var c Config

func Init() {
	d := read()
	c = unmarshal(d)
	log.Println("Successfully loaded configuration")
}

func Get() Config {
	return c
}

func read() []byte {
	bytes, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatalln(err)
	}

	return bytes
}

func unmarshal(d []byte) Config {
	var c Config

	err := toml.Unmarshal(d, &c)
	if err != nil {
		log.Fatalln(err)
	}

	return c
}
