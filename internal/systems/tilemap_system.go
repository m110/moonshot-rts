package systems

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m110/moonshot-rts/internal/archetypes"
	"github.com/m110/moonshot-rts/internal/archetypes/tiles"
	"github.com/m110/moonshot-rts/internal/assets/sprites"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"golang.org/x/image/colornames"
)

type TilemapConfig struct {
	OffsetX int
	OffsetY int

	Width  int
	Height int

	TileWidth  int
	TileHeight int
}

func (t TilemapConfig) TotalWidth() int {
	return t.Width * t.TileWidth
}

func (t TilemapConfig) TotalHeight() int {
	return t.Height * t.TileHeight
}

type TilemapSystem struct {
	BaseSystem

	world archetypes.Object

	tiles      []tiles.Tile
	debugTiles []archetypes.Object

	castlePosition engine.Point
}

func NewTilemapSystem(base BaseSystem) *TilemapSystem {
	return &TilemapSystem{
		BaseSystem: base,
	}
}

func (t *TilemapSystem) Start() {
	worldSprite := engine.NewFilledSprite(
		t.Config.TileMap.TotalWidth(),
		t.Config.TileMap.TotalHeight(),
		colornames.White,
	)

	t.world = archetypes.NewObject(worldSprite, components.LayerBackground)
	t.world.GetWorldSpace().Translate(
		float64(t.Config.TileMap.OffsetX),
		float64(t.Config.TileMap.OffsetY),
	)

	t.Spawner.Spawn(t.world)

	t.spawnTiles()
	t.spawnDebugTiles()
	t.spawnUnits()
}

func (t *TilemapSystem) spawnTiles() {
	w := t.Config.TileMap.Width
	h := t.Config.TileMap.Height

	desertEndX := engine.RandomRange(w/4, w/2)
	desertEndY := engine.RandomRange(h/4, h/2)

	seaStartX := engine.RandomRange(w-w/3, w)
	seaStartY := engine.RandomRange(h-h/3, h)

	t.castlePosition = engine.Point{
		X: engine.RandomRange(2, seaStartX-2),
		Y: engine.RandomRange(2, seaStartY-2),
	}

	for y := 0; y < t.Config.TileMap.Height; y++ {
		for x := 0; x < t.Config.TileMap.Width; x++ {
			position := engine.Vector{
				X: float64(x * t.Config.TileMap.TileWidth),
				Y: float64(y * t.Config.TileMap.TileWidth),
			}

			var forestChance int
			var mountainChance int

			var ground tiles.GroundType
			if x < desertEndX && y < desertEndY {
				forestChance = 0
				mountainChance = 4
				ground = tiles.GroundSand
			} else if x > seaStartX && y > seaStartY {
				forestChance = 0
				mountainChance = 0
				ground = tiles.GroundSea
			} else {
				forestChance = 3
				mountainChance = 2
				ground = tiles.GroundGrass
			}

			var tile tiles.Tile
			if x == t.castlePosition.X && y == t.castlePosition.Y {
				tile = tiles.NewBuildingTile(ground, components.BuildingSettlement)
			} else {
				if rand.Intn(10) < forestChance {
					forestType := tiles.ForestType(rand.Intn(3))
					tile = tiles.NewForestTile(ground, forestType)
				} else if rand.Intn(10) < mountainChance {
					mountainsType := archetypes.MountainType(rand.Intn(3))
					tile = tiles.NewMountainsTile(ground, mountainsType)
				} else {
					tile = tiles.NewGroundTile(ground)
				}
			}

			t.tiles = append(t.tiles, tile)
			t.world.GetWorldSpace().AddChild(tile)
			tile.GetWorldSpace().Translate(position.X, position.Y)
			t.Spawner.Spawn(tile)
		}
	}

}

func (t *TilemapSystem) spawnDebugTiles() {
	for y := 0; y < t.Config.TileMap.Height; y++ {
		for x := 0; x < t.Config.TileMap.Width; x++ {
			pos := engine.Vector{
				X: float64(x * t.Config.TileMap.TileWidth),
				Y: float64(y * t.Config.TileMap.TileHeight),
			}
			tile := tiles.NewDebugTile(t.Config.TileMap.TileWidth, t.Config.TileMap.TileHeight)
			t.world.GetWorldSpace().AddChild(tile)
			tile.GetWorldSpace().Translate(pos.X, pos.Y)
			t.debugTiles = append(t.debugTiles, tile)
			t.Spawner.Spawn(tile)
		}
	}
}

func (t TilemapSystem) spawnUnits() {
	unitsX := func(o int) float64 {
		return float64((t.castlePosition.X+o)*t.Config.TileMap.TileWidth + t.Config.TileMap.TileWidth/2)
	}
	unitsY := func(o int) float64 {
		return float64((t.castlePosition.Y+o)*t.Config.TileMap.TileHeight + t.Config.TileMap.TileHeight/2)
	}

	spriteGetter := atlasSpriteGetter{}

	king := archetypes.NewUnit(components.TeamBlue, components.ClassKing, spriteGetter)
	t.world.GetWorldSpace().AddChild(king)
	king.GetWorldSpace().Translate(unitsX(0), unitsY(1))
	t.Spawner.Spawn(king)
}

func (t TilemapSystem) Update(_ float64) {
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		for _, t := range t.debugTiles {
			t.GetDrawable().Disabled = !t.GetDrawable().Disabled
		}
	}
}

func (t TilemapSystem) TileAtPosition(position engine.Vector) (tiles.Tile, bool) {
	for _, tile := range t.tiles {
		bounds := tile.GetClickable().Bounds
		bounds.Position = bounds.Position.Add(tile.GetWorldSpace().WorldPosition())
		if bounds.WithinBounds(position) {
			return tile, true
		}
	}

	return tiles.Tile{}, false
}

func (t TilemapSystem) Remove(e engine.Entity) {}

type atlasSpriteGetter struct{}

func (a atlasSpriteGetter) SpriteForUnit(team components.Team, class components.Class) engine.Sprite {
	return sprites.Units[team][class].Random()
}
