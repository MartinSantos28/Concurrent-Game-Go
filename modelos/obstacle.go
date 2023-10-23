package modelos
import (
	"math/rand"
)
type Obstacle struct {
	X float64
	Y float64
}

func NewObstacle(x, y float64) *Obstacle {
	return &Obstacle{
		X: x,
		Y: y,
	}
}

func (o *Obstacle) Move(speed float64) {
	o.X += speed
}

func (o *Obstacle) IsOutOfBounds(screenWidth float64) bool {
	return o.X > screenWidth
}

func GenerateRandomObstacle(screenWidth, obstacleSize float64) *Obstacle {
	y := float64(rand.Intn(int(screenHeight - obstacleSize)))
	return NewObstacle(-obstacleSize, y)
}