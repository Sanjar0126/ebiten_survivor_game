package game

import (
	"github.com/mlange-42/arche/ecs"
)

// Position component
type Position struct {
	X, Y float64
}

// Velocity component
type Velocity struct {
	X, Y float64
}

// Player component
type Player struct{}

// Enemy component
type Enemy struct{}

// Bullet component
type Bullet struct {
	Damage int
}

// Component IDs
var (
	PositionID ecs.ID
	VelocityID ecs.ID
	PlayerID   ecs.ID
	EnemyID    ecs.ID
	BulletID   ecs.ID
)

func init() {
	// Register components
	PositionID = ecs.ComponentID[Position]()
	VelocityID = ecs.ComponentID[Velocity]()
	PlayerID = ecs.ComponentID[Player]()
	EnemyID = ecs.ComponentID[Enemy]()
	BulletID = ecs.ComponentID[Bullet]()
}
