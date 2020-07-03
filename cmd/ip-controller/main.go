package main

import (
	"flag"
	"github.com/voteva/ip-controller/internal/app/server"
	"github.com/voteva/ip-controller/internal/app/store/db"
	"github.com/voteva/ip-controller/internal/config"
	"github.com/BurntSushi/toml"
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
