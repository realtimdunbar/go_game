package api_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/realtimdunbar/go_game/api"
)

func mockAPI() api.Server {
	s, err := api.New("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	s.Routes()

	return s
}

func TestIndexPlayers(t *testing.T) {
	// Arrange
	s := mockAPI()

	// Act
	req, err := http.NewRequest(http.MethodGet, "/players", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.IndexPlayers)
	handler.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestIndexGames(t *testing.T) {
	// Arrange
	s := mockAPI()

	// Act
	req, err := http.NewRequest(http.MethodGet, "/games", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.IndexGames)
	handler.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestIndexStones(t *testing.T) {
	// Arrange
	s := mockAPI()

	// Act
	req, err := http.NewRequest(http.MethodGet, "/stones", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.IndexStones)
	handler.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
