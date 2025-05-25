package base

import (
	u "spacegame/utils"

	e "github.com/hajimehoshi/ebiten/v2"
)


type Sprite struct {
	Img *e.Image
	Position *u.Vector
	Direction *u.Vector
	Size *u.Size
}

type ISprite interface {
	Draw(screen *e.Image)
	Update()
}
