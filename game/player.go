package game

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

var (
	lastShotTime = time.Now()
	bulletQueue  []BulletSpawnRequest // Queue for bullets to spawn
)

// BulletSpawnRequest represents a request to spawn a bullet
type BulletSpawnRequest struct {
	X, Y float64
}

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

	// Queue bullet spawn instead of spawning directly
	if time.Since(lastShotTime) > 500*time.Millisecond {
		// Add to bullet spawn queue instead of creating immediately
		bulletQueue = append(bulletQueue, BulletSpawnRequest{X: pos.X, Y: pos.Y})
		lastShotTime = time.Now()
	}
}

// ProcessBulletQueue spawns all queued bullets
// Call this outside of any world iteration
func ProcessBulletQueue(world *ecs.World) {
	for _, req := range bulletQueue {
		// Create the bullet entity
		e := world.NewEntity(PositionID, VelocityID, BulletID)
		world.Set(e, PositionID, &Position{X: req.X, Y: req.Y})
		world.Set(e, VelocityID, &Velocity{X: 0, Y: -5}) // Bullets move upward
		world.Set(e, BulletID, &Bullet{Damage: 1})
	}

	// Clear the queue after processing
	bulletQueue = []BulletSpawnRequest{}
}

func DrawPlayer(screen *ebiten.Image, pos *Position) {
	// Placeholder for player drawing
	DrawRectangle(screen, pos.X, pos.Y, 20, 20, color.RGBA{B: 255, A: 255})
}
