package objects

import (
	"github.com/m110/moonshot-rts/internal/atlas"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type Panel struct {
	Object

	buttons []Button
}

func NewPanel(sprites []engine.Sprite) Panel {
	p := Panel{
		Object: NewObject(atlas.PanelBrown, components.LayerUI),
	}

	p.Drawable.Sprite.Scale(engine.Vector{X: 1.5, Y: 1.5})

	x := 20
	y := 15
	for i, s := range sprites {
		b := NewButton(s)
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

type Button struct {
	Object

	pressed bool

	spriteTop      engine.Sprite
	spriteReleased engine.Sprite
	spritePressed  engine.Sprite
}

func NewButton(spriteTop engine.Sprite) Button {
	b := Button{
		Object:         NewObject(engine.Sprite{}, components.LayerUIButton),
		spriteTop:      spriteTop,
		spriteReleased: atlas.ButtonBeige,
		spritePressed:  atlas.ButtonBeigePressed,
	}

	b.spriteTop.SetPivot(engine.NewPivotForSprite(b.spriteTop, engine.PivotCenter))
	b.updateSprite()

	return b
}

func (b *Button) Press() {
	b.pressed = true
	b.updateSprite()
}

func (b *Button) Release() {
	b.pressed = false
	b.updateSprite()
}

func (b *Button) updateSprite() {
	var baseSprite engine.Sprite
	if b.pressed {
		baseSprite = b.spritePressed
	} else {
		baseSprite = b.spriteReleased
	}

	sprite := engine.NewBlankSprite(baseSprite.Size())
	sprite.Draw(baseSprite)
	sprite.DrawAtPosition(b.spriteTop, baseSprite.Width()/2, baseSprite.Height()/2)

	b.Drawable.Sprite = sprite
}
