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
	// TODO It seems each system doesn't new a Draw method
	// Possibly, there could be just one system with it (DrawingSystem)
	// Or perhaps just a few - for example, WorldDrawingSystem and UIDrawingSystem could make sense
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

	baseSystem := systems.NewBaseSystem(g.config, g.eventBus, g)
	g.systems = []System{
		systems.NewTilemapSystem(baseSystem),
		systems.NewDrawingSystem(baseSystem),
		systems.NewSelectionSystem(baseSystem),
		systems.NewUISystem(baseSystem),
		systems.NewResourcesSystem(baseSystem),
		systems.NewUnitControlSystem(baseSystem),
		systems.NewBuildingControlSystem(baseSystem),
		systems.NewClickingSystem(baseSystem),
		systems.NewButtonsSystem(baseSystem),
		systems.NewTimeActionsSystem(baseSystem),
		systems.NewProgressBarSystem(baseSystem),
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
		case *systems.ClickingSystem:
			system.Add(building)
		case *systems.TimeActionsSystem:
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
		case *systems.ClickingSystem:
			system.Add(unit)
		case *systems.TimeActionsSystem:
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
		case *systems.ClickingSystem:
			system.Add(panel)
		}
	}

	for _, c := range panel.GetWorldSpace().Children {
		b, ok := c.(objects.PanelButton)
		if ok {
			g.SpawnPanelButton(b)
		}
	}
}

func (g *Game) RemovePanel(panel objects.Panel) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Remove(panel)
		case *systems.ClickingSystem:
			system.Remove(panel)
		}
	}

	for _, c := range panel.GetWorldSpace().Children {
		b, ok := c.(objects.PanelButton)
		if ok {
			g.RemovePanelButton(b)
		}
	}
}

func (g *Game) SpawnPanelButton(button objects.PanelButton) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Add(button)
		case *systems.ClickingSystem:
			system.Add(button)
		case *systems.ButtonsSystem:
			system.Add(button)
		}
	}
}

func (g *Game) RemovePanelButton(button objects.PanelButton) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Remove(button)
		case *systems.ClickingSystem:
			system.Remove(button)
		case *systems.ButtonsSystem:
			system.Remove(button)
		}
	}
}

func (g *Game) SpawnProgressBar(progressBar objects.ProgressBar) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Add(progressBar)
		case *systems.ProgressBarSystem:
			system.Add(progressBar)
		}
	}
}

func (g *Game) RemoveProgressBar(progressBar objects.ProgressBar) {
	for _, s := range g.systems {
		switch system := s.(type) {
		case *systems.DrawingSystem:
			system.Remove(progressBar)
		case *systems.ProgressBarSystem:
			system.Remove(progressBar)
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
