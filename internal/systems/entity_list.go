package systems

import "github.com/m110/moonshot-rts/internal/engine"

type EntityList struct {
	entities []engine.Entity
}

func (l *EntityList) All() []engine.Entity {
	return l.entities
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

func (l *EntityMap) All() []engine.Entity {
	var entities []engine.Entity
	for _, e := range l.entities {
		entities = append(entities, e)
	}
	return entities
}

func (l EntityMap) Empty() bool {
	return len(l.entities) == 0
}

func (l *EntityMap) ByID(id engine.EntityID) (engine.Entity, bool) {
	if l.entities == nil {
		return nil, false
	}
	e, ok := l.entities[id]
	return e, ok
}

func (l *EntityMap) Add(entity engine.Entity) {
	if l.entities == nil {
		l.entities = map[engine.EntityID]engine.Entity{}
	}
	l.entities[entity.ID()] = entity
}

func (l *EntityMap) Remove(entity engine.Entity) {
	delete(l.entities, entity.ID())
}
