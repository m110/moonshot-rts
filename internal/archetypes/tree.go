package archetypes

import (
	"github.com/m110/moonshot-rts/internal/assets/sprites"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type Tree struct {
	Object
}

func NewTree(treeType TreeType) Tree {
	var sprite engine.Sprite
	size := engine.RandomRange(0, 1)

	switch treeType {
	case TreeStandard:
		if size == 0 {
			sprite = sprites.TreeBig
		} else {
			sprite = sprites.TreeSmall
		}
	case TreePine:
		if size == 0 {
			sprite = sprites.PineBig
		} else {
			sprite = sprites.PineSmall
		}
	}

	return Tree{
		Object: NewObject(sprite, components.LayerObjects),
	}
}
