package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/S5477/go-simple-server/internal/app/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to cinfig file")
}

func main() {
	flag.Parse()
	config := api.NewConfig()

	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Println("can not find configs file. using default values: ", err)
	}

	server := api.New(config)

	log.Fatal(server.Run())
}
