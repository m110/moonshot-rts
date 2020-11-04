package tiles

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"golang.org/x/image/colornames"
)

func NewHighlightTile(width int, height int) Tile {
	c := colornames.Cyan
	c.A = 75
	highlightedSprite := engine.NewFilledSprite(width, height, c)

	return Tile{
		BaseEntity: engine.NewBaseEntity(),
		WorldSpace: &components.WorldSpace{},
		Drawable: &components.Drawable{
			Sprite:   highlightedSprite,
			Layer:    components.LayerForeground,
			Disabled: true,
		},
	}
}
