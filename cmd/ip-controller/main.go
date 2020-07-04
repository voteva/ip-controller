package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/voteva/ip-controller/internal/app/config"
	"github.com/voteva/ip-controller/internal/app/api"
	"github.com/voteva/ip-controller/internal/app/store/db"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/local.toml", "")
}

func main() {
	flag.Parse()

	conf := config.NewConfig()
	if _, err := toml.DecodeFile(configPath, conf); err != nil {
		log.Fatal(err)
	}

	if err := db.Connect(conf.DatabaseURL); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	srv := api.New(conf.BindAddr)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
