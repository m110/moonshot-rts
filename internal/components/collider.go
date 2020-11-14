package components

import "github.com/m110/moonshot-rts/internal/engine"

type CollisionLayer int

const (
	CollisionLayerGround CollisionLayer = iota
	CollisionLayerUnits
	CollisionLayerBuildings
)

type Collider struct {
	Bounds  engine.Rect
	Layer   CollisionLayer
	Overlay DrawableOwner

	collisions map[engine.EntityID]engine.Entity
}

func (c Collider) HasCollision(other engine.Entity) bool {
	_, ok := c.collisions[other.ID()]
	return ok
}

func (c *Collider) AddCollision(other engine.Entity) {
	if c.collisions == nil {
		c.collisions = map[engine.EntityID]engine.Entity{}
	}

	c.collisions[other.ID()] = other
}

func (c *Collider) RemoveCollision(other engine.Entity) {
	delete(c.collisions, other.ID())
}
