package systems

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type areaOccupyEntity interface {
	engine.Entity
	components.AreaOccupantOwner
}

type AreaOccupySystem struct {
	BaseSystem

	entities EntityMap
}

type EntityOccupiedArea struct {
	Entity engine.Entity
	Area   engine.Entity
}

type EntityStoppedOccupyingArea struct {
	Entity engine.Entity
	Area   engine.Entity
}

func NewAreaOccupySystem(base BaseSystem) *AreaOccupySystem {
	t := &AreaOccupySystem{
		BaseSystem: base,
	}

	t.EventBus.Subscribe(EntityMoved{}, t)
	t.EventBus.Subscribe(EntityReachedTarget{}, t)
	t.EventBus.Subscribe(EntitiesCollided{}, t)
	t.EventBus.Subscribe(EntitiesOutOfCollision{}, t)

	return t
}

func (t AreaOccupySystem) Start() {
}

func (t *AreaOccupySystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntityMoved:
		ent, ok := t.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		entity := ent.(areaOccupyEntity)
		occupant := entity.GetAreaOccupant()
		if occupant.Occupying && occupant.OccupiedArea != nil {
			occupant.Occupying = false
			t.EventBus.Publish(EntityStoppedOccupyingArea{
				Entity: entity,
				Area:   entity.GetAreaOccupant().OccupiedArea,
			})
		}
	case EntityReachedTarget:
		ent, ok := t.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		entity := ent.(areaOccupyEntity)
		occupant := entity.GetAreaOccupant()

		if occupant.OccupiedArea != nil {
			occupant.Occupying = true
			t.EventBus.Publish(EntityOccupiedArea{
				Entity: entity,
				Area:   occupant.OccupiedArea,
			})
		}

	case EntitiesCollided:
		ent, ok := t.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}
		entity := ent.(areaOccupyEntity)
		occupant := entity.GetAreaOccupant()

		_, ok = event.Other.(components.AreaOwner)
		if !ok {
			return
		}

		if occupant.OccupiedArea == nil {
			occupant.OccupiedArea = event.Other
		} else {
			occupant.NextArea = event.Other
		}
	case EntitiesOutOfCollision:
		ent, ok := t.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}
		entity := ent.(areaOccupyEntity)
		occupant := entity.GetAreaOccupant()

		if occupant.OccupiedArea == nil || !occupant.OccupiedArea.Equals(event.Other) {
			return
		}

		if occupant.NextArea != nil {
			occupant.OccupiedArea = occupant.NextArea
			occupant.NextArea = nil
		}
	}
}

func (t AreaOccupySystem) Update(dt float64) {
}

func (t *AreaOccupySystem) Add(entity areaOccupyEntity) {
	t.entities.Add(entity)
}

func (t *AreaOccupySystem) Remove(entity engine.Entity) {
	t.entities.Remove(entity)
}
