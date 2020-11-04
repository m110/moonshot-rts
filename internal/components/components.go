package components

import (
	"github.com/m110/rts/internal/engine"
)

const AllLayers = 6

type DrawingLayer int

const (
	LayerBackground DrawingLayer = iota
	LayerGround
	LayerObjects
	LayerUnits
	LayerForeground
	LayerUI
)

const (
	UILayerBackground DrawingLayer = iota
	UILayerText
)

// TODO find better suffix than Owner
type DrawableOwner interface {
	GetDrawable() *Drawable
}

type Drawable struct {
	Sprite engine.Sprite
	Layer  DrawingLayer

	Disabled bool
}

func (d *Drawable) Enable() {
	d.Disabled = false
}

func (d *Drawable) Disable() {
	d.Disabled = true
}

type BoxBoundaryOwner interface {
	GetBoxBoundary() *BoxBoundary
}

type BoxBoundary struct {
	// Bounds defines position relative to WorldSpace
	Bounds engine.Rect
}

func BoxBoundaryFromSprite(sprite engine.Sprite) *BoxBoundary {
	w, h := sprite.Size()
	pivot := sprite.Pivot()
	bounds := engine.Rect{
		Position: engine.Vector{
			X: -pivot.X,
			Y: -pivot.Y,
		},
		Width:  float64(w),
		Height: float64(h),
	}

	return &BoxBoundary{
		Bounds: bounds,
	}
}

type SelectableOwner interface {
	GetSelectable() *Selectable
}

type Selectable struct {
	Selected        bool
	GroupSelectable bool
	Overlay         DrawableOwner
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
	Class Class
}

type UnitSpawnerOwner interface {
	GetUnitSpawner() *UnitSpawner
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
