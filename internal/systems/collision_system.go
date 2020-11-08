package systems

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

// TODO Come up with a better way to define this
var collisionsMatrix = map[components.CollisionLayer][]components.CollisionLayer{
	components.CollisionLayerGround: {
		components.CollisionLayerUnits,
	},
	components.CollisionLayerUnits: {
		components.CollisionLayerGround,
		components.CollisionLayerBuildings,
	},
	components.CollisionLayerBuildings: {
		components.CollisionLayerUnits,
	},
}

type collisionEntity interface {
	engine.Entity
	components.WorldSpaceOwner
	components.ColliderOwner
}

type EntitiesCollided struct {
	Entity engine.Entity
	Other  engine.Entity
}

type EntitiesOutOfCollision struct {
	Entity engine.Entity
	Other  engine.Entity
}

type CollisionSystem struct {
	BaseSystem
	entities              EntityMap
	recentlyMovedEntities EntityList
}

func NewCollisionSystem(base BaseSystem) *CollisionSystem {
	c := &CollisionSystem{
		BaseSystem: base,
	}

	c.EventBus.Subscribe(EntityMoved{}, c)

	return c
}

func (c CollisionSystem) Start() {
}

func (c *CollisionSystem) Update(dt float64) {
	for _, e := range c.recentlyMovedEntities.All() {
		for _, o := range c.entities.All() {
			entity := e.(collisionEntity)
			other := o.(collisionEntity)

			if entity.Equals(other) {
				continue
			}

			// TODO This should be moved to a common func
			entityBounds := entity.GetCollider().Bounds
			entityBounds.Position = entityBounds.Position.Add(entity.GetWorldSpace().WorldPosition())

			otherBounds := other.GetCollider().Bounds
			otherBounds.Position = otherBounds.Position.Add(other.GetWorldSpace().WorldPosition())

			intersects := entityBounds.Intersects(otherBounds)

			// Already collide with each other
			if entity.GetCollider().HasCollision(other) {
				// No longer collide
				if !intersects {
					entity.GetCollider().RemoveCollision(other)
					other.GetCollider().RemoveCollision(entity)

					c.EventBus.Publish(EntitiesOutOfCollision{
						Entity: entity,
						Other:  other,
					})
					c.EventBus.Publish(EntitiesOutOfCollision{
						Entity: other,
						Other:  entity,
					})
				}
				continue
			}

			if !intersects {
				continue
			}

			collisions := collisionsMatrix[entity.GetCollider().Layer]
			collides := false
			for _, l := range collisions {
				if l == other.GetCollider().Layer {
					collides = true
					break
				}
			}

			if !collides {
				continue
			}

			entity.GetCollider().AddCollision(other)
			other.GetCollider().AddCollision(entity)

			c.EventBus.Publish(EntitiesCollided{
				Entity: entity,
				Other:  other,
			})
			c.EventBus.Publish(EntitiesCollided{
				Entity: other,
				Other:  entity,
			})
		}
	}

	c.recentlyMovedEntities.Clear()
}

func (c *CollisionSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntityMoved:
		_, ok := c.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		c.recentlyMovedEntities.Add(event.Entity)
	}
}

func (c *CollisionSystem) Add(entity collisionEntity) {
	c.entities.Add(entity)
	c.recentlyMovedEntities.Add(entity)
}

func (c *CollisionSystem) Remove(entity engine.Entity) {
	c.entities.Remove(entity)
}
