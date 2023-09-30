package modelos

type Player struct {
	X float64
	Y float64
}

func NewPlayer() *Player {
	return &Player{
		X: (800 - 32) / 2,
		Y: 500 - 32 * 4,
	}
}
