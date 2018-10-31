package main

import (
	"log"
	"net/http"

	"github.com/jasonmccallister/go_game/api"
	"github.com/jasonmccallister/go_game/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	server, err := api.New("mysql", "gotest:gotest@db/local_gotest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	defer server.DB.Close()

	server.DB.AutoMigrate(&models.Player{})

	log.Fatal(http.ListenAndServe(":80", server.Router))
}
