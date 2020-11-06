package units

import (
	"time"

	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/objects"
)

type Unit struct {
	*engine.BaseEntity
	*components.WorldSpace
	*components.Citizen
	*components.Drawable
	*components.Movable
	*components.Selectable
	*components.Clickable
	*components.Builder
	*components.TimeActions
}

type spriteGetter interface {
	SpriteForUnit(components.Team, components.Class) engine.Sprite
}

func NewUnit(team components.Team, class components.Class, spriteGetter spriteGetter) Unit {
	sprite := spriteGetter.SpriteForUnit(team, class)
	w, h := sprite.Size()
	overlay := objects.NewOverlay(w+20, h+20, engine.PivotBottom)

	var options []components.BuilderOption
	switch class {
	case components.ClassWorker:
		options = []components.BuilderOption{
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
	}

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
		&components.Builder{
			Options: options,
		},
		&components.TimeActions{},
	}

	u.GetWorldSpace().AddChild(u, overlay)
	overlay.GetWorldSpace().Translate(0, 10)

	return u
}
