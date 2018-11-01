package main

import (
	"log"
	"net/http"

	"github.com/realtimdunbar/go_game/api"
	"github.com/realtimdunbar/go_game/models"
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
