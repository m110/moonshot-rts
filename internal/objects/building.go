package objects

import (
	"github.com/m110/rts/internal/atlas"
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
)

type BuildingType int

const (
	BuildingSettlement BuildingType = iota
	BuildingBarracks
)

type SettlementType int

const (
	SettlementColony SettlementType = iota
	SettlementVillage
	SettlementCastle
)

type Building struct {
	Object
	*components.Selectable
	*components.BoxBoundary
	*components.UnitSpawner
}

func NewBuilding(position engine.Vector, buildingType BuildingType) Building {
	var bottomSprite, topSprite engine.Sprite
	var class components.Class

	switch buildingType {
	case BuildingSettlement:
		bottomSprite = atlas.Castle
		topSprite = atlas.CastleTop
		class = components.ClassWorker
	case BuildingBarracks:
		bottomSprite = atlas.Barracks
		topSprite = atlas.BarracksTop
		class = components.ClassWarrior
	}

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
		BoxBoundary: components.BoxBoundaryFromSprite(bottomSprite),
		UnitSpawner: &components.UnitSpawner{
			Class: class,
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

func (b Building) GetBoxBoundary() *components.BoxBoundary {
	return b.BoxBoundary
}

func (b Building) GetUnitSpawner() *components.UnitSpawner {
	return b.UnitSpawner
}
