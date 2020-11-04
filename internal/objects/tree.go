package objects

import (
	"github.com/m110/rts/internal/atlas"
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
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
			sprite = atlas.TreeBig
		} else {
			sprite = atlas.TreeSmall
		}
	case TreePine:
		if size == 0 {
			sprite = atlas.PineBig
		} else {
			sprite = atlas.PineSmall
		}
	}

	return Tree{
		Object: NewObject(sprite, components.LayerObjects),
	}
}
