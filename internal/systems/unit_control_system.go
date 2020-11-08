package systems

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/m110/moonshot-rts/internal/atlas"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/objects"
	"github.com/m110/moonshot-rts/internal/tiles"
)

type unitControlEntity interface {
	engine.Entity
	components.WorldSpaceOwner
	components.MovableOwner
	components.BuilderOwner
	components.TimeActionsOwner
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
	buildingsQueued map[engine.EntityID]components.BuilderOption

	buildIcon    engine.Sprite
	actionButton *objects.PanelButton
	actionsPanel *objects.Panel

	highlightedTile   objects.Object
	tileSelectionMode bool
}

type EntityReachedTarget struct {
	Entity engine.Entity
}

func NewUnitControlSystem(base BaseSystem, tileFinder tileFinder) *UnitControlSystem {
	u := &UnitControlSystem{
		BaseSystem: base,
		tileFinder: tileFinder,
	}

	u.EventBus.Subscribe(PointClicked{}, u)
	u.EventBus.Subscribe(EntitySelected{}, u)
	u.EventBus.Subscribe(EntityUnselected{}, u)
	u.EventBus.Subscribe(EntityReachedTarget{}, u)
	u.EventBus.Subscribe(EntitiesCollided{}, u)
	u.EventBus.Subscribe(EntitiesOutOfCollision{}, u)

	return u
}

func (u *UnitControlSystem) Start() {
	u.buildMode = false
	u.buildingsQueued = map[engine.EntityID]components.BuilderOption{}

	u.buildIcon = atlas.Hammer
	u.buildIcon.Scale(engine.Vector{X: 0.5, Y: 0.5})

	u.highlightedTile = tiles.NewHighlightTile(u.Config.TileMap.TileWidth, u.Config.TileMap.TileHeight)
	u.Spawner.SpawnObject(u.highlightedTile)
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
		u.hideActionButton()
		u.hideActionsPanel()
	case PointClicked:
		for _, e := range u.activeEntities.All() {
			entity := e.(unitControlEntity)
			entity.GetMovable().SetTarget(event.Point)

			_, ok := u.buildingsQueued[entity.ID()]
			if ok {
				// Entity moved after commanded to build - cancel the building
				delete(u.buildingsQueued, entity.ID())
			}

			if u.buildMode {
				u.buildingsQueued[entity.ID()] = u.buildingToBuild
				u.buildMode = false
			}
		}
	case EntityReachedTarget:
		buildingType, ok := u.buildingsQueued[event.Entity.ID()]
		if !ok {
			return
		}
		entity, ok := u.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		// TODO this should be based on tiles
		uce := entity.(unitControlEntity)
		pos := uce.GetWorldSpace().WorldPosition()
		pos.Translate(0, -24)

		building := objects.NewBuilding(pos, buildingType.BuildingType)
		u.Spawner.SpawnBuilding(building)

		delete(u.buildingsQueued, event.Entity.ID())
	case EntitiesCollided:
		entity, ok := u.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

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
		if entity.GetMovable().Target != nil {
			if entity.GetWorldSpace().WorldPosition().Distance(*entity.GetMovable().Target) < 1.0 {
				u.EventBus.Publish(EntityReachedTarget{Entity: entity})
				entity.GetMovable().ClearTarget()
			} else {
				direction := entity.GetMovable().Target.Sub(entity.GetWorldSpace().WorldPosition()).Normalized()
				entity.GetWorldSpace().Translate(direction.Mul(50 * dt).Unpack())
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

	button := objects.NewPanelButton(components.UIColorBrown, u.buildIcon, func() {
		u.hideActionButton()
		u.showActionPanel()
	})

	entity.GetWorldSpace().AddChild(button)

	u.Spawner.SpawnPanelButton(button)
	u.actionButton = &button
}

func (u *UnitControlSystem) hideActionButton() {
	if u.actionButton == nil {
		return
	}

	u.Spawner.RemovePanelButton(*u.actionButton)
	u.actionButton = nil
}

func (u *UnitControlSystem) showActionPanel() {
	entity := u.activeEntities.All()[0].(unitControlEntity)
	options := entity.GetBuilder().Options

	var configs []objects.ButtonConfig
	for i := range options {
		o := options[i]
		sprite := objects.SpriteForBuilding(o.BuildingType)
		sprite.Scale(engine.Vector{X: 0.5, Y: 0.5})

		configs = append(configs, objects.ButtonConfig{
			Sprite: sprite,
			Action: func() {
				u.buildMode = true
				u.buildingToBuild = o
				u.hideActionsPanel()
			},
		})
	}

	panel := objects.NewFourButtonPanel(configs)
	u.Spawner.SpawnPanel(panel)

	entity.GetWorldSpace().AddChild(panel)

	u.actionsPanel = &panel
}

func (u *UnitControlSystem) hideActionsPanel() {
	if u.actionsPanel == nil {
		return
	}

	u.Spawner.RemovePanel(*u.actionsPanel)
	u.actionsPanel = nil
}

func (u *UnitControlSystem) Add(entity unitControlEntity) {
	u.entities.Add(entity)
}

func (u *UnitControlSystem) Remove(entity engine.Entity) {
	u.entities.Remove(entity)
}
