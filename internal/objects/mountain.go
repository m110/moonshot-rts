package objects

import (
	"github.com/m110/moonshot-rts/internal/atlas"
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
			sprite = atlas.StoneSmall
		case 1:
			sprite = atlas.StoneBig
		case 2:
			sprite = atlas.StoneTwo
		case 3:
			sprite = atlas.StoneThree
		}
	case MountainGold:
		sprite = atlas.GoldThree
	case MountainIron:
		sprite = atlas.IronThree
	}

	return Mountain{
		Object: NewObject(sprite, components.LayerObjects),
	}
}
