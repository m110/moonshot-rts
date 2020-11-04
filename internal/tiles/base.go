package tiles

import (
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
)

type Tile struct {
	*engine.BaseEntity
	*components.WorldSpace
	*components.Drawable
	*components.BoxBoundary
}

func (t Tile) GetWorldSpace() *components.WorldSpace {
	return t.WorldSpace
}

func (t Tile) GetDrawable() *components.Drawable {
	return t.Drawable
}

func (t Tile) GetBoxBoundary() *components.BoxBoundary {
	return t.BoxBoundary
}

type BaseTile struct {
	renderer tileRenderer
	bounds   engine.Rect
}

func NewBaseTile(renderer tileRenderer, bounds engine.Rect) BaseTile {
	return BaseTile{
		renderer: renderer,
		bounds:   bounds,
	}
}

func (t BaseTile) SpriteAtLayer(layer int) (engine.Sprite, bool) {
	sprites := t.renderer.Sprites()
	if layer > len(sprites)-1 {
		return engine.Sprite{}, false
	}

	return sprites[layer], true
}

func (t BaseTile) Bounds() engine.Rect {
	return t.bounds
}

type tileRenderer interface {
	Sprites() []engine.Sprite
}
