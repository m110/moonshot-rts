package systems

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type ResourcesUpdated struct {
	Resources components.Resources
}

type ResourcesSystem struct {
	base BaseSystem

	resources components.Resources
}

func NewResourcesSystem(config Config, eventBus *engine.EventBus, spawner spawner) *ResourcesSystem {
	return &ResourcesSystem{
		base: NewBaseSystem(config, eventBus, spawner),
	}
}

func (r ResourcesSystem) Start() {
	r.resources = components.Resources{}
	r.base.EventBus.Publish(ResourcesUpdated{Resources: r.resources})
}

func (r ResourcesSystem) Update(dt float64) {
}

func (r ResourcesSystem) Draw(canvas engine.Sprite) {
}

func (r ResourcesSystem) Remove(entity engine.Entity) {
}
