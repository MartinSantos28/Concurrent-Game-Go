package game

import (
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
		playerX:       (screenWidth - playerSize) / 2, 
		playerY:       screenHeight - playerSize*4,
		closeChannel:  make(chan struct{}),
		score:         0,
		timeRemaining: 60,
		lives:         3,
		obstacleSpeed: 5,
	}
}

type Game struct {
	playerX      float64
	playerY      float64
	obstacles    []float64
	closeChannel chan struct{}
	lives        int
	timeRemaining float64
	score        int
	obstacleSpeed float64
	hasStarted bool 
}

// Resto de las funciones como NewGame(), Update(), Draw(), etc.
func (g *Game) moveObstacles() {
	for i := 0; i < len(g.obstacles); i += 2 {
		g.obstacles[i] += g.obstacleSpeed
		if g.obstacles[i] > screenWidth {
			g.score++
			g.obstacles = append(g.obstacles[:i], g.obstacles[i+2:]...)
			i -= 2
		}
	}
}



func (g *Game) generateObstacles() {
	for {
		select {
		case <-g.closeChannel:
			return
		default:
			y := float64(rand.Intn(screenHeight - obstacleSize))
			g.obstacles = append(g.obstacles, -obstacleSize, y)
			time.Sleep(time.Second)
		}
	}
}

func (g *Game) checkCollisions() {
	for i := 0; i < len(g.obstacles); i += 2 {
		ox, oy := g.obstacles[i], g.obstacles[i+1]

		if g.playerX < ox+obstacleSize && g.playerX+playerSize > ox &&
			g.playerY < oy+playerSize && g.playerY+playerSize > oy {
			g.lives-- // Decrementa el contador de vidas
			if g.lives == 0 {
				log.Fatal("¡Juego Terminado!")
			} else {
				log.Printf("¡Has sido golpeado! Te quedan %d vidas", g.lives)
				g.obstacles = append(g.obstacles[:i], g.obstacles[i+2:]...) 
				i -= 2
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
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.playerY > 0 {
		g.playerY -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.playerY+playerSize < screenHeight {
		g.playerY += playerSpeed
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
	op.GeoM.Translate(g.playerX, g.playerY)
	screen.DrawImage(playerImage, op)

	
	for i := 0; i < len(g.obstacles); i += 2 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.2, 0.2)
		op.GeoM.Translate(g.obstacles[i], g.obstacles[i+1])
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
	ebiten.SetWindowTitle("Dodge the Obstacles!")
	return ebiten.RunGame(game)
}
