package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

func SpawnBullet(world *ecs.World, x, y float64) {
	e := world.NewEntity(PositionID, VelocityID, BulletID)
	world.Set(e, PositionID, &Position{X: x, Y: y})
	world.Set(e, VelocityID, &Velocity{X: 0, Y: -5}) // Bullets move upward
	world.Set(e, BulletID, &Bullet{Damage: 1})
}

func UpdateBullet(world *ecs.World, e ecs.Entity, pos *Position) {
	// Update position based on velocity
	if vel, ok := world.Get(e, VelocityID).(*Velocity); ok {
		pos.X += vel.X
		pos.Y += vel.Y
	}

	// Remove bullet if it goes off-screen
	if pos.Y < 0 {
		world.Remove(e)
	}
}

func DrawBullet(screen *ebiten.Image, pos *Position) {
	// Placeholder for bullet drawing
	// Replace with your own asset drawing logic
	DrawRectangle(screen, pos.X, pos.Y, 5, 10, color.RGBA{B: 255, A: 255})
}
