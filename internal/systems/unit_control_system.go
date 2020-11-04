package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
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
}

func NewUnitControlSystem(config Config, eventBus *engine.EventBus, spawner spawner) *UnitControlSystem {
	return &UnitControlSystem{
		base: NewBaseSystem(config, eventBus, spawner),
	}
}

func (u *UnitControlSystem) Start() {
	u.base.EventBus.Subscribe(EntitySelected{}, u)
	u.base.EventBus.Subscribe(EntityUnselected{}, u)
}

func (u UnitControlSystem) Update(dt float64) {
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

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !u.activeEntities.Empty() {
			x, y := ebiten.CursorPosition()
			target := engine.Vector{X: float64(x), Y: float64(y)}

			for _, e := range u.activeEntities.All() {
				entity := e.(unitControlEntity)
				entity.GetMovable().SetTarget(target)
			}
		}
	}
}

func (u UnitControlSystem) Draw(canvas engine.Sprite) {
}

func (u *UnitControlSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntitySelected:
		foundEntity, ok := u.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		u.activeEntities.Add(foundEntity)
	case EntityUnselected:
		u.activeEntities.Remove(event.Entity)
	}
}

func (u *UnitControlSystem) Add(entity unitControlEntity) {
	u.entities.Add(entity)
}

func (u *UnitControlSystem) Remove(entity engine.Entity) {
	u.entities.Remove(entity)
}
