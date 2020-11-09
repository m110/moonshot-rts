package tiles

import (
	"github.com/m110/moonshot-rts/internal/archetypes"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"golang.org/x/image/colornames"
)

func NewDebugTile(width int, height int) archetypes.Object {
	c := colornames.Pink
	c.A = 120
	debugSprite := engine.NewFilledSprite(width-2, height-2, c)

	debugPointSprite := engine.NewFilledSprite(1, 1, colornames.Red)

	sprite := engine.NewBlankSprite(width, height)

	sprite.DrawAtPosition(debugPointSprite, 0, 0)
	sprite.DrawAtPosition(debugSprite, 1, 1)
	sprite.DrawAtPosition(debugPointSprite, width/2, height/2)

	return archetypes.Object{
		BaseEntity: engine.NewBaseEntity(),
		WorldSpace: &components.WorldSpace{},
		Drawable: &components.Drawable{
			Sprite:   sprite,
			Layer:    components.LayerForeground,
			Disabled: true,
		},
	}
}
