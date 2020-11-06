package components

import (
	"time"

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

type UnitSpawnerOption struct {
	Class     Class
	SpawnTime time.Duration
}

type UnitSpawner struct {
	Options []UnitSpawnerOption
	Queue   []UnitSpawnerOption
	Timer   *engine.CountdownTimer
}

func (u *UnitSpawner) AddToQueue(option UnitSpawnerOption) {
	u.Queue = append(u.Queue, option)
}

func (u *UnitSpawner) PopFromQueue() (UnitSpawnerOption, bool) {
	if len(u.Queue) > 0 {
		opt := u.Queue[0]
		u.Queue = u.Queue[1:]
		return opt, true
	}

	return UnitSpawnerOption{}, false
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

type BuilderOption struct {
	BuildingType BuildingType
	SpawnTime    time.Duration
}

type Builder struct {
	Options []BuilderOption
	Queue   []BuilderOption
	Timer   *engine.CountdownTimer
}

func (b *Builder) AddToQueue(option BuilderOption) {
	b.Queue = append(b.Queue, option)
}

func (b *Builder) PopFromQueue() (BuilderOption, bool) {
	if len(b.Queue) > 0 {
		opt := b.Queue[0]
		b.Queue = b.Queue[1:]
		return opt, true
	}

	return BuilderOption{}, false
}

type BuilderOwner interface {
	GetBuilder() *Builder
}

type TimeActions struct {
	Timers []*engine.CountdownTimer
}

type TimeActionsOwner interface {
	GetTimeActions() *TimeActions
}

func (t *TimeActions) AddTimer(timer *engine.CountdownTimer) {
	t.Timers = append(t.Timers, timer)
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

type ProgressBarSprites struct {
	Left  engine.Sprite
	Mid   engine.Sprite
	Right engine.Sprite
}

type ProgressBar struct {
	Background ProgressBarSprites
	Foreground ProgressBarSprites
	Progress   float64
}

func (p *ProgressBar) SetProgress(progress float64) {
	if progress < 0 {
		progress = 0
	} else if progress > 1.0 {
		progress = 1.0
	}
	p.Progress = progress
}

type ProgressBarOwner interface {
	GetProgressBar() *ProgressBar
}
