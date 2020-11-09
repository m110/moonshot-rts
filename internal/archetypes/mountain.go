package archetypes

import (
	"github.com/m110/moonshot-rts/internal/assets/sprites"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type MountainType int

const (
	MountainStone MountainType = iota
	MountainGold
	MountainIron
)

type Mountain struct {
	Object
}

func NewMountain(mountainType MountainType) Mountain {
	var sprite engine.Sprite

	switch mountainType {
	case MountainStone:
		size := engine.RandomRange(0, 3)
		switch size {
		case 0:
			sprite = sprites.StoneSmall
		case 1:
			sprite = sprites.StoneBig
		case 2:
			sprite = sprites.StoneTwo
		case 3:
			sprite = sprites.StoneThree
		}
	case MountainGold:
		sprite = sprites.GoldThree
	case MountainIron:
		sprite = sprites.IronThree
	}

	return Mountain{
		Object: NewObject(sprite, components.LayerObjects),
	}
}
