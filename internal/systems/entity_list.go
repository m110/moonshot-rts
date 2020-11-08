package systems

import "github.com/m110/moonshot-rts/internal/engine"

type EntityList struct {
	entities []engine.Entity
}

func (l *EntityList) All() []engine.Entity {
	return l.entities
}

func (l *EntityList) Clear() {
	l.entities = nil
}

func (l EntityList) Empty() bool {
	return len(l.entities) == 0
}

func (l *EntityList) ByID(id engine.EntityID) (engine.Entity, bool) {
	for _, e := range l.entities {
		if e.ID() == id {
			return e, true
		}
	}

	return nil, false
}

func (l *EntityList) Add(entity engine.Entity) {
	l.entities = append(l.entities, entity)
}

func (l *EntityList) Remove(entity engine.Entity) {
	foundIndex := -1
	for i, e := range l.entities {
		if e.Equals(entity) {
			foundIndex = i
			break
		}
	}

	if foundIndex >= 0 {
		l.entities = append(l.entities[:foundIndex], l.entities[foundIndex+1:]...)
	}
}

type EntityMap struct {
	entities map[engine.EntityID]engine.Entity
}

func (m *EntityMap) All() []engine.Entity {
	var entities []engine.Entity
	for _, e := range m.entities {
		entities = append(entities, e)
	}
	return entities
}

func (m *EntityMap) Clear() {
	m.entities = map[engine.EntityID]engine.Entity{}
}

func (m EntityMap) Empty() bool {
	return len(m.entities) == 0
}

func (m *EntityMap) ByID(id engine.EntityID) (engine.Entity, bool) {
	if m.entities == nil {
		return nil, false
	}
	e, ok := m.entities[id]
	return e, ok
}

func (m *EntityMap) Add(entity engine.Entity) {
	if m.entities == nil {
		m.entities = map[engine.EntityID]engine.Entity{}
	}
	m.entities[entity.ID()] = entity
}

func (m *EntityMap) Remove(entity engine.Entity) {
	delete(m.entities, entity.ID())
}
