package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/systems"
)

type Game struct {
	config systems.Config

	eventBus *engine.EventBus

	systems []systems.System
	drawers []systems.Drawer
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

	tilemapSystem := systems.NewTilemapSystem(baseSystem)
	g.systems = []systems.System{
		tilemapSystem,
		systems.NewDrawingSystem(baseSystem),
		systems.NewSelectionSystem(baseSystem),
		systems.NewUISystem(baseSystem),
		systems.NewResourcesSystem(baseSystem),
		systems.NewUnitControlSystem(baseSystem, tilemapSystem),
		systems.NewBuildingControlSystem(baseSystem),
		systems.NewClickingSystem(baseSystem),
		systems.NewButtonsSystem(baseSystem),
		systems.NewTimeActionsSystem(baseSystem),
		systems.NewProgressBarSystem(baseSystem),
		systems.NewCollisionSystem(baseSystem),
	}

	for _, s := range g.systems {
		d, ok := s.(systems.Drawer)
		if ok {
			g.drawers = append(g.drawers, d)
		}

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
	for _, s := range g.drawers {
		s.Draw(screenSprite)
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g Game) Systems() []systems.System {
	return g.systems
}
