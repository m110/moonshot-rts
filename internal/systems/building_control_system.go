package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
	"github.com/m110/rts/internal/units"
)

type buildingControlEntity interface {
	engine.Entity
	components.WorldSpaceOwner
	// TODO building interface
}

type BuildingControlSystem struct {
	base BaseSystem

	entities EntityList

	activeEntity buildingControlEntity
}

func NewBuildingControlSystem(config Config, eventBus *engine.EventBus, spawner spawner) *BuildingControlSystem {
	return &BuildingControlSystem{
		base: NewBaseSystem(config, eventBus, spawner),
	}
}

func (b *BuildingControlSystem) Start() {
	b.base.EventBus.Subscribe(EntitySelected{}, b)
	b.base.EventBus.Subscribe(EntityUnselected{}, b)
}

func (b BuildingControlSystem) Update(dt float64) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if b.activeEntity != nil {
			// TODO improve this
			spawner, ok := b.activeEntity.(components.UnitSpawnerOwner)
			if !ok {
				return
			}

			unit := units.NewUnit(components.TeamBlue, spawner.GetUnitSpawner().Class, atlasSpriteGetter{})
			b.base.Spawner.SpawnUnit(unit)

			x, y := ebiten.CursorPosition()
			unit.GetWorldSpace().SetInWorld(float64(x), float64(y))
		}
	}
}

func (b BuildingControlSystem) Draw(canvas engine.Sprite) {
}

func (b *BuildingControlSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntitySelected:
		foundEntity, ok := b.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		b.activeEntity = foundEntity.(buildingControlEntity)
	case EntityUnselected:
		// TODO think this through
		b.activeEntity = nil
	}
}

func (b *BuildingControlSystem) Add(entity buildingControlEntity) {
	b.entities.Add(entity)
}

func (b *BuildingControlSystem) Remove(entity engine.Entity) {
	b.entities.Remove(entity)
}
