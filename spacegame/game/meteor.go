package game

import (
	"math/rand"
	"spacegame/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	image *ebiten.Image
	speed float64
	position Vector
}

func NewMeteor() *Meteor {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]
	speed := rand.Float64() * 8

	position := Vector{
		x: rand.Float64() * SCREEN_WIDTH,
		y: -100,
	}

	return &Meteor{
		image: image,
		speed: speed,
		position: position,
	}
}

func (m *Meteor) Update(){
	m.position.y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(m.position.x, m.position.y)

	screen.DrawImage(m.image, opt)
}

func (m *Meteor) Collider() Rect {
	bounds := m.image.Bounds()
	return NewRect(m.position, bounds.Dx(), bounds.Dy())
}