package api_test

import (
	"net/http"
	"testing"

	"github.com/realtimdunbar/go_game/api"
)

func TestIndexPlayers(t *testing.T) {
	// Arrange
	s, err := api.New("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	s.Routes()

	// Act
	req, err := http.NewRequest(http.MethodGet, "players", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if req.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %v instead", req.Response.StatusCode)
	}
}
