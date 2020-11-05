package units

import (
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
}

type spriteGetter interface {
	SpriteForUnit(components.Team, components.Class) engine.Sprite
}

func NewUnit(team components.Team, class components.Class, spriteGetter spriteGetter) Unit {
	sprite := spriteGetter.SpriteForUnit(team, class)
	w, h := sprite.Size()
	overlay := objects.NewOverlay(w+20, h+20, engine.PivotBottom)

	var buildings []components.BuildingType
	switch class {
	case components.ClassWorker:
		buildings = []components.BuildingType{
			components.BuildingBarracks,
			components.BuildingForge,
			components.BuildingChapel,
			components.BuildingTower,
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
			Buildings: buildings,
		},
	}

	u.GetWorldSpace().AddChild(u, overlay)
	overlay.GetWorldSpace().Translate(0, 10)

	return u
}
