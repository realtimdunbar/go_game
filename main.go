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
	Color  string `jsonapi:"color"`
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
	router.HandleFunc("/games", IndexGames).Methods("GET")
	router.HandleFunc("/games/{id}", ShowGame).Methods("GET")
	router.HandleFunc("/games", CreateGame).Methods("POST")
	router.HandleFunc("/games/{id}", DeleteGame).Methods("DELETE")
	router.HandleFunc("/moves", IndexMoves).Methods("GET")
	router.HandleFunc("/moves/{id}", ShowMove).Methods("GET")
	router.HandleFunc("/moves", CreateMove).Methods("POST")
	router.HandleFunc("/moves/{id}", DeleteMove).Methods("DELETE")

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

func IndexGames(w http.ResponseWriter, r *http.Request) {
	var games []Game
	db.Find(&games)
	json.NewEncoder(w).Encode(&games)
}

func ShowGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var game Game
	db.First(&game, params["id"])
	json.NewEncoder(w).Encode(&game)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var game Game
	json.NewDecoder(r.Body).Decode(&game)
	db.Create(&game)
	json.NewEncoder(w).Encode(&game)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var game Game
	db.First(&game, params["id"])
	db.Delete(&game)

	var games []Game
	db.Find(&games)
	json.NewEncoder(w).Encode(&games)
}

func IndexMoves(w http.ResponseWriter, r *http.Request) {
	var moves []Move
	db.Find(&moves)
	json.NewEncoder(w).Encode(&moves)
}

func ShowMove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var move Move
	db.First(&move, params["id"])
	json.NewEncoder(w).Encode(&move)
}

func CreateMove(w http.ResponseWriter, r *http.Request) {
	var move Move
	json.NewDecoder(r.Body).Decode(&move)
	db.Create(&move)
	json.NewEncoder(w).Encode(&move)
}

func DeleteMove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var move Move
	db.First(&move, params["id"])
	db.Delete(&move)

	var moves []Move
	db.Find(&moves)
	json.NewEncoder(w).Encode(&moves)
}
