package game

import (
	"fmt"
	"image/color"
	"spacegame/assets"

	e "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	Player            *Player
	Lasers            [MAX_N_LASERS]*Laser
	Meteors           [MAX_N_METEORS]*Meteor
	MeteorsSpawnTimer *Timer
	MeteorsCoolDown   int
	Score int
}

func NewGame() *Game {
	g := &Game{}

	g.MeteorsCoolDown = METEORS_BASE_COOLDOWN
	g.MeteorsSpawnTimer = NewTimer(METEORS_BASE_COOLDOWN)

	player := NewPlayer(g)
	g.Player = player
	return g
}

func (g *Game) Update() error {
	g.Player.Update()

	for _, laser := range g.Lasers {
		if laser != nil {
			laser.Update()
		}
	}

	for midx, meteor := range g.Meteors {
		if meteor == nil {continue}

		meteor.Update()
		mr := meteor.GetRect()
		pr := g.Player.GetRect()

		// check meteor's collision with the player
		if pr.PixelPerfectCollision(mr) {
			g.RestartGame()
			return nil
		}

		// check collision with the lasers
		for _, laser := range g.Lasers {
			if laser == nil {continue}

			r := laser.GetRect()
			if r.PixelPerfectCollision(mr) {
				g.Meteors[midx] = nil
				g.Score++
			}
		} 
	}

	SpawnMeteor(g)
	return nil
}

func (g *Game) Draw(screen *e.Image) {
	text.Draw(screen, fmt.Sprintf("Score: %d", g.Score), assets.ScoreFont, 20, 40, color.White)
	g.Player.Draw(screen)

	for _, laser := range g.Lasers {
		if laser != nil {
			laser.Draw(screen)
		}
	}
	for _, meteor := range g.Meteors {
		if meteor != nil {
			meteor.Draw(screen)
		}
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			`
Player:
	direction: %v
    position: %v
	size: %v
	speed: %v

Lasers
	current array: %v
	qtd: %d

Meteors
	current array: %v
	qtd: %d
	CoolDown time: %v
	Timer: %v
			`,
			g.Player.Direction,
			g.Player.Position,
			g.Player.Size,
			g.Player.Speed,

			g.Lasers,
			len(g.Lasers),

			g.Meteors,
			len(g.Meteors),
			g.MeteorsCoolDown,
			g.MeteorsSpawnTimer,
		),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) RestartGame() {
	g.Lasers = [MAX_N_LASERS]*Laser{}
	g.Meteors = [MAX_N_METEORS]*Meteor{}
	g.MeteorsSpawnTimer.resetTimer()
	g.MeteorsCoolDown = METEORS_BASE_COOLDOWN
	g.Score = 0
}
