package main

import (
	"log"
	"net/http"

	"github.com/realtimdunbar/go_game/api"
	"github.com/realtimdunbar/go_game/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	server, err := api.New("mysql", "gotest:gotest@tcp(db:3306)/local_gotest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	defer server.DB.Close()

	server.DB.AutoMigrate(&models.Player{})
	server.DB.AutoMigrate(&models.Game{})
	server.DB.AutoMigrate(&models.Stone{})

	log.Fatal(http.ListenAndServe(":8080", server.Router))
}
