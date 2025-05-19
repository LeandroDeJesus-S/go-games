package game

import (
	"spacegame/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image              *ebiten.Image
	position           *Vector
	game               *Game
	laserShootingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite
	bounds := image.Bounds()

	halfImgX := float64(bounds.Dx()) / 2
	x := SCREEN_WIDTH / 2 - halfImgX
	y := SCREEN_HEIGHT * .75

	position := &Vector{x, y}
	return &Player{
		image: image,
		position: position,
		game: game,
		laserShootingTimer: NewTimer(15),
	}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.x -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.x += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserShootingTimer.isReady() {
		bounds := p.image.Bounds()

		halfX := float64(bounds.Dx()) / 2
		halfY := float64(bounds.Dy()) / 2

		laserPos := Vector{
			p.position.x + halfX,
			p.position.y - halfY,
		}
		laser := NewLaser(laserPos)
		p.game.addLaser(laser)

		p.laserShootingTimer.resetTimer()
	}

	p.laserShootingTimer.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}

	opt.GeoM.Translate(p.position.x, p.position.y)
	screen.DrawImage(p.image, opt)
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()
	return NewRect(*p.position, bounds.Dx(), bounds.Dy())
}