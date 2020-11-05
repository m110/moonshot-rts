package components

import (
	"github.com/m110/moonshot-rts/internal/engine"
)

type SelectableOwner interface {
	GetSelectable() *Selectable
}

type Selectable struct {
	Selected bool
	Overlay  DrawableOwner
}

func (s *Selectable) Select() {
	s.Selected = true
	s.Overlay.GetDrawable().Disabled = false
}

func (s *Selectable) Unselect() {
	s.Selected = false
	s.Overlay.GetDrawable().Disabled = true
}

type MovableOwner interface {
	GetMovable() *Movable
}

type Movable struct {
	Target *engine.Vector
}

func (m *Movable) ClearTarget() {
	m.Target = nil
}

func (m *Movable) SetTarget(target engine.Vector) {
	m.Target = &target
}

type MovementArea struct {
	// TODO ?
	Speed float64
}

type Size struct {
	Width  int
	Height int
}

func (s *Size) Set(w int, h int) {
	s.Width = w
	s.Height = h
}

type SizeOwner interface {
	GetSize() *Size
}

type UnitSpawner struct {
	Classes []Class
}

type UnitSpawnerOwner interface {
	GetUnitSpawner() *UnitSpawner
}

type BuildingType int

const (
	BuildingSettlement BuildingType = iota
	BuildingBarracks
	BuildingChapel
	BuildingForge
	BuildingTower
)

type SettlementType int

const (
	SettlementColony SettlementType = iota
	SettlementVillage
	SettlementCastle
)

type Builder struct {
	Buildings []BuildingType
}

type BuilderOwner interface {
	GetBuilder() *Builder
}

type Team int

const (
	TeamBlue Team = iota
	TeamRed
	TeamGreen
	TeamGray
)

type Class int

const (
	ClassWorker Class = iota
	ClassWarrior
	ClassKnight
	ClassPriest
	ClassKing
)

type Citizen struct {
	Team  Team
	Class Class
}
