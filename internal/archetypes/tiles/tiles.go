package tiles

import (
	"math/rand"

	"github.com/m110/moonshot-rts/internal/archetypes"
	"github.com/m110/moonshot-rts/internal/assets/sprites"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type GroundType int

const (
	GroundGrass GroundType = iota
	GroundSand
	GroundSea
)

func NewGroundTile(groundType GroundType) Tile {
	var sprite engine.Sprite
	switch groundType {
	case GroundGrass:
		sprite = sprites.Grass1
	case GroundSand:
		sprite = sprites.Sand1
	case GroundSea:
		sprite = sprites.Water1
	}

	return Tile{
		BaseEntity: engine.NewBaseEntity(),
		WorldSpace: &components.WorldSpace{},
		Drawable: &components.Drawable{
			Sprite: sprite,
			Layer:  components.LayerGround,
		},
		Clickable: &components.Clickable{
			Bounds: components.BoundsFromSprite(sprite),
		},
		Collider: &components.Collider{
			Bounds: components.BoundsFromSprite(sprite),
			Layer:  components.CollisionLayerGround,
		},
		ResourcesSource: &components.ResourcesSource{},
		Area:            &components.Area{},
	}
}

type ForestType int

const (
	ForestStandard ForestType = iota
	ForestPine
	ForestMixed
)

func NewForestTile(groundType GroundType, forestType ForestType) Tile {
	t := NewGroundTile(groundType)

	treesCount := engine.RandomRange(2, 5)
	y := engine.RandomRange(5, 15)
	for i := 0; i < treesCount; i++ {
		y += engine.RandomRange(5, 10)

		tree := archetypes.NewTree(forestToTripType(forestType))
		t.GetWorldSpace().AddChild(tree)
		tree.Translate(
			float64(rand.Intn(50)+5),
			float64(y),
		)
	}

	t.ResourcesSource.Resources.Wood = treesCount

	return t
}

func forestToTripType(forestType ForestType) archetypes.TreeType {
	switch forestType {
	case ForestStandard:
		return archetypes.TreeStandard
	case ForestPine:
		return archetypes.TreePine
	case ForestMixed:
		r := engine.RandomRange(0, 1)
		if r == 0 {
			return archetypes.TreeStandard
		} else {
			return archetypes.TreePine
		}
	default:
		return archetypes.TreeStandard
	}
}

func NewMountainsTile(groundType GroundType, mountainType archetypes.MountainType) Tile {
	t := NewGroundTile(groundType)

	width := int(t.Clickable.Bounds.Width)
	height := int(t.Clickable.Bounds.Height)
	widthOffset := width / 4.0
	heightOffset := height / 4.0

	mountain := archetypes.NewMountain(mountainType)
	mountain.GetWorldSpace().SetLocal(
		float64(engine.RandomRange(widthOffset, width-widthOffset)),
		float64(engine.RandomRange(heightOffset, height-heightOffset)),
	)
	t.GetWorldSpace().AddChild(mountain)

	switch mountainType {
	case archetypes.MountainStone:
		t.ResourcesSource.Resources.Stone = 1
	case archetypes.MountainIron:
		t.ResourcesSource.Resources.Iron = 1
	case archetypes.MountainGold:
		t.ResourcesSource.Resources.Gold = 1
	}

	return t
}

func NewBuildingTile(groundType GroundType, buildingType components.BuildingType) Tile {
	t := NewGroundTile(groundType)

	buildingPos := engine.Vector{
		X: t.Clickable.Bounds.Width / 2.0,
		Y: t.Clickable.Bounds.Height,
	}

	building := archetypes.NewBuilding(buildingPos, buildingType)
	t.GetWorldSpace().AddChild(building)

	return t
}
