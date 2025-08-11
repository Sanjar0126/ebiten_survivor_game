package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

type Game struct {
	World  *ecs.World
	Width  int
	Height int
}

func NewGame(world *ecs.World, width, height int) *Game {
	g := &Game{
		World:  world,
		Width:  width,
		Height: height,
	}

	RegisterComponents(world)

	// Initialize ECS systems
	InitializeSystems(world, width, height)

	return g
}

func (g *Game) Update() error {
	// Update ECS systems
	UpdateSystems(g.World)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render ECS systems
	DrawSystems(g.World, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}
