package game

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

var (
	lastShotTime = time.Now()
)

func SpawnPlayer(world *ecs.World, x, y float64) {
	e := world.NewEntity(PositionID, VelocityID, PlayerID)
	world.Set(e, PositionID, &Position{X: x, Y: y})
	world.Set(e, VelocityID, &Velocity{})
}

func UpdatePlayer(world *ecs.World, e ecs.Entity, pos *Position) {
	// Movement
	speed := 2.0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		pos.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		pos.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		pos.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		pos.X += speed
	}

	// Auto-shoot every 500ms
	if time.Since(lastShotTime) > 500*time.Millisecond {
		SpawnBullet(world, pos.X, pos.Y)
		lastShotTime = time.Now()
	}
}

func DrawPlayer(screen *ebiten.Image, pos *Position) {
	// Placeholder for player drawing
	// Replace with your own asset drawing logic
	DrawRectangle(screen, pos.X, pos.Y, 20, 20, color.RGBA{B: 255, A: 255})
}
