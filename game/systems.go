package game

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

func InitializeSystems(world *ecs.World, width, height int) {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Spawn player at the center of the screen
	SpawnPlayer(world, float64(width)/2, float64(height)/2)

	// Spawn initial wave of enemies
	SpawnEnemyWave(world, width, height, 5)
}

func UpdateSystems(world *ecs.World) {
	// Update player
	query := ecs.NewQuery(world, ecs.All(PositionID, PlayerID))
	for query.Next() {
		e := query.Entity()
		pos := (*Position)(world.Get(e, PositionID))
		UpdatePlayer(world, e, pos)
	}

	// Update bullets
	query = ecs.NewQuery(world, ecs.All(PositionID, BulletID))
	for query.Next() {
		e := query.Entity()
		pos := (*Position)(world.Get(e, PositionID))
		UpdateBullet(world, e, pos)
	}

	// Update enemies
	query = ecs.NewQuery(world, ecs.All(PositionID, EnemyID))
	for query.Next() {
		e := query.Entity()
		pos := (*Position)(world.Get(e, PositionID))
		UpdateEnemy(world, e, pos)
	}

	// Check collisions
	CheckCollisions(world)
}

func DrawSystems(world *ecs.World, screen *ebiten.Image) {
	// Draw player
	query := ecs.NewQuery(world, ecs.All(PositionID, PlayerID))
	for query.Next() {
		pos := (*Position)(world.Get(query.Entity(), PositionID))
		DrawPlayer(screen, pos)
	}

	// Draw enemies
	query = ecs.NewQuery(world, ecs.All(PositionID, EnemyID))
	for query.Next() {
		pos := (*Position)(world.Get(query.Entity(), PositionID))
		DrawEnemy(screen, pos)
	}

	// Draw bullets
	query = ecs.NewQuery(world, ecs.All(PositionID, BulletID))
	for query.Next() {
		pos := (*Position)(world.Get(query.Entity(), PositionID))
		DrawBullet(screen, pos)
	}
}
