package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/realtimdunbar/go_game/api"
)

var (
	flagPort    *string
	flagMigrate *bool
)

func main() {
	flag.StringVar(flagPort, "port", "80", "which port the server should listen on")
	flag.BoolVar(flagMigrate, "migrate", true, "if the application should run database migrations")
	flag.Parse()

	server, err := api.New("mysql", "gotest:gotest@tcp(db:3306)/local_gotest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer server.DB.Close()
	if flagMigrate {
		api.MigrateDB()
	}

	log.Fatal(http.ListenAndServe(":8080", server.Router))
}
