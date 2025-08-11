package game

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche/ecs"
)

func InitializeSystems(world *ecs.World, width, height int) {
	// Modern approach to seeding random number generator
	// rand.Seed is deprecated in newer Go versions
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Spawn player at the center of the screen
	SpawnPlayer(world, float64(width)/2, float64(height)/2)

	// Spawn initial wave of enemies
	SpawnEnemyWave(world, width, height, 5)
}

func UpdateSystems(world *ecs.World) {
	// Step 1: Update player
	playerFilter := ecs.All(PositionID, PlayerID)
	playerQuery := world.Query(playerFilter)
	for playerQuery.Next() {
		e := playerQuery.Entity()
		pos := (*Position)(world.Get(e, PositionID))
		UpdatePlayer(world, e, pos)
	}
	
	// Step 2: Process any bullet spawning requests
	// This happens AFTER player updates but BEFORE bullet movement
	ProcessBulletQueue(world)

	// Step 3: Update bullets
	bulletFilter := ecs.All(PositionID, BulletID)
	bulletQuery := world.Query(bulletFilter)
	for bulletQuery.Next() {
		e := bulletQuery.Entity()
		pos := (*Position)(world.Get(e, PositionID))
		UpdateBullet(world, e, pos)
	}

	// Step 4: Update enemies
	enemyFilter := ecs.All(PositionID, EnemyID)
	enemyQuery := world.Query(enemyFilter)
	for enemyQuery.Next() {
		e := enemyQuery.Entity()
		pos := (*Position)(world.Get(e, PositionID))
		UpdateEnemy(world, e, pos)
	}

	// Step 5: Clean up bullets that went off-screen
	CleanupBullets(world)

	// Step 6: Check collisions 
	CheckCollisions(world)
}

func DrawSystems(world *ecs.World, screen *ebiten.Image) {
	// Draw player
	playerFilter := ecs.All(PositionID, PlayerID)
	playerQuery := world.Query(playerFilter)
	for playerQuery.Next() {
		pos := (*Position)(world.Get(playerQuery.Entity(), PositionID))
		DrawPlayer(screen, pos)
	}

	// Draw enemies
	enemyFilter := ecs.All(PositionID, EnemyID)
	enemyQuery := world.Query(enemyFilter) 
	for enemyQuery.Next() {
		pos := (*Position)(world.Get(enemyQuery.Entity(), PositionID))
		DrawEnemy(screen, pos)
	}

	// Draw bullets
	bulletFilter := ecs.All(PositionID, BulletID)
	bulletQuery := world.Query(bulletFilter)
	for bulletQuery.Next() {
		pos := (*Position)(world.Get(bulletQuery.Entity(), PositionID))
		DrawBullet(screen, pos)
	}
}