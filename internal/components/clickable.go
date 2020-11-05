package components

import "github.com/m110/moonshot-rts/internal/engine"

type ClickableOwner interface {
	GetClickable() *Clickable
}

type Clickable struct {
	// Bounds defines position relative to WorldSpace
	Bounds engine.Rect

	ByOverlay bool
}

func BoundsFromSprite(sprite engine.Sprite) engine.Rect {
	w, h := sprite.Size()
	pivot := sprite.Pivot()
	return engine.Rect{
		Position: engine.Vector{
			X: -pivot.X,
			Y: -pivot.Y,
		},
		Width:  float64(w),
		Height: float64(h),
	}
}
