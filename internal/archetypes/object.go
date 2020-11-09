package archetypes

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
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
