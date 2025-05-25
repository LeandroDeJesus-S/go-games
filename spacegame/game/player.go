package game

import (
	"spacegame/assets"
	"spacegame/base"
	u "spacegame/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

// Player represents the player structure
type Player struct {
	*base.Sprite
	Speed            float64
	ShootingCoolDown *Timer
	Game             *Game
}

// Cretes a new Player
func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite
	bounds := sprite.Bounds()

	size := &u.Size{Width: bounds.Dx(), Height: bounds.Dy()}
	initialPosition := &u.Vector{
		X: float64(SCREEN_WIDTH/2 - size.Width/2),
		Y: 550,
	}
	initialDirection := &u.Vector{X: STOPPED, Y: STOPPED}
	shootingTimer := NewTimer(20)

	return &Player{
		Sprite: &base.Sprite{
			Img:       sprite,
			Size:      size,
			Position:  initialPosition,
			Direction: initialDirection,
		},
		Speed: .8,

		ShootingCoolDown: shootingTimer,
		Game:             game,
	}
}

/*
GetDirection sets the player direction and return it. The directions is 1 if the player
is moving to right, -1 if is moving to left or 0 if he is not moving. The player moves
only in the X direction
*/
func (p *Player) GetDirection() int {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA):
		p.Direction.X = LEFT
		return LEFT

	case ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD):
		p.Direction.X = RIGHT
		return RIGHT

	default:
		p.Direction.X = STOPPED
		return STOPPED
	}
}

// Move handles the player's movement mechanic
func (p *Player) Move() {
	direction := p.GetDirection()
	if direction == STOPPED {
		return
	}

	MIN_X, MAX_X := 0.0, float64(SCREEN_WIDTH-p.Size.Width)

	newPos := p.Position.X + float64(direction)*p.Speed*BASE_SPEED
	isIntoScreen := IsBetween(newPos, MIN_X, MAX_X)

	if isIntoScreen {
		p.Position.X = newPos
		return
	}

	if direction == LEFT {
		p.Position.X = 0
		return
	}
	p.Position.X = float64(SCREEN_WIDTH - p.Size.Width)

}

// Shoot handles the shooting mechanic
func (p *Player) Shoot() {
	p.ShootingCoolDown.Update()

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.ShootingCoolDown.isReady() {
		for i := range p.Game.Lasers {
			if p.Game.Lasers[i] != nil {
				continue
			}
			p.Game.Lasers[i] = NewLaser(p.Game)
			break
		}
		p.ShootingCoolDown.resetTimer()
	}
}

func (p *Player) Update() {
	p.Move()
	p.Shoot()
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.Img, opt)
}

func (p *Player) GetRect() *Rect {
	return NewRect(p.Sprite)
}
