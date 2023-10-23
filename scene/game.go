package game

import (
	"dino/modelos"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"fmt"
	"math/rand"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"time"
	
)

const (
	screenWidth   = 800
	screenHeight  = 500
	playerSize    = 32
	obstacleSize  = 50
	playerSpeed   = 20
	obstacleSpeed = 5
	
)
func NewGame() *Game {
	return &Game{
		Player: modelos.NewPlayer(),
		closeChannel:  make(chan struct{}),
		score:         0,
		timeRemaining: 60,
		lives:         3,
		obstacleSpeed: 5,
	}
}

type Game struct {
	
	Player       *modelos.Player
	Obstacles    []*modelos.Obstacle

	closeChannel chan struct{}
	lives        int
	timeRemaining float64
	score        int
	obstacleSpeed float64
	hasStarted bool
	hasLose bool 
}


func (g *Game) moveObstacles() {
	for i := len(g.Obstacles) - 1; i >= 0; i-- {
		obstacle := g.Obstacles[i]
		obstacle.Move(g.obstacleSpeed) 
		if obstacle.IsOutOfBounds(screenWidth) { 
			g.score++
			
			g.Obstacles = append(g.Obstacles[:i], g.Obstacles[i+1:]...)
		}
	}
}



func (g *Game) generateObstacles() {
	for {
		select {
		case <-g.closeChannel:
			return
		default:
			newObstacle := modelos.GenerateRandomObstacle(screenWidth, obstacleSize)
			g.Obstacles = append(g.Obstacles, newObstacle)
			time.Sleep(time.Second)
		}
	}
}


func (g *Game) checkCollisions() {
    for i := len(g.Obstacles) - 1; i >= 0; i-- {  
        obstacle := g.Obstacles[i]
        ox, oy := obstacle.X, obstacle.Y

        if g.Player.X < ox+obstacleSize && g.Player.X+playerSize > ox &&
            g.Player.Y < oy+obstacleSize && g.Player.Y+playerSize > oy {
            g.lives--
            if g.lives == 0 {
                g.hasLose = true
            } else {
                log.Printf("¡Has sido golpeado! Te quedan %d vidas", g.lives)
                
                g.Obstacles = append(g.Obstacles[:i], g.Obstacles[i+1:]...)
            }
        }
    }
}


func (g *Game) Update() error {
	if !g.hasStarted {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.hasStarted = true
			go g.generateObstacles()
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.Player.Y > 0 {
		g.Player.MoveUp() 
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.Player.Y+playerSize < screenHeight {
		g.Player.MoveDown() 
	}

	go g.moveObstacles()
	go g.checkCollisions()
	g.timeRemaining -= 1.0 / 60.0
	if g.timeRemaining <= 0 {
		log.Fatal("¡Tiempo agotado!")
	}
	if int(g.timeRemaining)%10 == 0 {
		g.obstacleSpeed += 0.1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    if !g.hasStarted {
        ebitenutil.DebugPrint(screen, "Presiona ENTER para comenzar!")
        return
    }
    
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(g.Player.X, g.Player.Y)
	screen.DrawImage(playerImage, op)


    if g.hasLose {
        ebitenutil.DebugPrint(screen, "Has perdido")
        return
    }
    
    for _, obstacle := range g.Obstacles {
        op := &ebiten.DrawImageOptions{}
        op.GeoM.Scale(0.2, 0.2)
        op.GeoM.Translate(obstacle.X, obstacle.Y)  
        screen.DrawImage(obstacleImage, op)
    }

    ebitenutil.DebugPrint(screen, fmt.Sprintf("Vidas: %d\nTiempo restante: %02d:%02d\nPuntuación: %d",
        g.lives, int(g.timeRemaining)/60, int(g.timeRemaining)%60, g.score))
}


func (g *Game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}


func Run() error {
	if err := loadImages(); err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())
	game := NewGame()
	game.Update()
	go game.generateObstacles()
	
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Esquiva los pichis!")
	return ebiten.RunGame(game)
}
