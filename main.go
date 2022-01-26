package main

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	file, err := os.Open("./circle.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	c, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	circle = ebiten.NewImageFromImage(c)
	game := &Game{}
	cellsInit()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Cells")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
