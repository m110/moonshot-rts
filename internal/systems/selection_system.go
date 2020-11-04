package systems

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m110/rts/internal/components"
	"github.com/m110/rts/internal/engine"
	"github.com/m110/rts/internal/objects"
)

type selectionEntity interface {
	engine.Entity
	components.WorldSpaceOwner
	components.SelectableOwner
	components.BoxBoundaryOwner
}

type SelectionSystem struct {
	base BaseSystem

	entities EntityList

	selectedEntities []selectionEntity

	overlay        objects.Overlay
	overlayAnchor  engine.Vector
	overlayEnabled bool
}

type EntitySelected struct {
	Entity engine.Entity
}

type EntityUnselected struct {
	Entity engine.Entity
}

func NewSelectionSystem(config Config, eventBus *engine.EventBus, spawner spawner) *SelectionSystem {
	return &SelectionSystem{
		base: NewBaseSystem(config, eventBus, spawner),
	}
}

func (s *SelectionSystem) Start() {
	s.overlay = objects.NewOverlay(0, 0, engine.PivotTopLeft)
	s.overlay.GetDrawable().Disable()
	s.base.Spawner.SpawnDrawingEntity(s.overlay)
}

func (s *SelectionSystem) Update(dt float64) {
	cx, cy := ebiten.CursorPosition()
	cursorX := float64(cx)
	cursorY := float64(cy)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if s.overlayEnabled {
			s.hideOverlay()
			r := engine.Rect{
				Position: s.overlay.WorldSpace.WorldPosition(),
				Width:    float64(s.overlay.GetSize().Width),
				Height:   float64(s.overlay.GetSize().Height),
			}
			for _, e := range s.entities.All() {
				entity := e.(selectionEntity)
				bounds := entity.GetBoxBoundary().Bounds
				bounds.Position = bounds.Position.Add(entity.GetWorldSpace().WorldPosition())
				if bounds.Intersects(r) && entity.GetSelectable().GroupSelectable {
					entity.GetSelectable().Select()
					s.selectedEntities = append(s.selectedEntities, entity)

					s.base.EventBus.Publish(EntitySelected{
						Entity: entity,
					})
				} else {
					entity.GetSelectable().Unselect()
				}
			}
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if len(s.selectedEntities) > 0 {
			for _, e := range s.selectedEntities {
				e.GetSelectable().Unselect()
				s.base.EventBus.Publish(EntityUnselected{Entity: e})
			}
			s.selectedEntities = nil
		} else {
			v := engine.Vector{X: cursorX, Y: cursorY}
			for _, e := range s.entities.All() {
				entity := e.(selectionEntity)
				bounds := entity.GetBoxBoundary().Bounds
				bounds.Position = bounds.Position.Add(entity.GetWorldSpace().WorldPosition())
				if len(s.selectedEntities) == 0 && bounds.WithinBounds(v) {
					entity.GetSelectable().Select()
					s.selectedEntities = append(s.selectedEntities, entity)

					s.base.EventBus.Publish(EntitySelected{
						Entity: entity,
					})
				} else {
					entity.GetSelectable().Unselect()
				}
			}

			if len(s.selectedEntities) == 0 {
				s.showOverlay(cursorX, cursorY)
			}
		}
	}

	if s.overlayEnabled {
		s.updateOverlay(cursorX, cursorY)
	}
}

func (s *SelectionSystem) showOverlay(x float64, y float64) {
	s.overlayAnchor = engine.Vector{X: x, Y: y}
	s.overlay.GetWorldSpace().SetInWorld(x, y)
	s.overlay.GetDrawable().Enable()
	s.overlayEnabled = true
}

func (s *SelectionSystem) updateOverlay(x float64, y float64) {
	var pos engine.Vector

	switch {
	case x < s.overlayAnchor.X && y < s.overlayAnchor.Y:
		pos = engine.Vector{X: x, Y: y}
	case x < s.overlayAnchor.X && y > s.overlayAnchor.Y:
		pos = engine.Vector{X: x, Y: s.overlayAnchor.Y}
	case x > s.overlayAnchor.X && y < s.overlayAnchor.Y:
		pos = engine.Vector{X: s.overlayAnchor.X, Y: y}
	default:
		pos = engine.Vector{X: s.overlayAnchor.X, Y: s.overlayAnchor.Y}
	}

	s.overlay.WorldSpace.SetInWorld(pos.X, pos.Y)
	s.overlay.Size.Set(
		int(math.Abs(x-s.overlayAnchor.X)),
		int(math.Abs(y-s.overlayAnchor.Y)),
	)
	s.overlay.Drawable.Sprite = objects.NewRectangleSprite(s.overlay, engine.PivotTopLeft)
}

func (s *SelectionSystem) hideOverlay() {
	s.overlayEnabled = false
	s.overlay.Drawable.Disable()
}

func (s SelectionSystem) Draw(canvas engine.Sprite) {
}

func (s *SelectionSystem) Add(entity selectionEntity) {
	s.entities.Add(entity)
}

func (s SelectionSystem) Remove(entity engine.Entity) {
	s.entities.Remove(entity)
}
