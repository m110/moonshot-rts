package systems

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type resourcesEntity interface {
	engine.Entity
	components.ColliderOwner
	components.MovableOwner
	components.ResourcesCollectorOwner
}

type ResourcesUpdated struct {
	Resources components.Resources
}

type ResourcesSystem struct {
	BaseSystem

	entities  EntityMap
	resources components.Resources
}

func NewResourcesSystem(base BaseSystem) *ResourcesSystem {
	return &ResourcesSystem{
		BaseSystem: base,
	}
}

func (r *ResourcesSystem) Start() {
	r.EventBus.Subscribe(EntityMoved{}, r)
	r.EventBus.Subscribe(EntityReachedTarget{}, r)
	r.EventBus.Subscribe(EntitiesCollided{}, r)
	r.updateResources()
}

func (r *ResourcesSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntityMoved:
		ent, ok := r.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}
		entity := ent.(resourcesEntity)
		entity.GetResourcesCollector().Collecting = false
		r.updateResources()

	case EntityReachedTarget:
		ent, ok := r.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}
		entity := ent.(resourcesEntity)
		entity.GetResourcesCollector().Collecting = true
		r.updateResources()

	case EntitiesCollided:
		ent, ok := r.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}
		entity := ent.(resourcesEntity)

		source, ok := event.Other.(components.ResourcesSourceOwner)
		if !ok {
			return
		}

		entity.GetResourcesCollector().CurrentResources = source.GetResourcesSource().Resources
	}
}

func (r *ResourcesSystem) updateResources() {
	r.resources = components.Resources{}
	for _, e := range r.entities.All() {
		entity := e.(resourcesEntity)
		if entity.GetResourcesCollector().Collecting {
			r.resources.Update(entity.GetResourcesCollector().CurrentResources)
		}
	}

	r.EventBus.Publish(ResourcesUpdated{Resources: r.resources})
}

func (r ResourcesSystem) Update(dt float64) {}

func (r *ResourcesSystem) Add(entity resourcesEntity) {
	r.entities.Add(entity)
}

func (r *ResourcesSystem) Remove(entity engine.Entity) {
	r.entities.Remove(entity)
}
