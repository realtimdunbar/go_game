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

type Game struct {
	gorm.Model
	GameID      int64  `jsonapi:"primary,game_id"`
	Board       int    `jsonapi:"attr,board"`
	WhitePlayer Player `gorm:"foreignkey:PlayerID;association_foreignkey:Refer" jsonapi:"white_player"`
	BlackPlayer Player `gorm:"foreignkey:PlayerID;association_foreignkey:Refer" jsonapi:"black_player"`
	Moves       []Move `jsonapi:"moves"`
	Winner      string `jsonapi:"winner"`
	Loser       string `jsonapi:"loser"`
}

type Move struct {
	gorm.Model
	MoveID int64  `jsonapi:"primary,move_id"`
	GameID int64  `gorm:"foreignkey:GameID;association_foreignkey:Refer" jsonapi:"game_id"`
	X      string `jsonapi:"x"`
	Y      string `jsonapi:"y"`
}

type Player struct {
	gorm.Model
	PlayerID int64  `jsonapi:"primary,player_id"`
	Name     string `jsonapi:"name"`
	Handicap string `jsonapi:"handicap"`
	Rating   string `jsonapi:"rating"`
}

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("mysql", "gotest:gotest@tcp(db:3306)/local_gotest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Printf("failed to connect database because %s", err)
	}

	defer db.Close()

	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Game{})
	db.AutoMigrate(&Move{})

	router.HandleFunc("/players", IndexPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", ShowPlayer).Methods("GET")
	router.HandleFunc("/players", CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", DeletePlayer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func IndexPlayers(w http.ResponseWriter, r *http.Request) {
	var players []Player
	db.Find(&players)
	json.NewEncoder(w).Encode(&players)
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
