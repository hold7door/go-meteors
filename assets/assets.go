package assets

import (
	"embed"
	"image"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

var PlayerSprite = mustLoadImage("assets/PNG/playerShip1_blue.png")
var MeteorSprite = mustLoadImages("assets/PNG/Meteors/*.png")
var LaserSprite = mustLoadImage("assets/PNG/Lasers/laserBlue01.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(name string) []*ebiten.Image {
	matches, err := fs.Glob(assets, name)
	if err != nil {
		panic(err)
	}
	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}
	return images
}
