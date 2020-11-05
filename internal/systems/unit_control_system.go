package systems

import (
	"github.com/m110/moonshot-rts/internal/atlas"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/objects"
)

type unitControlEntity interface {
	engine.Entity
	components.WorldSpaceOwner
	components.MovableOwner
}

type UnitControlSystem struct {
	base BaseSystem

	entities EntityList

	activeEntities EntityList

	buildIcon    engine.Sprite
	actionButton *objects.PanelButton
	actionsPanel *objects.Panel
}

func NewUnitControlSystem(config Config, eventBus *engine.EventBus, spawner spawner) *UnitControlSystem {
	return &UnitControlSystem{
		base: NewBaseSystem(config, eventBus, spawner),
	}
}

func (u *UnitControlSystem) Start() {
	u.buildIcon = atlas.Hammer
	u.buildIcon.Scale(engine.Vector{X: 0.5, Y: 0.5})

	u.base.EventBus.Subscribe(PointClicked{}, u)
	u.base.EventBus.Subscribe(EntitySelected{}, u)
	u.base.EventBus.Subscribe(EntityUnselected{}, u)
}

func (u UnitControlSystem) Update(dt float64) {
	u.moveEntities(dt)
}

func (u *UnitControlSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntitySelected:
		foundEntity, ok := u.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		u.activeEntities.Add(foundEntity)
		if len(u.activeEntities.All()) > 1 {
			u.hideActionButton()
		} else {
			u.showActionButton()
		}
	case EntityUnselected:
		u.activeEntities.Remove(event.Entity)
		u.hideActionButton()
		u.hideActionsPanel()
	case PointClicked:
		for _, e := range u.activeEntities.All() {
			entity := e.(unitControlEntity)
			entity.GetMovable().SetTarget(event.Point)
		}
	}
}

func (u *UnitControlSystem) moveEntities(dt float64) {
	for _, e := range u.entities.All() {
		entity := e.(unitControlEntity)
		if entity.GetMovable().Target != nil {
			if entity.GetWorldSpace().WorldPosition().Distance(*entity.GetMovable().Target) < 1.0 {
				entity.GetMovable().ClearTarget()
			} else {
				direction := entity.GetMovable().Target.Sub(entity.GetWorldSpace().WorldPosition()).Normalized()
				entity.GetWorldSpace().Translate(direction.Mul(50 * dt).Unpack())
			}
		}
	}
}

func (u *UnitControlSystem) showActionButton() {
	button := objects.NewPanelButton(components.UIColorBrown, u.buildIcon, func() {
		u.hideActionButton()
		u.showActionPanel()
	})

	entity := u.activeEntities.All()[0].(unitControlEntity)
	entity.GetWorldSpace().AddChild(entity, button)

	u.base.Spawner.SpawnPanelButton(button)
	u.actionButton = &button
}

func (u *UnitControlSystem) hideActionButton() {
	if u.actionButton == nil {
		return
	}

	u.base.Spawner.RemovePanelButton(*u.actionButton)
	u.actionButton = nil
}

func (u *UnitControlSystem) showActionPanel() {
	var configs []objects.ButtonConfig
	/*
		for _, class := range spawner.GetUnitSpawner().Classes {
			configs = append(configs, objects.ButtonConfig{})
		}
	*/

	panel := objects.NewFourButtonPanel(configs)
	u.base.Spawner.SpawnPanel(panel)

	entity := u.activeEntities.All()[0].(unitControlEntity)
	entity.GetWorldSpace().AddChild(entity, panel)

	u.actionsPanel = &panel
}

func (u *UnitControlSystem) hideActionsPanel() {
	if u.actionsPanel == nil {
		return
	}

	u.base.Spawner.RemovePanel(*u.actionsPanel)
	u.actionsPanel = nil
}

func (u *UnitControlSystem) Add(entity unitControlEntity) {
	u.entities.Add(entity)
}

func (u *UnitControlSystem) Remove(entity engine.Entity) {
	u.entities.Remove(entity)
}

func (u UnitControlSystem) Draw(canvas engine.Sprite) {
}
