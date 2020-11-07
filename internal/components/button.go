package components

import (
	"github.com/m110/moonshot-rts/internal/engine"
)

type UIColor int

const (
	UIColorBeige UIColor = iota
	UIColorBrown
)

type Button struct {
	Action  func()
	Pressed bool

	SpriteTop      engine.Sprite
	SpriteReleased engine.Sprite
	SpritePressed  engine.Sprite
}
