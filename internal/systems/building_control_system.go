package systems

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/objects"
	"github.com/m110/moonshot-rts/internal/units"
)

type buildingControlEntity interface {
	engine.Entity
	components.WorldSpaceOwner
}

type BuildingControlSystem struct {
	base BaseSystem

	entities EntityList

	activeBuilding buildingControlEntity

	buildPanel *objects.Panel

	spawnedUnits []units.Unit
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

func (b *BuildingControlSystem) Update(dt float64) {
	// TODO Rework this (with events?)
	units := b.spawnedUnits
	b.spawnedUnits = nil
	for _, u := range units {
		if u.Target == nil {
			u.Clickable.Enable()
		} else {
			b.spawnedUnits = append(b.spawnedUnits, u)
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

		b.activeBuilding = foundEntity.(buildingControlEntity)
		b.ShowBuildPanel()
	case EntityUnselected:
		if b.activeBuilding != nil && event.Entity.Equals(b.activeBuilding) {
			b.activeBuilding = nil
			b.base.Spawner.RemovePanel(*b.buildPanel)
			b.buildPanel = nil
		}
	}
}

func (b *BuildingControlSystem) ShowBuildPanel() {
	getter := atlasSpriteGetter{}

	spawner, ok := b.activeBuilding.(components.UnitSpawnerOwner)
	if !ok {
		// TODO Add more building types?
		return
	}

	team := components.TeamBlue

	var configs []objects.ButtonConfig
	for i := range spawner.GetUnitSpawner().Classes {
		class := spawner.GetUnitSpawner().Classes[i]
		configs = append(configs, objects.ButtonConfig{
			Sprite: getter.SpriteForUnit(team, class),
			Action: func() { b.spawnUnit(team, class) },
		})
	}
	buildPanel := objects.NewFourButtonPanel(configs)
	b.base.Spawner.SpawnPanel(buildPanel)

	pos := b.activeBuilding.GetWorldSpace().WorldPosition()
	buildPanel.GetWorldSpace().SetInWorld(pos.X, pos.Y)

	b.buildPanel = &buildPanel
}

func (b *BuildingControlSystem) spawnUnit(team components.Team, class components.Class) {
	// Sanity check, this shouldn't happen
	if b.activeBuilding == nil {
		return
	}

	unit := units.NewUnit(team, class, atlasSpriteGetter{})
	b.base.Spawner.SpawnUnit(unit)

	pos := b.activeBuilding.GetWorldSpace().WorldPosition()
	unit.GetWorldSpace().SetInWorld(pos.X, pos.Y)

	// TODO This should be based on tiles, not absolute position
	target := pos
	target.Translate(
		float64(engine.RandomRange(-32, 32)),
		float64(engine.RandomRange(12, 32)),
	)
	unit.Clickable.Disable()
	unit.GetMovable().SetTarget(target)
	b.spawnedUnits = append(b.spawnedUnits, unit)
}

func (b *BuildingControlSystem) Add(entity buildingControlEntity) {
	b.entities.Add(entity)
}

func (b *BuildingControlSystem) Remove(entity engine.Entity) {
	b.entities.Remove(entity)
}
