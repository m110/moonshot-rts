package systems

import (
	"fmt"

	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/objects"
	"github.com/m110/moonshot-rts/internal/tiles"
	"github.com/m110/moonshot-rts/internal/units"
)

type Spawner struct {
	systemsProvider systemsProvider
}

func NewSpawner(systemsProvider systemsProvider) Spawner {
	return Spawner{
		systemsProvider: systemsProvider,
	}
}

func (s Spawner) Spawn(e engine.Entity) {
	switch entity := e.(type) {
	case tiles.Tile:
		s.spawnTile(entity)
	case objects.Building:
		s.spawnBuilding(entity)
	case units.Unit:
		s.spawnUnit(entity)
	case objects.Panel:
		s.spawnPanel(entity)
	case objects.PanelButton:
		s.spawnPanelButton(entity)
	case objects.ProgressBar:
		s.spawnProgressBar(entity)
	default:
		// Fallback to drawing system for the simplest entities
		if drawingEntity, ok := e.(DrawingEntity); ok {
			s.spawnDrawingEntity(drawingEntity)
		} else {
			panic(fmt.Sprintf("No suitable function found to spawn entity %#v", entity))
		}
	}

	// Spawn all children
	if worldSpaceOwner, ok := e.(components.WorldSpaceOwner); ok {
		for _, child := range worldSpaceOwner.GetWorldSpace().Children {
			s.Spawn(child.(engine.Entity))
		}
	}
}

func (s Spawner) Destroy(e engine.Entity) {
	switch entity := e.(type) {
	case objects.Panel:
		s.destroyPanel(entity)
	case objects.PanelButton:
		s.destroyPanelButton(entity)
	case objects.ProgressBar:
		s.destroyProgressBar(entity)
	default:
		panic(fmt.Sprintf("No suitable function found to destroy entity %#v", entity))
	}

	// Destroy all children
	if worldSpaceOwner, ok := e.(components.WorldSpaceOwner); ok {
		for _, child := range worldSpaceOwner.GetWorldSpace().Children {
			s.Destroy(child.(engine.Entity))
		}
	}
}

func (s Spawner) spawnTile(tile tiles.Tile) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(tile)
		case *CollisionSystem:
			system.Add(tile)
		}
	}
}

func (s Spawner) spawnBuilding(building objects.Building) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(building)
		case *SelectionSystem:
			system.Add(building)
		case *BuildingControlSystem:
			system.Add(building)
		case *ClickingSystem:
			system.Add(building)
		case *TimeActionsSystem:
			system.Add(building)
		case *CollisionSystem:
			system.Add(building)
		}
	}
}

func (s Spawner) spawnUnit(unit units.Unit) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(unit)
		case *SelectionSystem:
			system.Add(unit)
		case *UnitControlSystem:
			system.Add(unit)
		case *ClickingSystem:
			system.Add(unit)
		case *TimeActionsSystem:
			system.Add(unit)
		case *CollisionSystem:
			system.Add(unit)
		}
	}
}

func (s Spawner) spawnPanel(panel objects.Panel) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(panel)
		case *ClickingSystem:
			system.Add(panel)
		}
	}
}

func (s Spawner) destroyPanel(panel objects.Panel) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Remove(panel)
		case *ClickingSystem:
			system.Remove(panel)
		}
	}
}

func (s Spawner) spawnPanelButton(button objects.PanelButton) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(button)
		case *ClickingSystem:
			system.Add(button)
		case *ButtonsSystem:
			system.Add(button)
		}
	}
}

func (s Spawner) destroyPanelButton(button objects.PanelButton) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Remove(button)
		case *ClickingSystem:
			system.Remove(button)
		case *ButtonsSystem:
			system.Remove(button)
		}
	}
}

func (s Spawner) spawnProgressBar(progressBar objects.ProgressBar) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(progressBar)
		case *ProgressBarSystem:
			system.Add(progressBar)
		}
	}
}

func (s Spawner) destroyProgressBar(progressBar objects.ProgressBar) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Remove(progressBar)
		case *ProgressBarSystem:
			system.Remove(progressBar)
		}
	}
}

func (s Spawner) spawnDrawingEntity(entity DrawingEntity) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(entity)
		}
	}
}
