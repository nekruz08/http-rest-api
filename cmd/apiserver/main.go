package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"github.com/nekruz08/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}
func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}


	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}

//
//m, err := migrate.New(
//"github://mattes:personal-access-token@mattes/migrate_test",
//"postgres://postgres:postgres@localhost:5432/restapi_dev?sslmode=disable")
//m.Steps(2)

