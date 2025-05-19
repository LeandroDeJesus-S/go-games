package game

import (
	"spacegame/assets"

	"github.com/hajimehoshi/ebiten/v2"
)


type Laser struct {
	image *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite
	bouds := image.Bounds()

	halfX := bouds.Dx() / 2
	halfY := bouds.Dy() / 2

	position.x -= float64(halfX)
	position.y -= float64(halfY)

	return &Laser{image, position}
}

func (l *Laser) Update(){
	speed := 7.0

	l.position.y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(l.position.x, l.position.y)
	screen.DrawImage(l.image, opt)
}

func (l *Laser) Collider() Rect {
	bounds := l.image.Bounds()
	return NewRect(l.position, bounds.Dx(), bounds.Dy())
}
