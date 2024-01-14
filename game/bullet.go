package game

import (
	"game/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const bulletSpeedPerSecond = 350

type Bullet struct {
	position Vector
	sprite   *ebiten.Image
	rotation float64
}

func NewBullet(pos Vector, rotation float64) *Bullet {
	sprite := assets.LaserSprite
	bounds := sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos.X -= halfW
	pos.Y -= halfH

	b := &Bullet{
		position: pos,
		sprite:   sprite,
		rotation: rotation,
	}

	return b
}

func (b *Bullet) Update() {
	speed := bulletSpeedPerSecond / float64(ebiten.TPS())

	b.position.X += math.Sin(b.rotation) * speed
	b.position.Y += math.Cos(b.rotation) * speed
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	bounds := b.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(b.position.X, b.position.Y)

	screen.DrawImage(b.sprite, op)
}
