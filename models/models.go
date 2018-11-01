package models

import (
	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Board       int     `jsonapi:"attr,board"`
	WhitePlayer Player  `gorm:"foreignkey:ID;association_foreignkey:Refer" jsonapi:"white_player"`
	BlackPlayer Player  `gorm:"foreignkey:ID;association_foreignkey:Refer" jsonapi:"black_player"`
	Stones      []Stone `jsonapi: "stones"`
	Winner      Player  `gorm:"foreignkey:ID;association_foreignkey:Refer" jsonapi:"winner"`
	Loser       Player  `gorm:"foreignkey:ID;association_foreignkey:Refer" jsonapi:"loser"`
}

type Stone struct {
	gorm.Model
	GameID        int64  `gorm:"foreignkey:ID;association_foreignkey:Refer" jsonapi:"game_id"`
	X             string `jsonapi:"x"`
	Y             string `jsonapi:"y"`
	LibertyTop    *Stone `jsonapi:"liberty_top"`
	LibertyBottom *Stone `jsonapi:"liberty_bottom"`
	LibertyLeft   *Stone `jsonapi:"liberty_left"`
	LibertyRight  *Stone `jsonapi:"liberty_right"`
}

type Player struct {
	gorm.Model
	Name     string `jsonapi:"name"`
	Handicap string `jsonapi:"handicap"`
	Rating   string `jsonapi:"rating"`
}
