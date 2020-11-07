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
	components.TimeActionsOwner
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

	for _, e := range b.entities.All() {
		entity, ok := e.(components.UnitSpawnerOwner)
		if !ok {
			continue
		}

		// TODO this nil-check doesn't look good
		timer := entity.GetUnitSpawner().Timer
		if timer != nil {
			timer.Update(dt)
			for _, c := range e.(buildingControlEntity).GetWorldSpace().Children {
				// TODO casting to concrete struct is a hack, there should be a better way to do this
				p, ok := c.(objects.ProgressBar)
				if ok {
					p.SetProgress(timer.PercentDone())
				}
			}
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

	// TODO get this from a component
	team := components.TeamBlue

	var configs []objects.ButtonConfig
	for i := range spawner.GetUnitSpawner().Options {
		o := spawner.GetUnitSpawner().Options[i]

		configs = append(configs, objects.ButtonConfig{
			Sprite: getter.SpriteForUnit(team, o.Class),
			Action: func() { b.startSpawningUnit(o) },
		})
	}
	buildPanel := objects.NewFourButtonPanel(configs)
	b.base.Spawner.SpawnPanel(buildPanel)

	pos := b.activeBuilding.GetWorldSpace().WorldPosition()
	buildPanel.GetWorldSpace().SetInWorld(pos.X, pos.Y)

	b.buildPanel = &buildPanel
}

func (b *BuildingControlSystem) startSpawningUnit(option components.UnitSpawnerOption) {
	// Sanity check, this shouldn't happen
	if b.activeBuilding == nil {
		return
	}

	entity := b.activeBuilding.(components.UnitSpawnerOwner)
	entity.GetUnitSpawner().AddToQueue(option)
	b.addSpawnUnitTimer(b.activeBuilding)
}

func (b *BuildingControlSystem) addSpawnUnitTimer(entity buildingControlEntity) {
	unitSpawner := entity.(components.UnitSpawnerOwner).GetUnitSpawner()

	if unitSpawner.Timer == nil {
		option, ok := unitSpawner.PopFromQueue()
		if !ok {
			return
		}

		b.showProgressBar(entity)

		timer := engine.NewCountdownTimer(option.SpawnTime, func() {
			// TODO get this from a component
			team := components.TeamBlue
			pos := entity.GetWorldSpace().WorldPosition()

			b.spawnUnit(pos, team, option.Class)

			unitSpawner.Timer = nil

			for _, c := range entity.GetWorldSpace().Children {
				// TODO casting to concrete struct is a hack, there should be a better way to do this
				p, ok := c.(objects.ProgressBar)
				if ok {
					b.base.Spawner.RemoveProgressBar(p)
					// TODO A RemoveChild is missing here. Not trivial for now, and despawning should work fine
				}
			}

			b.addSpawnUnitTimer(entity)
		})
		unitSpawner.Timer = timer
	}
}

func (b *BuildingControlSystem) showProgressBar(entity buildingControlEntity) {
	progressBar := objects.NewHorizontalProgressBar()
	b.base.Spawner.SpawnProgressBar(progressBar)
	entity.GetWorldSpace().AddChild(progressBar)
	// TODO better position
	progressBar.GetWorldSpace().Translate(0, -30)
}

func (b *BuildingControlSystem) spawnUnit(spawnPosition engine.Vector, team components.Team, class components.Class) {
	unit := units.NewUnit(team, class, atlasSpriteGetter{})
	b.base.Spawner.SpawnUnit(unit)

	unit.GetWorldSpace().SetInWorld(spawnPosition.X, spawnPosition.Y)

	// TODO This should be based on tiles, not absolute position
	target := spawnPosition
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
