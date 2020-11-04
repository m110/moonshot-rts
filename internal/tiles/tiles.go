package tiles

import (
	"math/rand"

	"github.com/m110/rts/internal/atlas"
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
	"github.com/m110/rts/internal/objects"
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
		sprite = atlas.Grass1
	case GroundSand:
		sprite = atlas.Sand1
	case GroundSea:
		sprite = atlas.Water1
	}

	return Tile{
		BaseEntity: engine.NewBaseEntity(),
		WorldSpace: &components.WorldSpace{},
		Drawable: &components.Drawable{
			Sprite: sprite,
			Layer:  components.LayerGround,
		},
		BoxBoundary: components.BoxBoundaryFromSprite(sprite),
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
	for i := 0; i <= treesCount; i++ {
		y += engine.RandomRange(5, 10)

		tree := objects.NewTree(forestToTripType(forestType))
		t.GetWorldSpace().AddChild(t, tree)
		tree.Translate(
			float64(rand.Intn(50)+5),
			float64(y),
		)
	}

	return t
}

func forestToTripType(forestType ForestType) objects.TreeType {
	switch forestType {
	case ForestStandard:
		return objects.TreeStandard
	case ForestPine:
		return objects.TreePine
	case ForestMixed:
		r := engine.RandomRange(0, 1)
		if r == 0 {
			return objects.TreeStandard
		} else {
			return objects.TreePine
		}
	default:
		return objects.TreeStandard
	}
}

func NewMountainsTile(groundType GroundType, mountainType objects.MountainType) Tile {
	t := NewGroundTile(groundType)

	width := int(t.BoxBoundary.Bounds.Width)
	height := int(t.BoxBoundary.Bounds.Height)
	widthOffset := width / 4.0
	heightOffset := height / 4.0

	mountain := objects.NewMountain(mountainType)
	mountain.GetWorldSpace().SetLocal(
		float64(engine.RandomRange(widthOffset, width-widthOffset)),
		float64(engine.RandomRange(heightOffset, height-heightOffset)),
	)
	t.GetWorldSpace().AddChild(t, mountain)

	return t
}

func NewBuildingTile(groundType GroundType, buildingType objects.BuildingType) Tile {
	t := NewGroundTile(groundType)

	buildingPos := engine.Vector{
		X: t.BoxBoundary.Bounds.Width / 2.0,
		Y: t.BoxBoundary.Bounds.Height,
	}

	building := objects.NewBuilding(buildingPos, buildingType)

	buildingHeight := float64(building.Sprite.Height())
	height := t.BoxBoundary.Bounds.Height

	// TODO This is a hack, but probably good enough for now
	if buildingHeight > height {
		p := building.Sprite.Pivot()
		p.Y += (buildingHeight - height) / 2.0
		building.Sprite.SetPivot(p)
	}

	t.GetWorldSpace().AddChild(t, building)

	return t
}
