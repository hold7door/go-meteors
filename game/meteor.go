package game

import (
	"game/assets"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	position Vector
	sprite   *ebiten.Image
	movement Vector
}

func newMeteor(baseVelocity float64) *Meteor {
	target := Vector{
		X: ScreenWidth / 2.0,
		Y: ScreenHeight / 2.0,
	}

	angle := rand.Float64() * 2 * math.Pi
	r := ScreenWidth / 2.0

	velocity := baseVelocity + rand.Float64()*1.5

	pos := Vector{
		X: target.X + r*math.Cos(angle),
		Y: target.Y + r*math.Sin(angle),
	}

	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}
	normalizedDirection := direction.Normalize()

	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	sprite := assets.MeteorSprite[rand.Intn(len(assets.MeteorSprite))]

	return &Meteor{
		position: pos,
		sprite:   sprite,
		movement: movement,
	}
}

func (m *Meteor) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
}
