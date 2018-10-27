package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Game []struct {
	ID          int64 `jsonapi:"primary,game_id"`
	Board       int   `jsonapi:"attr,board"`
	WhitePlayer struct {
		Player
	} `jsonapi:"white_player"`
	BlackPlayer struct {
		Player
	} `jsonapi:"black_player"`
	Moves []struct {
		Move
	} `jsonapi:"moves"`
	Winner string `jsonapi:"winner"`
	Loser  string `jsonapi:"loser"`
}

type Move struct {
	ID     int64  `jsonapi:"id"`
	GameID int64  `jsonapi:"game_id"`
	X      string `jsonapi:"x"`
	Y      string `jsonapi:"y"`
}

type Player struct {
	ID       int64  `jsonapi:"id"`
	Name     string `jsonapi:"name"`
	Handicap string `jsonapi:"handicap"`
	Rating   string `jsonapi:"rating"`
}

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("mysql", "gotest:gotest@db/local_gotest?charset=utf8&parseTime=True&loc=Local")
local_gotest
	if err != nil {
		fmt.Printf("failed to connect database because %s", err)
	}

	defer db.Close()

	db.AutoMigrate(&Player{})

	router.HandleFunc("/players", IndexPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", ShowPlayer).Methods("GET")
	router.HandleFunc("/players", CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", DeletePlayer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":80", router))
}

func IndexPlayers(w http.ResponseWriter, r *http.Request) {
	var player []Player
	db.Find(&player)
	json.NewEncoder(w).Encode(&player)
}

func ShowPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player Player
	db.First(&player, params["id"])
	json.NewEncoder(w).Encode(&player)
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player Player
	json.NewDecoder(r.Body).Decode(&player)
	db.Create(&player)
	json.NewEncoder(w).Encode(&player)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player Player
	db.First(&player, params["id"])
	db.Delete(&player)

	var players []Player
	db.Find(&players)
	json.NewEncoder(w).Encode(&players)
}
