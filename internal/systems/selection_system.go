package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type selectionEntity interface {
	engine.Entity
	components.WorldSpaceOwner
	components.SelectableOwner
	components.ClickableOwner
}

type SelectionSystem struct {
	BaseSystem

	entities EntityList

	selectedEntities []selectionEntity
}

type EntitySelected struct {
	Entity engine.Entity
}

type EntityUnselected struct {
	Entity engine.Entity
}

func NewSelectionSystem(base BaseSystem) *SelectionSystem {
	s := &SelectionSystem{
		BaseSystem: base,
	}

	s.EventBus.Subscribe(EntityClicked{}, s)
	s.EventBus.Subscribe(EntitiesClicked{}, s)
	s.EventBus.Subscribe(PointClicked{}, s)

	return s
}

func (s *SelectionSystem) Start() {
}

func (s *SelectionSystem) Update(dt float64) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.unselectCurrentEntities()
	}
}

func (s *SelectionSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntityClicked:
		if len(s.selectedEntities) > 0 {
			_, ok := s.entities.ByID(event.Entity.ID())
			if ok {
				s.unselectCurrentEntities()
				return
			}
		}
		s.selectEntity(event.Entity)
	case EntitiesClicked:
		if len(s.selectedEntities) > 0 {
			s.unselectCurrentEntities()
			return
		}

		for _, entity := range event.Entities {
			s.selectEntity(entity)
		}
	case PointClicked:
		s.unselectCurrentEntities()
	}
}
func (s *SelectionSystem) selectEntity(e engine.Entity) bool {
	_, ok := s.entities.ByID(e.ID())
	if !ok {
		return false
	}

	entity := e.(selectionEntity)
	entity.GetSelectable().Select()
	s.selectedEntities = append(s.selectedEntities, entity)

	s.EventBus.Publish(EntitySelected{
		Entity: entity,
	})

	return true
}

func (s *SelectionSystem) unselectCurrentEntities() {
	if len(s.selectedEntities) == 0 {
		return
	}

	for _, e := range s.selectedEntities {
		e.GetSelectable().Unselect()
		s.EventBus.Publish(EntityUnselected{Entity: e})
	}
	s.selectedEntities = nil
}

func (s *SelectionSystem) Add(entity selectionEntity) {
	s.entities.Add(entity)
}

func (s SelectionSystem) Remove(entity engine.Entity) {
	s.entities.Remove(entity)
}
