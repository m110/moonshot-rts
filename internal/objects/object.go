package objects

import (
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
)

type TreeType int

const (
	TreeStandard TreeType = iota
	TreePine
)

type Object struct {
	*engine.BaseEntity
	*components.WorldSpace
	*components.Drawable
}

func (o Object) GetWorldSpace() *components.WorldSpace {
	return o.WorldSpace
}

func (o Object) GetDrawable() *components.Drawable {
	return o.Drawable
}

func NewObject(sprite engine.Sprite, layer components.DrawingLayer) Object {
	return Object{
		BaseEntity: engine.NewBaseEntity(),
		WorldSpace: &components.WorldSpace{},
		Drawable: &components.Drawable{
			Sprite: sprite,
			Layer:  layer,
		},
	}
}
