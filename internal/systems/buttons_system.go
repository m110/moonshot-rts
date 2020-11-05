package systems

import (
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type buttonsEntity interface {
	engine.Entity
	components.DrawableOwner
	components.ButtonOwner
}

type ButtonsSystem struct {
	base     BaseSystem
	entities EntityMap

	activeButton buttonsEntity
}

func NewButtonsSystem(config Config, eventBus *engine.EventBus, spawner spawner) *ButtonsSystem {
	return &ButtonsSystem{
		base: NewBaseSystem(config, eventBus, spawner),
	}
}

func (b *ButtonsSystem) Start() {
	b.base.EventBus.Subscribe(EntityClicked{}, b)
	b.base.EventBus.Subscribe(ClickReleased{}, b)
}

func (b *ButtonsSystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case EntityClicked:
		_, ok := b.entities.ByID(event.Entity.ID())
		if !ok {
			return
		}

		entity := event.Entity.(buttonsEntity)
		entity.GetButton().Pressed = true
		b.updateSprite(entity)
		entity.GetButton().Action()
		b.activeButton = entity
	case ClickReleased:
		if b.activeButton == nil {
			return
		}

		b.activeButton.GetButton().Pressed = false
		b.updateSprite(b.activeButton)
		b.activeButton = nil
	}
}

func (b ButtonsSystem) Update(dt float64) {}

func (b ButtonsSystem) Draw(canvas engine.Sprite) {}

func (b *ButtonsSystem) Add(entity buttonsEntity) {
	b.updateSprite(entity)
	b.entities.Add(entity)
}

func (b ButtonsSystem) Remove(entity engine.Entity) {
	b.entities.Remove(entity)
}

func (b *ButtonsSystem) updateSprite(entity buttonsEntity) {
	button := entity.GetButton()

	var baseSprite engine.Sprite
	if button.Pressed {
		baseSprite = button.SpritePressed
	} else {
		baseSprite = button.SpriteReleased
	}

	sprite := engine.NewBlankSprite(baseSprite.Size())
	sprite.Draw(baseSprite)
	sprite.DrawAtPosition(button.SpriteTop, baseSprite.Width()/2, baseSprite.Height()/2)

	entity.GetDrawable().Sprite = sprite
}
