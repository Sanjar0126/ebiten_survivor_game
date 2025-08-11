package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

// We remove SpawnBullet function since bullet spawning is now handled
// through the queue system in player.go

func UpdateBullet(world *ecs.World, e ecs.Entity, pos *Position) {
	// Update position based on velocity
	vel := (*Velocity)(world.Get(e, VelocityID))
	pos.X += vel.X
	pos.Y += vel.Y

	// Tracking of off-screen bullets is now done in CleanupBullets
}

// Function to remove bullets that are off-screen
// This should be called after all bullet updates are complete
func CleanupBullets(world *ecs.World) {
	bulletsToRemove := []ecs.Entity{}
	
	bulletFilter := ecs.All(PositionID, BulletID)
	bulletQuery := world.Query(bulletFilter)
	
	for bulletQuery.Next() {
		entity := bulletQuery.Entity()
		pos := (*Position)(world.Get(entity, PositionID))
		
		// Check if bullet is off-screen
		if pos.Y < 0 {
			bulletsToRemove = append(bulletsToRemove, entity)
		}
	}
	
	// Remove bullets after iteration is complete
	for _, entity := range bulletsToRemove {
		world.Remove(entity)
	}
}

func DrawBullet(screen *ebiten.Image, pos *Position) {
	// Placeholder for bullet drawing
	DrawRectangle(screen, pos.X, pos.Y, 5, 10, color.RGBA{B: 255, A: 255})
}