package modelos

const (
	playerSpeed = 20 
	playerSize = 32
	screenHeight = 500
)

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

func (p *Player) MoveUp() {
	if p.Y > 0 {
		p.Y -= playerSpeed
	}
}

func (p *Player) MoveDown() {
	if p.Y+playerSize < screenHeight {
		p.Y += playerSpeed
	}
}