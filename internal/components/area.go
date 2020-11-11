package components

import "github.com/m110/moonshot-rts/internal/engine"

type Area struct {
}

// TODO This should be a "has" relation, not "is"
type AreaOccupant struct {
	OccupiedArea engine.Entity
	NextArea     engine.Entity
}
