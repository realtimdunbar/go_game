package models

type Game []struct {
	ID          int64 `jsonapi:"primary,game_id"`
	Board       int   `jsonapi:"attr,board"`
	WhitePlayer struct {
		Player
	} `jsonapi:"white_player"`
	BlackPlayer struct {
		Player
	} `jsonapi:"black_player"`
	Stones []struct {
		Stone
	} `jsonapi: Stones"`
	Winner string `jsonapi:"winner"`
	Loser  string `jsonapi:"loser"`
}

type Stone struct {
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
