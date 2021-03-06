package archetypes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type Overlay struct {
	*engine.BaseEntity
	*components.WorldSpace
	*components.Drawable
	*components.Size
}

func NewOverlay(width int, height int, pivotType engine.PivotType, c color.RGBA) Overlay {
	o := Overlay{
		BaseEntity: engine.NewBaseEntity(),
		WorldSpace: &components.WorldSpace{},
		Size: &components.Size{
			Width:  width,
			Height: height,
		},
	}

	o.Drawable = &components.Drawable{
		Sprite:   NewRectangleSprite(o, pivotType, c),
		Layer:    components.LayerUI,
		Disabled: true,
	}

	return o
}

func NewRectangleSprite(owner components.SizeOwner, pivotType engine.PivotType, c color.RGBA) engine.Sprite {
	width := float64(owner.GetSize().Width)
	height := float64(owner.GetSize().Height)

	c.A = 175
	sprite := engine.NewBlankSprite(int(width), int(height))
	sprite.SetPivot(engine.NewPivotForSprite(sprite, pivotType))
	lineSize := 3.0
	ebitenutil.DrawRect(sprite.Image(), 0, 0, width, lineSize, c)
	ebitenutil.DrawRect(sprite.Image(), 0, 0, lineSize, height, c)
	ebitenutil.DrawRect(sprite.Image(), width-lineSize, 0, lineSize, height, c)
	ebitenutil.DrawRect(sprite.Image(), 0, height-lineSize, width, lineSize, c)

	return sprite
}
