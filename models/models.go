package models

import (
	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Board       int     `json:"attr,board"`
	WhitePlayer Player  `gorm:"foreignkey:ID;association_foreignkey:Refer" json:"white_player"`
	BlackPlayer Player  `gorm:"foreignkey:ID;association_foreignkey:Refer" json:"black_player"`
	Stones      []Stone `json:"stones"`
	Winner      Player  `gorm:"foreignkey:ID;association_foreignkey:Refer" json:"winner"`
	Loser       Player  `gorm:"foreignkey:ID;association_foreignkey:Refer" json:"loser"`
}

type Stone struct {
	gorm.Model
	GameID        int64  `gorm:"foreignkey:ID;association_foreignkey:Refer" json:"game_id"`
	X             string `json:"x"`
	Y             string `json:"y"`
	LibertyTop    *Stone `json:"liberty_top"`
	LibertyBottom *Stone `json:"liberty_bottom"`
	LibertyLeft   *Stone `json:"liberty_left"`
	LibertyRight  *Stone `json:"liberty_right"`
	Color         string `json:"color"`
}

type Player struct {
	gorm.Model
	Name     string `json:"name"`
	Handicap string `json:"handicap"`
	Rating   string `json:"rating"`
}
