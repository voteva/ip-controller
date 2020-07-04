package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/voteva/ip-controller/internal/app/ip-controller/config"
	"github.com/voteva/ip-controller/internal/app/ip-controller/server"
	"github.com/voteva/ip-controller/internal/app/ip-controller/store/db"
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

	srv := server.New(conf.BindAddr)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
