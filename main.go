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
