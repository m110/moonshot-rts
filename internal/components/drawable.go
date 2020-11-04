package components

import "github.com/m110/moonshot-rts/internal/engine"

const AllLayers = 7

type DrawingLayer int

const (
	LayerBackground DrawingLayer = iota
	LayerGround
	LayerObjects
	LayerUnits
	LayerForeground
	LayerUI
	LayerUIButton
)

const (
	UILayerBackground DrawingLayer = iota
	UILayerText
)

// TODO find better suffix than Owner
type DrawableOwner interface {
	GetDrawable() *Drawable
}

type Drawable struct {
	Sprite engine.Sprite
	Layer  DrawingLayer

	Disabled bool
}

func (d *Drawable) Enable() {
	d.Disabled = false
}

func (d *Drawable) Disable() {
	d.Disabled = true
}
