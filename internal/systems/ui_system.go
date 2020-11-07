package systems

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/objects"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

type UIConfig struct {
	OffsetX int
	OffsetY int

	Width  int
	Height int

	Font font.Face
}

type UISystem struct {
	base BaseSystem

	resources objects.Object
	fps       objects.Object
}

func NewUISystem(config Config, eventBus *engine.EventBus, spawner spawner) *UISystem {
	return &UISystem{
		base: NewBaseSystem(config, eventBus, spawner),
	}
}

func (u *UISystem) Start() {
	u.base.EventBus.Subscribe(ResourcesUpdated{}, u)

	sprite := engine.NewFilledSprite(u.base.Config.UI.Width, u.base.Config.UI.Height, colornames.Black)
	ui := objects.NewObject(sprite, components.UILayerBackground)
	ui.Translate(
		float64(u.base.Config.UI.OffsetX),
		float64(u.base.Config.UI.OffsetY),
	)

	resourcesSprite := engine.NewBlankSprite(u.base.Config.UI.Width, u.base.Config.UI.Height)
	u.resources = objects.NewObject(resourcesSprite, components.UILayerText)
	ui.GetWorldSpace().AddChild(u.resources)
	u.resources.Translate(20, 20)

	fpsSprite := engine.NewBlankSprite(200, u.base.Config.UI.Height)
	u.fps = objects.NewObject(fpsSprite, components.UILayerText)
	ui.GetWorldSpace().AddChild(u.fps)
	u.fps.Translate(float64(u.base.Config.UI.Width-200), 20)

	u.base.Spawner.SpawnObject(ui)
}

func (u UISystem) updateResources(resources components.Resources) {
	content := fmt.Sprintf("Food: %v Wood: %v Stone: %v Gold: %v Iron: %v",
		resources.Food, resources.Wood, resources.Stone, resources.Gold, resources.Iron)

	u.resources.Sprite.Image().Fill(colornames.Black)
	bounds := text.BoundString(u.base.Config.UI.Font, content)
	text.Draw(
		u.resources.Sprite.Image(),
		content,
		u.base.Config.UI.Font,
		0, bounds.Dy(), colornames.White,
	)
}

func (u UISystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case ResourcesUpdated:
		u.updateResources(event.Resources)
	}
}

func (u UISystem) Update(_ float64) {
	content := fmt.Sprintf("FPS: %2.0f", ebiten.CurrentFPS())
	u.fps.Sprite.Image().Fill(colornames.Black)
	bounds := text.BoundString(u.base.Config.UI.Font, content)
	text.Draw(
		u.fps.Sprite.Image(),
		content,
		u.base.Config.UI.Font,
		0, bounds.Dy(), colornames.White,
	)
}

func (u UISystem) Draw(canvas engine.Sprite) {
}

func (u UISystem) Remove(entity engine.Entity) {
}
