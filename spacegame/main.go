package main

import (
	"spacegame/game"
	
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		panic(err)
	}
}
