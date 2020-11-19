package systems

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/m110/moonshot-rts/internal/archetypes"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
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
	BaseSystem

	resources archetypes.Object
	fps       archetypes.Object
}

func NewUISystem(base BaseSystem) *UISystem {
	return &UISystem{
		BaseSystem: base,
	}
}

func (u *UISystem) Start() {
	u.EventBus.Subscribe(ResourcesUpdated{}, u)

	sprite := engine.NewFilledSprite(u.Config.UI.Width, u.Config.UI.Height, colornames.Black)
	ui := archetypes.NewObject(sprite, components.UILayerBackground)
	ui.Translate(
		float64(u.Config.UI.OffsetX),
		float64(u.Config.UI.OffsetY),
	)

	resourcesSprite := engine.NewBlankSprite(u.Config.UI.Width, u.Config.UI.Height)
	u.resources = archetypes.NewObject(resourcesSprite, components.UILayerText)
	ui.GetWorldSpace().AddChild(u.resources)
	u.resources.Translate(20, 20)

	fpsSprite := engine.NewBlankSprite(200, u.Config.UI.Height)
	u.fps = archetypes.NewObject(fpsSprite, components.UILayerText)
	ui.GetWorldSpace().AddChild(u.fps)
	u.fps.Translate(float64(u.Config.UI.Width-200), 20)

	u.Spawner.Spawn(ui)
}

func (u UISystem) updateResources(used components.Resources, available components.Resources) {
	content := fmt.Sprintf("Food: %v/%v Wood: %v/%v Stone: %v/%v Gold: %v/%v Iron: %v/%v",
		used.Food, available.Food,
		used.Wood, available.Wood,
		used.Stone, available.Stone,
		used.Gold, available.Gold,
		used.Iron, available.Iron,
	)

	u.resources.Sprite.Image().Fill(colornames.Black)
	bounds := text.BoundString(u.Config.UI.Font, content)
	text.Draw(
		u.resources.Sprite.Image(),
		content,
		u.Config.UI.Font,
		0, bounds.Dy(), colornames.White,
	)
}

func (u UISystem) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case ResourcesUpdated:
		u.updateResources(event.UsedResources, event.AvailableResources)
	}
}

func (u UISystem) Update(_ float64) {
	content := fmt.Sprintf("FPS: %2.0f", ebiten.CurrentFPS())
	u.fps.Sprite.Image().Fill(colornames.Black)
	bounds := text.BoundString(u.Config.UI.Font, content)
	text.Draw(
		u.fps.Sprite.Image(),
		content,
		u.Config.UI.Font,
		0, bounds.Dy(), colornames.White,
	)
}

func (u UISystem) Remove(entity engine.Entity) {}
