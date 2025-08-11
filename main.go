package main

import (
	"log"

	"github.com/Sanjar0126/ebiten_survivor_game/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func main() {
	// Initialize ECS world
	world := ecs.NewWorld()

	// Create Game
	g := game.NewGame(&world, ScreenWidth, ScreenHeight)

	// Run Game
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}