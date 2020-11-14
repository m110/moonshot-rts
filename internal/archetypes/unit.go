package archetypes

import (
	"time"

	"golang.org/x/image/colornames"

	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type Unit struct {
	*engine.BaseEntity
	*components.WorldSpace
	*components.Citizen
	*components.Drawable
	*components.Movable
	*components.Selectable
	*components.Clickable
	*components.Collider
	*components.AreaOccupant
}

type Worker struct {
	Unit
	*components.Builder
	*components.ResourcesCollector
}

type spriteGetter interface {
	SpriteForUnit(components.Team, components.Class) engine.Sprite
}

func NewUnit(team components.Team, class components.Class, spriteGetter spriteGetter) Unit {
	sprite := spriteGetter.SpriteForUnit(team, class)
	w, h := sprite.Size()
	overlay := NewOverlay(w+20, h+20, engine.PivotBottom, colornames.White)

	colliderBounds := engine.Rect{
		Position: engine.Vector{
			X: -5,
			Y: -5,
		},
		Width:  10,
		Height: 5,
	}
	colliderOverlay := NewOverlay(int(colliderBounds.Width), int(colliderBounds.Height), engine.PivotTopLeft, colornames.Red)

	u := Unit{
		engine.NewBaseEntity(),
		&components.WorldSpace{},
		&components.Citizen{
			Team:  team,
			Class: class,
		},
		&components.Drawable{
			Sprite: sprite,
			Layer:  components.LayerUnits,
		},
		&components.Movable{},
		&components.Selectable{
			Overlay: overlay,
		},
		&components.Clickable{
			Bounds:    components.BoundsFromSprite(sprite),
			ByOverlay: true,
		},
		&components.Collider{
			Bounds:  colliderBounds,
			Layer:   components.CollisionLayerUnits,
			Overlay: colliderOverlay,
		},
		&components.AreaOccupant{},
	}

	u.GetWorldSpace().AddChild(overlay)
	overlay.GetWorldSpace().Translate(0, 10)

	u.GetWorldSpace().AddChild(colliderOverlay)
	colliderOverlay.GetWorldSpace().Translate(colliderBounds.Position.X, colliderBounds.Position.Y)

	return u
}

func NewWorker(team components.Team, spriteGetter spriteGetter) Worker {
	unit := NewUnit(team, components.ClassWorker, spriteGetter)

	options := []components.BuilderOption{
		{
			BuildingType: components.BuildingBarracks,
			SpawnTime:    time.Second * 10,
		},
		{
			BuildingType: components.BuildingForge,
			SpawnTime:    time.Second * 5,
		},
		{
			BuildingType: components.BuildingChapel,
			SpawnTime:    time.Second * 15,
		},
		{
			BuildingType: components.BuildingTower,
			SpawnTime:    time.Second * 20,
		},
	}

	return Worker{
		Unit: unit,
		Builder: &components.Builder{
			Options: options,
		},
		ResourcesCollector: &components.ResourcesCollector{},
	}
}
