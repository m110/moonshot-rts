package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/objects"
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
		if b.activeBuilding != nil {
			/*
				// TODO improve this
				spawner, ok := b.activeBuilding.(components.UnitSpawnerOwner)
				if !ok {
					return
				}

				unit := units.NewUnit(components.TeamBlue, spawner.GetUnitSpawner().Class, atlasSpriteGetter{})
				b.base.Spawner.SpawnUnit(unit)

				x, y := ebiten.CursorPosition()
				unit.GetWorldSpace().SetInWorld(float64(x), float64(y))

			*/
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
	sprites := []engine.Sprite{
		getter.SpriteForUnit(components.TeamBlue, components.ClassWorker),
		getter.SpriteForUnit(components.TeamBlue, components.ClassWarrior),
		getter.SpriteForUnit(components.TeamBlue, components.ClassKnight),
		getter.SpriteForUnit(components.TeamBlue, components.ClassPriest),
	}
	buildPanel := objects.NewPanel(sprites)
	b.base.Spawner.SpawnPanel(buildPanel)

	pos := b.activeBuilding.GetWorldSpace().WorldPosition()
	buildPanel.GetWorldSpace().SetInWorld(pos.X, pos.Y)

	// TODO register the button's entity ID as callback for spawning a unit
	// Buttons need a clickable component which triggers this action

	b.buildPanel = &buildPanel
}

func (b *BuildingControlSystem) Add(entity buildingControlEntity) {
	b.entities.Add(entity)
}

func (b *BuildingControlSystem) Remove(entity engine.Entity) {
	b.entities.Remove(entity)
}
