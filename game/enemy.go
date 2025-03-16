package game

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

func SpawnEnemy(world *ecs.World, x, y float64) {
	e := world.NewEntity(PositionID, VelocityID, EnemyID)
	world.Set(e, PositionID, &Position{X: x, Y: y})
	world.Set(e, VelocityID, &Velocity{})
}

func SpawnEnemyWave(world *ecs.World, width, height, count int) {
	for i := 0; i < count; i++ {
		x := rand.Float64() * float64(width)
		y := rand.Float64() * float64(height)
		SpawnEnemy(world, x, y)
	}
}

func UpdateEnemy(world *ecs.World, e ecs.Entity, pos *Position) {
	// Retrieve player position
	var playerPos *Position
	query := ecs.NewQuery(world, ecs.All(PositionID, PlayerID))
	for query.Next() {
		playerPos = (*Position)(world.Get(query.Entity(), PositionID))
	}

	if playerPos == nil {
		return // No player found
	}

	// Move towards player
	speed := 1.0
	angle := math.Atan2(playerPos.Y-pos.Y, playerPos.X-pos.X)
	pos.X += speed * math.Cos(angle)
	pos.Y += speed * math.Sin(angle)
}

func DrawEnemy(screen *ebiten.Image, pos *Position) {
	// Placeholder for enemy drawing
	// Replace with your own asset drawing logic
	DrawRectangle(screen, pos.X, pos.Y, 20, 20, color.RGBA{R: 255, A: 255})
}
