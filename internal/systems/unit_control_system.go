package systems

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/m110/moonshot-rts/internal/archetypes"
	"github.com/m110/moonshot-rts/internal/archetypes/tiles"
	"github.com/m110/moonshot-rts/internal/assets/sprites"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type unitControlEntity interface {
	engine.Entity
	components.WorldSpaceOwner
	components.MovableOwner
	components.BuilderOwner
	components.TimeActionsOwner
	components.ColliderOwner
}

type tileFinder interface {
	TileAtPosition(position engine.Vector) (tiles.Tile, bool)
}

type UnitControlSystem struct {
	BaseSystem
	tileFinder tileFinder

	entities EntityList

	activeEntities EntityList

	buildMode       bool
	buildingToBuild components.BuilderOption
	buildingsQueued map[engine.EntityID]queuedBuilding

	buildIcon    engine.Sprite
	actionButton *archetypes.PanelButton
	actionsPanel *archetypes.Panel

	highlightedTile   archetypes.Object
	tileSelectionMode bool
}

type queuedBuilding struct {
	DestinationTile tiles.Tile
	Option          components.BuilderOption
}

type EntityReachedTarget struct {
	Entity engine.Entity
}

type EntityMoved struct {
	Entity engine.Entity
}

func NewUnitControlSystem(base BaseSystem, tileFinder tileFinder) *UnitControlSystem {
	u := &UnitControlSystem{
		BaseSystem:      base,
		tileFinder:      tileFinder,
		buildingsQueued: map[engine.EntityID]queuedBuilding{},
	}

	u.EventBus.Subscribe(PointClicked{}, u)
	u.EventBus.Subscribe(EntitySelected{}, u)
	u.EventBus.Subscribe(EntityUnselected{}, u)
	u.EventBus.Subscribe(EntitiesCollided{}, u)
	u.EventBus.Subscribe(EntitiesOutOfCollision{}, u)

	return u
}

func (u *UnitControlSystem) Start() {
	u.buildMode = false

	u.buildIcon = sprites.Hammer
	u.buildIcon.Scale(engine.Vector{X: 0.5, Y: 0.5})

	u.highlightedTile = tiles.NewHighlightTile(u.Config.TileMap.TileWidth, u.Config.TileMap.TileHeight)
	u.Spawner.Spawn(u.highlightedTile)
}

func (u UnitControlSystem) Update(dt float64) {
	u.moveEntities(dt)
	u.updateHighlightedTile()
}

func (u *UnitControlSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntitySelected:
		foundEntity, ok := u.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		u.tileSelectionMode = true

		u.activeEntities.Add(foundEntity)
		if len(u.activeEntities.All()) > 1 {
			u.hideActionButton()
		} else {
			u.showActionButton()
		}
	case EntityUnselected:
		u.activeEntities.Remove(event.Entity)
		u.tileSelectionMode = false
		u.buildMode = false
		u.hideActionButton()
		u.hideActionsPanel()
	case PointClicked:
		for _, e := range u.activeEntities.All() {
			entity := e.(unitControlEntity)

			_, ok := u.buildingsQueued[entity.ID()]
			if ok {
				// Entity moved after commanded to build - cancel the building
				delete(u.buildingsQueued, entity.ID())
			}

			if u.buildMode {
				tile, ok := u.tileFinder.TileAtPosition(event.Point)
				if !ok {
					fmt.Println("Tile not found at position", event.Point)
					return
				}

				u.buildMode = false

				if !canBuildOnTile(tile) {
					return
				}

				u.buildingsQueued[entity.ID()] = queuedBuilding{
					DestinationTile: tile,
					Option:          u.buildingToBuild,
				}

				if entity.GetCollider().HasCollision(tile) {
					u.attemptBuildOnCollision(entity, tile)
					return
				}
			}

			entity.GetMovable().SetTarget(event.Point)
		}
	case EntitiesCollided:
		entity, ok := u.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		u.attemptBuildOnCollision(entity, event.Other)
		fmt.Println(entity, "collides with", event.Other)
	case EntitiesOutOfCollision:
		entity, ok := u.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}
		fmt.Println(entity, "out of collision with", event.Other)
	}
}

