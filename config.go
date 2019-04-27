package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

var config Configuration

type Configuration struct {
	Mode    string
	General General
}
type General struct {
	ServerAddress string
}

func LoadConfiguration() {
	path := "./config.toml"

	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Printf("Couldn't read config file at [%s]\n", path)
		log.Fatal(err)
	}
}
