package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_"image/png"
)

var (
	playerImage   *ebiten.Image
	obstacleImage *ebiten.Image
)

func loadImages() error {
	var err error

	playerImage, _, err = ebitenutil.NewImageFromFile("./assets/player.png") // Ajustado para encontrar la imagen
	if err != nil {
		return err
	}

	obstacleImage, _, err = ebitenutil.NewImageFromFile("./assets/obstacle.png") // Ajustado
	return err
}
