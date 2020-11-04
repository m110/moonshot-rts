package objects

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
	"golang.org/x/image/colornames"
)

type Overlay struct {
	*engine.BaseEntity
	*components.WorldSpace
	*components.Drawable
	*components.Size
}

func NewOverlay(width int, height int, pivotType engine.PivotType) Overlay {
	o := Overlay{
		BaseEntity: engine.NewBaseEntity(),
		WorldSpace: &components.WorldSpace{},
		Size: &components.Size{
			Width:  width,
			Height: height,
		},
	}

	o.Drawable = &components.Drawable{
		Sprite:   NewRectangleSprite(o, pivotType),
		Layer:    components.LayerUI,
		Disabled: true,
	}

	return o
}

func (o Overlay) GetSize() *components.Size {
	return o.Size
}

func NewRectangleSprite(owner components.SizeOwner, pivotType engine.PivotType) engine.Sprite {
	width := float64(owner.GetSize().Width)
	height := float64(owner.GetSize().Height)

	c := colornames.White
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

func (o Overlay) GetWorldSpace() *components.WorldSpace {
	return o.WorldSpace
}

func (o Overlay) GetDrawable() *components.Drawable {
	return o.Drawable
}
