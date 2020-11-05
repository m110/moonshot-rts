package objects

import (
	"github.com/m110/moonshot-rts/internal/atlas"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type Panel struct {
	Object
	*components.Clickable

	buttons []PanelButton
}

type ButtonConfig struct {
	Sprite engine.Sprite
	Action func()
}

func NewFourButtonPanel(buttonConfigs []ButtonConfig) Panel {
	p := Panel{
		Object: NewObject(atlas.PanelBrown, components.LayerUIPanel),
	}

	p.Drawable.Sprite.Scale(engine.Vector{X: 1.5, Y: 1.5})
	p.Clickable = &components.Clickable{
		Bounds: components.BoundsFromSprite(p.Drawable.Sprite),
	}

	x := 20
	y := 15
	for i, s := range buttonConfigs {
		b := NewPanelButton(components.UIColorBeige, s.Sprite, s.Action)
		p.buttons = append(p.buttons, b)
		p.WorldSpace.AddChild(p, b)
		b.WorldSpace.Translate(float64(x), float64(y))

		if i%2 != 0 {
			y += 60
			x = 20
		} else {
			x += 60
		}
	}

	return p
}

func (p Panel) GetClickable() *components.Clickable {
	return p.Clickable
}

type PanelButton struct {
	Object
	*components.Clickable
	*components.Button
}

func NewPanelButton(color components.UIColor, spriteTop engine.Sprite, action func()) PanelButton {
	spriteTop.SetPivot(engine.NewPivotForSprite(spriteTop, engine.PivotCenter))

	var spriteReleased, spritePressed engine.Sprite
	switch color {
	case components.UIColorBeige:
		spriteReleased = atlas.ButtonBeige
		spritePressed = atlas.ButtonBeigePressed
	case components.UIColorBrown:
		spriteReleased = atlas.ButtonBrown
		spritePressed = atlas.ButtonBrownPressed
	}

	b := PanelButton{
		Object: NewObject(engine.Sprite{}, components.LayerUIButton),
		Button: &components.Button{
			Action:         action,
			Pressed:        false,
			SpriteTop:      spriteTop,
			SpriteReleased: spriteReleased,
			SpritePressed:  spritePressed,
		},
	}

	b.Clickable = &components.Clickable{
		Bounds: components.BoundsFromSprite(b.Button.SpriteReleased),
	}

	return b
}

func (b PanelButton) GetClickable() *components.Clickable {
	return b.Clickable
}

func (b PanelButton) GetButton() *components.Button {
	return b.Button
}