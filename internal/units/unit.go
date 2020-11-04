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
	*components.BoxBoundary
}

func (u Unit) GetWorldSpace() *components.WorldSpace {
	return u.WorldSpace
}

func (u Unit) GetDrawable() *components.Drawable {
	return u.Drawable
}

func (u Unit) GetSelectable() *components.Selectable {
	return u.Selectable
}

func (u Unit) GetMovable() *components.Movable {
	return u.Movable
}

func (u Unit) GetBoxBoundary() *components.BoxBoundary {
	return u.BoxBoundary
}

type spriteGetter interface {
	SpriteForUnit(components.Team, components.Class) engine.Sprite
}

func NewUnit(team components.Team, class components.Class, spriteGetter spriteGetter) Unit {
	sprite := spriteGetter.SpriteForUnit(team, class)
	w, h := sprite.Size()
	overlay := objects.NewOverlay(w+20, h+20, engine.PivotBottom)

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
			GroupSelectable: true,
			Overlay:         overlay,
		},
		components.BoxBoundaryFromSprite(sprite),
	}

	u.GetWorldSpace().AddChild(u, overlay)
	overlay.GetWorldSpace().Translate(0, 10)

	return u
}
