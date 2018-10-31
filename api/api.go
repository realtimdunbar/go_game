package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jasonmccallister/go_game/models"
	"github.com/jinzhu/gorm"
)

type server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// New will create a new server struct for the API configuration
func New(dialect, conn string) (server, error) {
	db, err := gorm.Open(dialect, conn)
	if err != nil {
		return s, err
	}

	s := server{
		Router: mux.NewRouter(),
	}

	s.DB = db

	return s, nil
}

func (s *server) routes() {
	s.Router.HandleFunc("/players", s.IndexPlayers).Methods("GET")
	s.Router.HandleFunc("/players/{id}", s.ShowPlayer).Methods("GET")
	s.Router.HandleFunc("/players", s.CreatePlayer).Methods("POST")
	s.Router.HandleFunc("/players/{id}", s.DeletePlayer).Methods("DELETE")
}

func (s *server) IndexPlayers(w http.ResponseWriter, r *http.Request) {
	var player []models.Player
	s.DB.Find(&player)
	json.NewEncoder(w).Encode(&player)
}

func (s *server) ShowPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player models.Player
	s.DB.First(&player, params["id"])
	json.NewEncoder(w).Encode(&player)
}

func (s *server) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	json.NewDecoder(r.Body).Decode(&player)
	s.DB.Create(&player)
	json.NewEncoder(w).Encode(&player)
}

func (s *server) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player models.Player
	s.DB.First(&player, params["id"])
	s.DB.Delete(&player)

	var players []models.Player
	s.DB.Find(&players)
	json.NewEncoder(w).Encode(&players)
}
