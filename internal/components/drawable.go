package components

import "github.com/m110/moonshot-rts/internal/engine"

const AllLayers = 8

type DrawingLayer int

const (
	LayerBackground DrawingLayer = iota
	LayerGround
	LayerObjects
	LayerUnits
	LayerForeground
	LayerUI
	LayerUIPanel
	LayerUIButton
)

const (
	UILayerBackground DrawingLayer = iota
	UILayerText
)

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
