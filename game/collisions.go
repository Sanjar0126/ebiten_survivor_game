package game

import (
	"github.com/mlange-42/arche/ecs"
)

func CheckCollisions(world *ecs.World) {
	bulletQuery := ecs.NewQuery(world, ecs.All(PositionID, BulletID))
	enemyQuery := ecs.NewQuery(world, ecs.All(PositionID, EnemyID))

	for bulletQuery.Next() {
		bulletEntity := bulletQuery.Entity()
		bulletPos := (*Position)(world.Get(bulletEntity, PositionID))

		for enemyQuery.Next() {
			enemyEntity := enemyQuery.Entity()
			enemyPos := (*Position)(world.Get(enemyEntity, PositionID))

			if IsColliding(bulletPos, enemyPos, 5, 10, 20, 20) {
				// Handle collision (e.g., remove both entities)
				world.Remove(bulletEntity)
				world.Remove(enemyEntity)
			}
		}
	}
}

func IsColliding(pos1, pos2 *Position, w1, h1, w2, h2 float64) bool {
	return pos1.X < pos2.X+w2 &&
		pos1.X+w1 > pos2.X &&
		pos1.Y < pos2.Y+h2 &&
		pos1.Y+h1 > pos2.Y
}
