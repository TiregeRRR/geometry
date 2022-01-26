package main

import (
	"math/rand"
)

type vector2D struct {
	x, y float64
}

func (v *vector2D) Add(vec vector2D) {
	v.x += vec.x
	v.y += vec.y
}

type cell struct {
	position vector2D
	velocity vector2D
	id       int
}

func getNewCell() *cell {
	c := cell{
		position: vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: vector2D{((rand.Float64() * 2) - 1) * 3, ((rand.Float64() * 2) - 1) * 3},
	}
	return &c
}

func getNewCellCord(x, y int) *cell {
	c := cell{
		position: vector2D{float64(x), float64(y)},
		velocity: vector2D{((rand.Float64() * 2) - 1) * 3, ((rand.Float64() * 2) - 1) * 3},
	}
	return &c
}

func (c *cell) move() {
	c.position.Add(c.velocity)
	if c.position.x+c.velocity.x >= screenWidth || c.position.x+c.velocity.x <= 0 {
		c.velocity.x = -c.velocity.x
	}
	if c.position.y+c.velocity.y >= screenHeight || c.position.y+c.velocity.y <= 0 {
		c.velocity.y = -c.velocity.y
	}
}
