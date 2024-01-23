package game

import (
	"fmt"
	"game/assets"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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
	bullets          []*Bullet
	baseVelocity     float64
	score            int
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

	for _, b := range g.bullets {
		b.Update()
	}

	// TODO: Understand this

	// Check for meteor/bullet collisions
	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.Collider().Intersects(b.Collider()) {
				updatedMeteors := g.meteors[:i]
				updatedBullets := g.bullets[:j]

				if i+1 < len(g.meteors) {
					g.meteors = append(updatedMeteors, g.meteors[i+1:]...)
				}
				if j+1 < len(g.bullets) {
					g.bullets = append(updatedBullets, g.bullets[j+1:]...)
				}
				g.score++
			}
		}
	}

	// Check for meteor/player collisions
	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
			break
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, b := range g.bullets {
		b.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, ScreenWidth/2-100, 50, color.White)
}

func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Reset() {
	g.player = newPlayer(g)
	g.meteors = nil
	g.bullets = nil
	g.score = 0
	g.meteorSpawnTimer.Reset()
	g.baseVelocity = baseMeteorVelocity
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		baseVelocity:     baseMeteorVelocity,
	}

	g.player = newPlayer(g)

	return g
}
