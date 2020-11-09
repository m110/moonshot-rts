package archetypes

import (
	"time"

	"github.com/m110/moonshot-rts/internal/assets/sprites"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type Building struct {
	Object
	*components.Selectable
	*components.Clickable
	*components.Collider
	*components.UnitSpawner
	*components.TimeActions
}

func NewBuilding(position engine.Vector, buildingType components.BuildingType) Building {
	var options []components.UnitSpawnerOption
	switch buildingType {
	case components.BuildingSettlement:
		options = []components.UnitSpawnerOption{
			{
				Class:     components.ClassWorker,
				SpawnTime: 5 * time.Second,
			},
		}
	case components.BuildingBarracks:
		options = []components.UnitSpawnerOption{
			{
				Class:     components.ClassWarrior,
				SpawnTime: 5 * time.Second,
			},
			{
				Class:     components.ClassKnight,
				SpawnTime: 10 * time.Second,
			},
		}
	case components.BuildingChapel:
		options = []components.UnitSpawnerOption{
			{
				Class:     components.ClassPriest,
				SpawnTime: 15 * time.Second,
			},
		}
	}

	bottomSprite, topSprite := SpritesForBuilding(buildingType)

	w, h := bottomSprite.Size()
	if !topSprite.IsZero() {
		h += topSprite.Height()
	}
	overlay := NewOverlay(w, h, engine.PivotBottom)

	b := Building{
		Object: NewObject(bottomSprite, components.LayerObjects),
		Selectable: &components.Selectable{
			Overlay: overlay,
		},
		Clickable: &components.Clickable{
			Bounds: components.BoundsFromSprite(overlay.GetDrawable().Sprite),
		},
		Collider: &components.Collider{
			Bounds: components.BoundsFromSprite(bottomSprite),
			Layer:  components.CollisionLayerBuildings,
		},
		UnitSpawner: &components.UnitSpawner{
			Options: options,
		},
		TimeActions: &components.TimeActions{},
	}

	b.GetWorldSpace().AddChild(overlay)
	b.WorldSpace.SetLocal(position.X, position.Y)

	if !topSprite.IsZero() {
		topSprite := NewObject(topSprite, components.LayerForeground)
		topSprite.GetWorldSpace().SetLocal(0, float64(-bottomSprite.Height()))
		b.GetWorldSpace().AddChild(topSprite)
	}

	return b
}

func SpritesForBuilding(buildingType components.BuildingType) (bottomSprite engine.Sprite, topSprite engine.Sprite) {
	switch buildingType {
	case components.BuildingSettlement:
		return sprites.Castle, sprites.CastleTop
	case components.BuildingBarracks:
		return sprites.Barracks, sprites.BarracksTop
	case components.BuildingChapel:
		return sprites.Chapel, sprites.ChapelTop
	case components.BuildingForge:
		return sprites.Forge, engine.Sprite{}
	case components.BuildingTower:
		return sprites.Tower, sprites.TowerTop
	}

	return engine.Sprite{}, engine.Sprite{}
}

func SpriteForBuilding(buildingType components.BuildingType) engine.Sprite {
	bottom, top := SpritesForBuilding(buildingType)
	if top.IsZero() {
		return bottom
	}

	bottomWidth, bottomHeight := bottom.Size()
	_, topHeight := top.Size()

	width := bottomWidth
	height := bottomHeight + topHeight

	sprite := engine.NewBlankSprite(width, height)
	sprite.DrawAtPosition(top, width/2, topHeight)
	sprite.DrawAtPosition(bottom, width/2, height)

	return sprite
}
