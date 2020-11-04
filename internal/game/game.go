package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/objects"
	"github.com/m110/moonshot-rts/internal/systems"
	"github.com/m110/moonshot-rts/internal/tiles"
	"github.com/m110/moonshot-rts/internal/units"
)

type Game struct {
	config systems.Config

	eventBus *engine.EventBus
	systems  []System
}

type System interface {
	Start()
	Update(dt float64)
	Draw(canvas engine.Sprite)
	Remove(entity engine.Entity)
}

func NewGame(config systems.Config) *Game {
	g := &Game{
		config: config,
	}
	g.Start()

	return g
}

// Start starts a new game. It can also work as a restart, so all important initialisation
// should happen here, instead of NewGame.
// Why not initialize systems in their constructors? For example, the subscribers can be set up
// in the constructors, and when the Start() is executing, you know all systems are in place.
// So it's possible to use spawner in the Start() methods.
func (g *Game) Start() {
	g.eventBus = engine.NewEventBus()

	g.systems = []System{
		systems.NewTilemapSystem(g.config, g.eventBus, g),
		systems.NewDrawingSystem(g.config, g.eventBus, g),
		systems.NewSelectionSystem(g.config, g.eventBus, g),
		systems.NewUISystem(g.config, g.eventBus, g),
		systems.NewResourcesSystem(g.config, g.eventBus, g),
		systems.NewUnitControlSystem(g.config, g.eventBus, g),
		systems.NewBuildingControlSystem(g.config, g.eventBus, g),
	}

	for _, s := range g.systems {
		s.Start()
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Start()
		return nil
	}

	dt := 1.0 / float64(ebiten.MaxTPS())
	for _, s := range g.systems {
		s.Update(dt)
	}

	g.eventBus.Flush()

	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	screenSprite := engine.NewSpriteFromImage(screen)
	for _, s := range g.systems {
		s.Draw(screenSprite)
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) SpawnTile(tile tiles.Tile) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Add(tile)

			for _, c := range tile.GetWorldSpace().Children {
				d, ok := c.(systems.DrawingEntity)
				if ok {
					system.Add(d)
				}
			}
		}
	}

	for _, c := range tile.GetWorldSpace().Children {
		// TODO this seems like a hack
		building, ok := c.(objects.Building)
		if ok {
			g.SpawnBuilding(building)
		}
	}
}

func (g *Game) SpawnBuilding(building objects.Building) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Add(building)
			for _, c := range building.GetWorldSpace().Children {
				d, ok := c.(systems.DrawingEntity)
				if ok {
					system.Add(d)
				}
			}
		case *systems.SelectionSystem:
			system.Add(building)
		case *systems.BuildingControlSystem:
			system.Add(building)
		}
	}
}

func (g *Game) SpawnUnit(unit units.Unit) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Add(unit)
			for _, c := range unit.GetWorldSpace().Children {
				d, ok := c.(systems.DrawingEntity)
				if ok {
					system.Add(d)
				}
			}
		case *systems.SelectionSystem:
			system.Add(unit)
		case *systems.UnitControlSystem:
			system.Add(unit)
		}
	}
}

func (g *Game) SpawnObject(object objects.Object) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Add(object)
			for _, c := range object.GetWorldSpace().Children {
				d, ok := c.(systems.DrawingEntity)
				if ok {
					system.Add(d)
				}
			}
		}
	}
}

func (g *Game) SpawnPanel(panel objects.Panel) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Add(panel)
			for _, c := range panel.GetWorldSpace().Children {
				d, ok := c.(systems.DrawingEntity)
				if ok {
					system.Add(d)
				}
			}
		}
	}
}

func (g *Game) RemovePanel(panel objects.Panel) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Remove(panel)
			for _, c := range panel.GetWorldSpace().Children {
				d, ok := c.(systems.DrawingEntity)
				if ok {
					system.Remove(d)
				}
			}
		}
	}
}

func (g *Game) SpawnDrawingEntity(entity systems.DrawingEntity) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Add(entity)
		}
	}
}
