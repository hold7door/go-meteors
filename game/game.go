package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth        = 800
	ScreenHeight       = 600
	meteorSpawnTime    = 1 * time.Second
	baseMeteorVelocity = 0.25
)

type Game struct {
	meteorSpawnTimer *Timer
	player           *Player
	meteors          []*Meteor
	baseVelocity     float64
}

func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.isReady() {
		g.meteorSpawnTimer.Reset()

		m := newMeteor(g.baseVelocity)

		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		baseVelocity:     baseMeteorVelocity,
	}

	g.player = newPlayer()

	return g
}
