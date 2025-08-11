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
	// Initialization code that does not require arguments
}

func RegisterComponents(world *ecs.World) {
	// Register components
	PositionID = ecs.ComponentID[Position](world)
	VelocityID = ecs.ComponentID[Velocity](world)
	PlayerID = ecs.ComponentID[Player](world)
	EnemyID = ecs.ComponentID[Enemy](world)
	BulletID = ecs.ComponentID[Bullet](world)
}
