package systems

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type timeActionsEntity interface {
	engine.Entity
	components.TimeActionsOwner
}

type TimeActionsSystem struct {
	base     BaseSystem
	entities EntityList
}

func NewTimeActionsSystem(config Config, eventBus *engine.EventBus, spawner spawner) *TimeActionsSystem {
	return &TimeActionsSystem{
		base: NewBaseSystem(config, eventBus, spawner),
	}
}

func (t *TimeActionsSystem) Start() {
}

func (t *TimeActionsSystem) Update(dt float64) {
	for _, e := range t.entities.All() {
		entity := e.(timeActionsEntity)

		timers := entity.GetTimeActions().Timers
		entity.GetTimeActions().Timers = nil

		for _, t := range timers {
			t.Update(dt)
			if !t.Done() {
				entity.GetTimeActions().AddTimer(t)
			}
		}
	}
}

func (t TimeActionsSystem) Draw(canvas engine.Sprite) {
}

func (t *TimeActionsSystem) Add(entity timeActionsEntity) {
	t.entities.Add(entity)
}

func (t *TimeActionsSystem) Remove(entity engine.Entity) {
	t.entities.Remove(entity)
}
