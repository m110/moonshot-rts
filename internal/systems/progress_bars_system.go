package systems

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type progressBarEntity interface {
	engine.Entity
	components.DrawableOwner
	components.ProgressBarOwner
}

type ProgressBarSystem struct {
	base     BaseSystem
	entities EntityList
}

func NewProgressBarSystem(config Config, eventBus *engine.EventBus, spawner spawner) *ProgressBarSystem {
	return &ProgressBarSystem{
		base: NewBaseSystem(config, eventBus, spawner),
	}
}

func (p ProgressBarSystem) Start() {
}

func (p ProgressBarSystem) Update(dt float64) {
	for _, e := range p.entities.All() {
		entity := e.(progressBarEntity)
		updateProgressBarSprite(entity)
	}
}

func updateProgressBarSprite(entity progressBarEntity) {
	bar := entity.GetProgressBar()

	entity.GetDrawable().Sprite.Image().Clear()
	entity.GetDrawable().Sprite.Draw(bar.Background.Full)

	rect := image.Rect(0, 0, int(float64(bar.Background.Full.Width())*bar.Progress), bar.Background.Full.Height())
	foreground := entity.GetProgressBar().Foreground.Full.Image().SubImage(rect)
	entity.GetDrawable().Sprite.Image().DrawImage(ebiten.NewImageFromImage(foreground), nil)
}

func fillProgressBar(sprites components.ProgressBarSprites, midLength int) engine.Sprite {
	width := sprites.Left.Width() + midLength*sprites.Mid.Width() + sprites.Right.Width()
	height := sprites.Left.Width() + midLength*sprites.Mid.Width() + sprites.Right.Width()

	sprite := engine.NewBlankSprite(width, height)

	sprite.DrawAtPosition(sprites.Left, 0, 0)
	x := sprites.Left.Width()
	for i := 0; i < midLength; i++ {
		sprite.DrawAtPosition(sprites.Mid, x, 0)
		x += sprites.Mid.Width()
	}
	sprite.DrawAtPosition(sprites.Right, x, 0)
	return sprite
}

func (p ProgressBarSystem) Draw(canvas engine.Sprite) {}

func (p *ProgressBarSystem) Add(entity progressBarEntity) {
	midLength := 3

	bar := entity.GetProgressBar()
	bar.Background.Full = fillProgressBar(bar.Background, midLength)
	bar.Foreground.Full = fillProgressBar(bar.Foreground, midLength)

	entity.GetDrawable().Sprite = engine.NewBlankSprite(bar.Background.Full.Size())
	entity.GetDrawable().Sprite.SetPivot(engine.NewPivotForSprite(entity.GetDrawable().Sprite, engine.PivotCenter))

	p.entities.Add(entity)
}

func (p *ProgressBarSystem) Remove(entity engine.Entity) {
	p.entities.Remove(entity)
}
