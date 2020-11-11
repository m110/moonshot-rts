package systems

import (
	"fmt"

	"github.com/m110/moonshot-rts/internal/archetypes"

	"github.com/m110/moonshot-rts/internal/archetypes/tiles"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
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
	case archetypes.Building:
		s.spawnBuilding(entity)
	case archetypes.Unit:
		s.spawnUnit(entity)
	case archetypes.Worker:
		s.spawnWorker(entity)
	case archetypes.Panel:
		s.spawnPanel(entity)
	case archetypes.PanelButton:
		s.spawnPanelButton(entity)
	case archetypes.ProgressBar:
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
	case archetypes.Panel:
		s.destroyPanel(entity)
	case archetypes.PanelButton:
		s.destroyPanelButton(entity)
	case archetypes.ProgressBar:
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

func (s Spawner) spawnBuilding(building archetypes.Building) {
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

func (s Spawner) spawnUnit(unit archetypes.Unit) {
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
		case *CollisionSystem:
			system.Add(unit)
		case *AreaOccupySystem:
			system.Add(unit)
		}
	}
}

func (s Spawner) spawnWorker(worker archetypes.Worker) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(worker)
		case *SelectionSystem:
			system.Add(worker)
		case *UnitControlSystem:
			system.Add(worker)
		case *ClickingSystem:
			system.Add(worker)
		case *CollisionSystem:
			system.Add(worker)
		case *AreaOccupySystem:
			system.Add(worker)
		case *ResourcesSystem:
			system.Add(worker)
		}
	}
}

func (s Spawner) spawnPanel(panel archetypes.Panel) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(panel)
		case *ClickingSystem:
			system.Add(panel)
		}
	}
}

func (s Spawner) destroyPanel(panel archetypes.Panel) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Remove(panel)
		case *ClickingSystem:
			system.Remove(panel)
		}
	}
}

func (s Spawner) spawnPanelButton(button archetypes.PanelButton) {
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

func (s Spawner) destroyPanelButton(button archetypes.PanelButton) {
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

func (s Spawner) spawnProgressBar(progressBar archetypes.ProgressBar) {
	for _, sys := range s.systemsProvider.Systems() {
		switch system := sys.(type) {
		case *DrawingSystem:
			system.Add(progressBar)
		case *ProgressBarSystem:
			system.Add(progressBar)
		}
	}
}

func (s Spawner) destroyProgressBar(progressBar archetypes.ProgressBar) {
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
