package game

import (
	"math/rand"
	"spacegame/assets"
	"spacegame/base"
	u "spacegame/utils"

	e "github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	*base.Sprite
	Speed int
	Game *Game
}

// Initializes the Meteor structure
func NewMeteor(game *Game) *Meteor {
	spriteIndex := rand.Intn(len(assets.MeteorSprites))
	image := assets.MeteorSprites[spriteIndex]
	bounds := image.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	speed := 1
	pos := &u.Vector{
		X: rand.Float64() * SCREEN_WIDTH,
		Y: -100,
	}

	return &Meteor{
		Sprite: &base.Sprite{
			Img:      image,
			Position: pos,
			Size: &u.Size{Width: w, Height: h},
			Direction: &u.Vector{X: STOPPED, Y: DOWN},
		},
		Speed: speed,
		Game:  game,
	}
}

func (m *Meteor) Update() {
	m.Position.Y += float64(m.Speed) * BASE_SPEED
	m.Reset()
}

func (m *Meteor) Draw(screen *e.Image) {
	opt := &e.DrawImageOptions{}
	opt.GeoM.Translate(m.Position.X, m.Position.Y)

	screen.DrawImage(m.Sprite.Img, opt)
}

// IsOutOfScreen returns true if the meteor is off screen
func (m *Meteor) IsOutOfScreen() bool {
	return m.Position.Y > SCREEN_HEIGHT
}

// Reset removes all the meteors out of the screen when the array of meteors is full filled
func (m *Meteor) Reset() {
	if m.Game.Meteors[len(m.Game.Meteors)-1] == nil {
		return
	}

	for i, laser := range m.Game.Meteors {
		if laser != nil && laser.IsOutOfScreen() {
			m.Game.Meteors[i] = nil
		}
	}
}

func (m *Meteor) GetRect() *Rect {
	return NewRect(m.Sprite)
}

func SpawnMeteor(m *Game) {
	m.MeteorsSpawnTimer.Update()
	if !m.MeteorsSpawnTimer.isReady() {
		return
	}

	m.MeteorsSpawnTimer.resetTimer()
	for i := range m.Meteors {
		if m.Meteors[i] == nil {
			m.Meteors[i] = NewMeteor(m)
			return
		}
	}
}
