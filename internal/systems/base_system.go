package systems

import "github.com/m110/rts/internal/engine"

type BaseSystem struct {
	Config   Config
	EventBus *engine.EventBus
	Spawner  spawner
}

func NewBaseSystem(config Config, eventBus *engine.EventBus, spawner spawner) BaseSystem {
	return BaseSystem{
		Config:   config,
		EventBus: eventBus,
		Spawner:  spawner,
	}
}
