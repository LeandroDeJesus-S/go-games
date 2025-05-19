package game

import (
	"fmt"
	"image/color"
	"spacegame/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)


type Game struct {
	player *Player
	lasers []*Laser
	meteors []*Meteor
	meteorSpawnTimer *Timer
	score int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(25),
	}
	player := NewPlayer(g)
	g.player = player
	return g
}

func (g *Game) Update() error {
	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.isReady() {
		g.meteors = append(g.meteors, NewMeteor())
		g.meteorSpawnTimer.resetTimer()
	}

	for i, m := range g.meteors {
		m.Update()

		r := m.Collider()
		if r.CollidedWith(g.player.Collider()) {
			g.resetGame()
			return nil
		}

		for j, l := range g.lasers {
			r2 := l.Collider()
			if r2.CollidedWith(r) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score ++
			}
		}
	}

	

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), assets.FontUi, 20, 40, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) addLaser(l *Laser) {
	g.lasers = append(g.lasers, l)
}

func (g *Game) resetGame() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.resetTimer()
	g.score = 0
}