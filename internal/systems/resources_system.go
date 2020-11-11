package systems

import (
	"github.com/m110/moonshot-rts/internal/archetypes"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type resourcesEntity interface {
	engine.Entity
	components.AreaOccupantOwner
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

	tileOverlays map[resourcesEntity]archetypes.Object
}

func NewResourcesSystem(base BaseSystem) *ResourcesSystem {
	r := &ResourcesSystem{
		BaseSystem:   base,
		tileOverlays: map[resourcesEntity]archetypes.Object{},
	}

	r.EventBus.Subscribe(EntityOccupiedArea{}, r)
	r.EventBus.Subscribe(EntityStoppedOccupyingArea{}, r)

	return r
}

func (r *ResourcesSystem) Start() {
	r.updateResources()
}

func (r *ResourcesSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntityOccupiedArea:
		ent, ok := r.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		tile, ok := event.Area.(components.ResourcesSourceOwner)
		if !ok {
			return
		}

		entity := ent.(resourcesEntity)

		entity.GetResourcesCollector().CurrentResources = tile.GetResourcesSource().Resources
		entity.GetResourcesCollector().Collecting = true
		r.updateResources()
	case EntityStoppedOccupyingArea:
		ent, ok := r.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		entity := ent.(resourcesEntity)
		entity.GetResourcesCollector().Collecting = false
		r.updateResources()
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

	/*
		overlay := archetypes.NewOverlay(
			r.Config.TileMap.TileWidth,
			r.Config.TileMap.TileHeight,
			engine.PivotBottom,
		)

		entity.GetWorldSpace().AddChild(overlay)
		r.Spawner.Spawn(overlay)
	*/
}

func (r *ResourcesSystem) Remove(entity engine.Entity) {
	r.entities.Remove(entity)
}
