package archetypes

import (
	"time"

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
	overlay := NewOverlay(w+20, h+20, engine.PivotBottom)

	spriteBounds := components.BoundsFromSprite(sprite)

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
			Bounds:    spriteBounds,
			ByOverlay: true,
		},
		&components.Collider{
			Bounds: engine.Rect{
				Position: engine.Vector{
					X: spriteBounds.Width / 2,
					Y: spriteBounds.Height - 10,
				},
				Width:  10,
				Height: 10,
			},
			Layer: components.CollisionLayerUnits,
		},
		&components.AreaOccupant{},
	}

	u.GetWorldSpace().AddChild(overlay)
	overlay.GetWorldSpace().Translate(0, 10)

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
