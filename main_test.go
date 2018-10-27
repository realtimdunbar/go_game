package main

import (
	"testing"
)

func TestCreateBoard(t *testing.T) {
	s := 9
	p1 := "Tim"
	p2 := "Bob"
	ws := p1
	bs := p2

	g := CreateGame(s, p1, p2, ws, bs)

	for _, h := range g.Board {
		if len(h) != s {
			t.Errorf("expected %d", s)
		}
	}
	if len(g.Board) != s {
		t.Errorf("expected %d", s)
	}

	if g.PlayerOne != "Tim" {
		t.Errorf("expected %s", p1)
	}

	if g.PlayerTwo != "Bob" {
		t.Errorf("expected %s", p2)
	}

	if g.WhiteStones != "Tim" {
		t.Errorf("expected %s", ws)
	}

	if g.BlackStones != "Bob" {
		t.Errorf("expected %s", bs)
	}
}

func TestPlaceStone(t *testing.T) {

}
