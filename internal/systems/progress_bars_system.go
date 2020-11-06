package systems

import (
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
		entity.GetDrawable().Sprite = updateProgressBarSprite(entity.GetProgressBar())
	}
}

func updateProgressBarSprite(bar *components.ProgressBar) engine.Sprite {
	// TODO probably not a good idea to create the new image every frame
	midLength := 3

	background := fillProgressBar(bar.Background, midLength, 1.0)
	foreground := fillProgressBar(bar.Foreground, midLength, bar.Progress)

	background.Draw(foreground)
	background.SetPivot(engine.NewPivotForSprite(background, engine.PivotCenter))

	return background
}

func fillProgressBar(sprites components.ProgressBarSprites, midLength int, widthPercent float64) engine.Sprite {
	width := sprites.Left.Width() + midLength*sprites.Mid.Width() + sprites.Right.Width()
	height := sprites.Left.Width() + midLength*sprites.Mid.Width() + sprites.Right.Width()

	sprite := engine.NewBlankSprite(int(float64(width)*widthPercent), height)

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
	p.entities.Add(entity)
}

func (p *ProgressBarSystem) Remove(entity engine.Entity) {
	p.entities.Remove(entity)
}