func (u *UnitControlSystem) moveEntities(dt float64) {
	for _, e := range u.entities.All() {
		entity := e.(unitControlEntity)
		movable := entity.GetMovable()
		if movable.Disabled || movable.Target != nil {
			if entity.GetWorldSpace().WorldPosition().Distance(*movable.Target) < 1.0 {
				u.EventBus.Publish(EntityReachedTarget{Entity: entity})
				movable.ClearTarget()
			} else {
				direction := movable.Target.Sub(entity.GetWorldSpace().WorldPosition()).Normalized()
				entity.GetWorldSpace().Translate(direction.Mul(50 * dt).Unpack())
				u.EventBus.Publish(EntityMoved{Entity: entity})
			}
		}
	}
}

func (u *UnitControlSystem) updateHighlightedTile() {
	u.highlightedTile.GetDrawable().Disable()
	if u.tileSelectionMode {
		x, y := ebiten.CursorPosition()
		v := engine.Vector{X: float64(x), Y: float64(y)}
		tile, ok := u.tileFinder.TileAtPosition(v)
		if ok {
			u.highlightedTile.GetDrawable().Enable()
			u.highlightedTile.GetWorldSpace().SetInWorld(
				tile.GetWorldSpace().WorldPosition().X,
				tile.GetWorldSpace().WorldPosition().Y,
			)
		}
	}
}

func (u *UnitControlSystem) showActionButton() {
	entity := u.activeEntities.All()[0].(unitControlEntity)
	if len(entity.GetBuilder().Options) == 0 {
		return
	}

	button := archetypes.NewPanelButton(components.UIColorBrown, u.buildIcon, func() {
		u.hideActionButton()
		u.showActionPanel()
	})

	entity.GetWorldSpace().AddChild(button)

	u.Spawner.Spawn(button)
	u.actionButton = &button
}

func (u *UnitControlSystem) hideActionButton() {
	if u.actionButton == nil {
		return
	}

	u.Spawner.Destroy(*u.actionButton)
	u.actionButton = nil
}

func (u *UnitControlSystem) showActionPanel() {
	entity := u.activeEntities.All()[0].(unitControlEntity)
	options := entity.GetBuilder().Options

	var configs []archetypes.ButtonConfig
	for i := range options {
		o := options[i]
		sprite := archetypes.SpriteForBuilding(o.BuildingType)
		sprite.Scale(engine.Vector{X: 0.5, Y: 0.5})

		configs = append(configs, archetypes.ButtonConfig{
			Sprite: sprite,
			Action: func() {
				u.buildMode = true
				u.buildingToBuild = o
				u.hideActionsPanel()
			},
		})
	}

	panel := archetypes.NewFourButtonPanel(configs)
	u.Spawner.Spawn(panel)

	entity.GetWorldSpace().AddChild(panel)

	u.actionsPanel = &panel
}

func (u *UnitControlSystem) hideActionsPanel() {
	if u.actionsPanel == nil {
		return
	}

	u.Spawner.Destroy(*u.actionsPanel)
	u.actionsPanel = nil
}

func (u *UnitControlSystem) attemptBuildOnCollision(entity engine.Entity, other engine.Entity) {
	queued, ok := u.buildingsQueued[entity.ID()]
	if !ok {
		return
	}

	if !queued.DestinationTile.Equals(other) {
		return
	}

	uce := entity.(unitControlEntity)
	uce.GetMovable().ClearTarget()

	buildingPos := engine.Vector{
		X: float64(u.Config.TileMap.TileWidth) / 2.0,
		Y: float64(u.Config.TileMap.TileHeight),
	}

	building := archetypes.NewBuilding(buildingPos, queued.Option.BuildingType)
	queued.DestinationTile.GetWorldSpace().AddChild(building)
	u.Spawner.Spawn(building)

	delete(u.buildingsQueued, entity.ID())
}

func canBuildOnTile(tile tiles.Tile) bool {
	// TODO ?
	return len(tile.GetWorldSpace().Children) == 0
}

func (u *UnitControlSystem) Add(entity unitControlEntity) {
	u.entities.Add(entity)
}

func (u *UnitControlSystem) Remove(entity engine.Entity) {
	u.entities.Remove(entity)
}
