package tiles

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type Tile struct {
	*engine.BaseEntity
	*components.WorldSpace
	*components.Drawable
	*components.Clickable
	*components.Collider
}
