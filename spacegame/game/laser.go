package game

import (
	"spacegame/assets"
	b "spacegame/base"
	u "spacegame/utils"

	e "github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	*b.Sprite
	Game  *Game
	Speed float64
}

func NewLaser(game *Game) *Laser {
	sprite := assets.LaserSprite
	bounds := sprite.Bounds()
	size := &u.Size{Width: bounds.Dx(), Height: bounds.Dy()}
	speed := 1.

	position := &u.Vector{
		X: game.Player.Position.X + float64(game.Player.Size.Width/2-size.Width/2),
		Y: game.Player.Position.Y - float64(game.Player.Size.Width/2),
	}

	return &Laser{
		Sprite: &b.Sprite{
			Size:     size,
			Img:      sprite,
			Position: position,
		},
		Game:  game,
		Speed: speed,
	}
}

func (l *Laser) Update() {
	l.Position.Y += -l.Speed * BASE_SPEED
	l.Reset()
}

func (l *Laser) Draw(screen *e.Image) {
	opt := &e.DrawImageOptions{}
	opt.GeoM.Translate(l.Position.X, l.Position.Y)
	screen.DrawImage(l.Sprite.Img, opt)
}

// IsOutOfScreen returns true if the laser is off screen
func (l *Laser) IsOutOfScreen() bool {
	return l.Position.Y < 0-float64(l.Size.Height)
}

// Reset removes all the lasers out of the screen when the array of lasers is full filled
func (l *Laser) Reset() {
	if l.Game.Lasers[len(l.Game.Lasers)-1] == nil {
		return
	}

	for i, laser := range l.Game.Lasers {
		if laser != nil && laser.IsOutOfScreen() {
			l.Game.Lasers[i] = nil
		}
	}
}


func (l *Laser) GetRect() *Rect {
	return NewRect(l.Sprite)
}