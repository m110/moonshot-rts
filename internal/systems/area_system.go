package systems

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type areaEntity interface {
	engine.Entity
	components.AreaOwner
}

type AreaSystem struct {
	BaseSystem
	entities EntityMap
}

func NewAreaSystem(base BaseSystem) *AreaSystem {
	a := &AreaSystem{
		BaseSystem: base,
	}

	a.EventBus.Subscribe(EntityOccupiedArea{}, a)
	a.EventBus.Subscribe(EntityStoppedOccupyingArea{}, a)

	return a
}

func (a *AreaSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntityOccupiedArea:
		entity, ok := a.entities.ByID(event.Area.ID())
		if !ok {
			return
		}

		area := entity.(areaEntity).GetArea()
		area.Occupants++
		if area.Occupants > 0 {
			area.Overlay.GetDrawable().Enable()
		}
	case EntityStoppedOccupyingArea:
		entity, ok := a.entities.ByID(event.Area.ID())
		if !ok {
			return
		}

		area := entity.(areaEntity).GetArea()
		area.Occupants--
		if area.Occupants <= 0 {
			area.Overlay.GetDrawable().Disable()
		}
	}
}

func (a AreaSystem) Start() {
}

func (a *AreaSystem) Update(dt float64) {
}

func (a *AreaSystem) Add(entity areaEntity) {
	a.entities.Add(entity)
}

func (a *AreaSystem) Remove(entity engine.Entity) {
	a.entities.Remove(entity)
}
