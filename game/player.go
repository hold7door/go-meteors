package game

import (
	"fmt"
	"game/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const bulletSpawnOffset = 50.0

type Player struct {
	position Vector
	sprite   *ebiten.Image
	rotate   float64
}

func newPlayer() *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		position: pos,
		sprite:   sprite,
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotate -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotate += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		bounds := p.sprite.Bounds()

		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			X: p.position.X + halfW + math.Sin(p.rotate)*bulletSpawnOffset,
			Y: p.position.Y + halfH + math.Cos(p.rotate)*-bulletSpawnOffset,
		}

		bullet := NewBullet(spawnPos, p.rotate)
		fmt.Println(bullet)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotate)
	op.GeoM.Translate(halfW, halfH)
	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}
