package main

import (
	"spacegame/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(game.SCREEN_WIDTH, game.SCREEN_HEIGHT)
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		panic(err)
	}
}
