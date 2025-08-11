package game

import (
	"github.com/mlange-42/arche/ecs"
)

func CheckCollisions(world *ecs.World) {
	// Create lists to track entities that need to be removed
	bulletsToRemove := []ecs.Entity{}
	enemiesToRemove := []ecs.Entity{}
	
	bulletFilter := ecs.All(PositionID, BulletID)
	bulletQuery := world.Query(bulletFilter)
	
	enemyFilter := ecs.All(PositionID, EnemyID)
	enemyQuery := world.Query(enemyFilter)

	// Store enemy positions for efficiency
	type enemyData struct {
		entity ecs.Entity
		pos    *Position
	}
	
	enemies := make([]enemyData, 0)
	
	// First pass: collect all enemy data
	for enemyQuery.Next() {
		entity := enemyQuery.Entity()
		pos := (*Position)(world.Get(entity, PositionID))
		enemies = append(enemies, enemyData{entity: entity, pos: pos})
	}
	
	// Second pass: check for collisions and mark entities for removal
	for bulletQuery.Next() {
		bulletEntity := bulletQuery.Entity()
		bulletPos := (*Position)(world.Get(bulletEntity, PositionID))

		for _, enemy := range enemies {
			if IsColliding(bulletPos, enemy.pos, 5, 10, 20, 20) {
				// Don't remove immediately, just mark for removal
				bulletsToRemove = append(bulletsToRemove, bulletEntity)
				enemiesToRemove = append(enemiesToRemove, enemy.entity)
				break // Bullet hit something, move to next bullet
			}
		}
	}
	
	// After all iterations are complete, remove the marked entities
	for _, entity := range bulletsToRemove {
		world.Remove(entity)
	}
	
	for _, entity := range enemiesToRemove {
		world.Remove(entity)
	}
}

func IsColliding(pos1, pos2 *Position, w1, h1, w2, h2 float64) bool {
	return pos1.X < pos2.X+w2 &&
		pos1.X+w1 > pos2.X &&
		pos1.Y < pos2.Y+h2 &&
		pos1.Y+h1 > pos2.Y
}