package engine

import (
	"fmt"
	"sync/atomic"
)

type EntityID int64

var (
	nextEntityID int64
)

type Entity interface {
	ID() EntityID
	Equals(Entity) bool
}

type BaseEntity struct {
	id       EntityID
	parent   Entity
	children []Entity
}

func NewBaseEntity() *BaseEntity {
	return &BaseEntity{
		id: EntityID(atomic.AddInt64(&nextEntityID, 1)),
	}
}

func (e BaseEntity) ID() EntityID {
	return e.id
}

func (e BaseEntity) Equals(other Entity) bool {
	return e.ID() == other.ID()
}

func (e *BaseEntity) String() string {
	return fmt.Sprintf("Entity [%v]", e.id)
}
