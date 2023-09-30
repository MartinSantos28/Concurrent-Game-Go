package modelos

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
