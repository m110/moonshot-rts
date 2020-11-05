package objects

import (
	"github.com/m110/moonshot-rts/internal/atlas"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"golang.org/x/image/colornames"
)

type Building struct {
	Object
	*components.Selectable
	*components.Clickable
	*components.UnitSpawner
}

func NewBuilding(position engine.Vector, buildingType components.BuildingType) Building {
	var classes []components.Class
	switch buildingType {
	case components.BuildingSettlement:
		classes = []components.Class{components.ClassWorker}
	case components.BuildingBarracks:
		classes = []components.Class{components.ClassWarrior, components.ClassKnight}
	case components.BuildingChapel:
		classes = []components.Class{components.ClassPriest}
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
			Bounds: components.BoundsFromSprite(bottomSprite),
		},
		UnitSpawner: &components.UnitSpawner{
			Classes: classes,
		},
	}

	b.GetWorldSpace().AddChild(b, overlay)
	b.WorldSpace.SetLocal(position.X, position.Y)

	if !topSprite.IsZero() {
		topSprite := NewObject(topSprite, components.LayerForeground)
		topSprite.GetWorldSpace().SetLocal(0, float64(-bottomSprite.Height()))
		b.GetWorldSpace().AddChild(b, topSprite)
	}

	return b
}

func (b Building) GetSelectable() *components.Selectable {
	return b.Selectable
}

func (b Building) GetClickable() *components.Clickable {
	return b.Clickable
}

func (b Building) GetUnitSpawner() *components.UnitSpawner {
	return b.UnitSpawner
}

func SpritesForBuilding(buildingType components.BuildingType) (bottomSprite engine.Sprite, topSprite engine.Sprite) {
	switch buildingType {
	case components.BuildingSettlement:
		return atlas.Castle, atlas.CastleTop
	case components.BuildingBarracks:
		return atlas.Barracks, atlas.BarracksTop
	case components.BuildingChapel:
		return atlas.Chapel, atlas.ChapelTop
	case components.BuildingForge:
		return atlas.Forge, engine.Sprite{}
	case components.BuildingTower:
		return atlas.Tower, atlas.TowerTop
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
	sprite.Image().Fill(colornames.Green)
	sprite.DrawAtPosition(top, width/2, topHeight)
	sprite.DrawAtPosition(bottom, width/2, height)

	return sprite
}
