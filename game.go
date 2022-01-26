package main

import (
	"image/color"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth      = 1920
	screenHeight     = 1080
	maxNumberOfCells = 500
	viewRadius       = 140
)

var circle *ebiten.Image

func cellsInit() {
	for i := 0; i < numberOfCells; i++ {
		cells[i] = getNewCell()
		cells[i].id = i
	}
}

var (
	numberOfCells int = 120
	cells         [maxNumberOfCells]*cell
)

var grey = color.RGBA{
	R: 100,
	G: 100,
	B: 100,
	A: 255,
}

type Game struct{}

func (g *Game) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if numberOfCells < maxNumberOfCells {
			x, y := ebiten.CursorPosition()
			c := getNewCellCord(x, y)
			cells[numberOfCells] = c
			numberOfCells++
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		if numberOfCells > 0 {
			numberOfCells--
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < numberOfCells; i++ {
		cells[i].move()
	}
	for i := 0; i < numberOfCells; i++ {
		draw(cells[i], screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func draw(c *cell, screen *ebiten.Image) {
	posX, posY := int(c.position.x), int(c.position.y)
	fromX, toX := posX-viewRadius, posX+viewRadius
	fromY, toY := posY-viewRadius, posY+viewRadius
	if fromX < 0 {
		fromX = 0
	}
	if toX >= screenWidth {
		toX = screenWidth - 1
	}
	if fromY < 0 {
		fromY = 0
	}
	if toY >= screenHeight {
		toY = screenHeight - 1
	}
	for j := c.id + 1; j < numberOfCells; j++ {
		newC := cells[j]
		dist := math.Pow(c.position.x-newC.position.x, 2) + math.Pow(c.position.y-newC.position.y, 2)
		if dist <= viewRadius*viewRadius {
			var orange = color.RGBA{
				R: 225,
				G: 165,
				B: 0,
				A: uint8((1 - (dist / (viewRadius * viewRadius))) * 255),
			}
			ebitenutil.DrawLine(screen, c.position.x, c.position.y, newC.position.x, newC.position.y, orange)
		}
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.75, 0.75)
	op.GeoM.Translate(c.position.x-6, c.position.y-6)
	op.ColorM.Apply(grey)
	screen.DrawImage(circle, op)
}
