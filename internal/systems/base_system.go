package systems

import "github.com/m110/moonshot-rts/internal/engine"

type BaseSystem struct {
	Config   Config
	EventBus *engine.EventBus
	Spawner  Spawner
}

func NewBaseSystem(config Config, eventBus *engine.EventBus, systemsProvider systemsProvider) BaseSystem {
	return BaseSystem{
		Config:   config,
		EventBus: eventBus,
		Spawner:  NewSpawner(systemsProvider),
	}
}
