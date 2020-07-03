package main

import (
	"flag"
	"git.ozon.dev/tvoteva/22_ide/homework/internal/app/server"
	"git.ozon.dev/tvoteva/22_ide/homework/internal/app/store/db"
	"git.ozon.dev/tvoteva/22_ide/homework/internal/config"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
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
