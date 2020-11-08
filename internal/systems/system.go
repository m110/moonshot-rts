package systems

import (
	"github.com/m110/moonshot-rts/internal/engine"
)

type System interface {
	Start()
	Update(dt float64)
	Remove(entity engine.Entity)
}

type Drawer interface {
	Draw(canvas engine.Sprite)
}

type systemsProvider interface {
	Systems() []System
}
