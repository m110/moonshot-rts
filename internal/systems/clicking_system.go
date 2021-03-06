package systems

import (
	"math"

	"github.com/m110/moonshot-rts/internal/archetypes"
	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type clickingEntity interface {
	engine.Entity
	components.WorldSpaceOwner
	components.DrawableOwner
	components.ClickableOwner
}

type EntityClicked struct {
	Entity engine.Entity
}

type EntitiesClicked struct {
	Entities []engine.Entity
}

type PointClicked struct {
	Point engine.Vector
}

type ClickReleased struct {
}

type ClickingSystem struct {
	BaseSystem

	entities []EntityList

	overlay        archetypes.Overlay
	overlayAnchor  engine.Vector
	overlayEnabled bool
}

func NewClickingSystem(base BaseSystem) *ClickingSystem {
	return &ClickingSystem{
		BaseSystem: base,
		entities:   make([]EntityList, components.AllLayers),
	}
}

func (c *ClickingSystem) Start() {
	c.overlay = archetypes.NewOverlay(1, 1, engine.PivotTopLeft, colornames.White)
	c.overlay.GetDrawable().Disable()
	c.Spawner.Spawn(c.overlay)
}

func (c *ClickingSystem) Update(dt float64) {
	cx, cy := ebiten.CursorPosition()
	position := engine.Vector{X: float64(cx), Y: float64(cy)}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		c.EventBus.Publish(ClickReleased{})

		if c.overlayEnabled {
			c.hideOverlay()
			entities := c.findAllEntitiesInOverlay()
			if len(entities) > 0 {
				c.EventBus.Publish(EntitiesClicked{
					Entities: entities,
				})
			} else {
				// TODO should have a dedicated event?
				c.EventBus.Publish(PointClicked{
					Point: position,
				})
			}
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		entity := c.findFirstClickedEntity(position)
		if entity == nil {
			c.EventBus.Publish(PointClicked{
				Point: position,
			})

			c.showOverlay(position)
		} else {
			c.EventBus.Publish(EntityClicked{
				Entity: entity,
			})
		}
	}

	if c.overlayEnabled {
		c.updateOverlay(position)
	}
}

func (c ClickingSystem) findFirstClickedEntity(position engine.Vector) engine.Entity {
	for i := len(c.entities) - 1; i >= 0; i-- {
		l := c.entities[i]
		for _, e := range l.All() {
			entity := e.(clickingEntity)
			bounds := entity.GetClickable().Bounds
			bounds.Position = bounds.Position.Add(entity.GetWorldSpace().WorldPosition())

			if bounds.WithinBounds(position) {
				return e
			}
		}
	}

	return nil
}

func (c ClickingSystem) findAllEntitiesInOverlay() []engine.Entity {
	var entities []engine.Entity

	r := engine.Rect{
		Position: c.overlay.WorldSpace.WorldPosition(),
		Width:    float64(c.overlay.GetSize().Width),
		Height:   float64(c.overlay.GetSize().Height),
	}

	for _, l := range c.entities {
		for _, e := range l.All() {
			entity := e.(clickingEntity)
			clickable := entity.GetClickable()
			if clickable.Disabled || !clickable.ByOverlay {
				continue
			}

			bounds := entity.GetClickable().Bounds
			bounds.Position = bounds.Position.Add(entity.GetWorldSpace().WorldPosition())

			if bounds.Intersects(r) {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

func (c *ClickingSystem) showOverlay(cursor engine.Vector) {
	c.overlayAnchor = cursor
	c.overlay.GetWorldSpace().SetInWorld(cursor.X, cursor.Y)
	c.overlay.GetDrawable().Enable()
	c.overlayEnabled = true
}

func (c *ClickingSystem) updateOverlay(cursor engine.Vector) {
	var pos engine.Vector

	switch {
	case cursor.X < c.overlayAnchor.X && cursor.Y < c.overlayAnchor.Y:
		pos = engine.Vector{X: cursor.X, Y: cursor.Y}
	case cursor.X < c.overlayAnchor.X && cursor.Y > c.overlayAnchor.Y:
		pos = engine.Vector{X: cursor.X, Y: c.overlayAnchor.Y}
	case cursor.X > c.overlayAnchor.X && cursor.Y < c.overlayAnchor.Y:
		pos = engine.Vector{X: c.overlayAnchor.X, Y: cursor.Y}
	default:
		pos = engine.Vector{X: c.overlayAnchor.X, Y: c.overlayAnchor.Y}
	}

	c.overlay.WorldSpace.SetInWorld(pos.X, pos.Y)
	c.overlay.Size.Set(
		int(math.Abs(cursor.X-c.overlayAnchor.X)),
		int(math.Abs(cursor.Y-c.overlayAnchor.Y)),
	)

	if c.overlay.Size.Width == 0 {
		c.overlay.Size.Width = 1
	}

	if c.overlay.Size.Height == 0 {
		c.overlay.Size.Height = 1
	}

	c.overlay.Drawable.Sprite = archetypes.NewRectangleSprite(c.overlay, engine.PivotTopLeft, colornames.White)
}

func (c *ClickingSystem) hideOverlay() {
	c.overlayEnabled = false
	c.overlay.Drawable.Disable()
}

func (c *ClickingSystem) Add(entity clickingEntity) {
	c.entities[entity.GetDrawable().Layer].Add(entity)
}

func (c *ClickingSystem) Remove(entity engine.Entity) {
	for i := range c.entities {
		c.entities[i].Remove(entity)
	}
}
