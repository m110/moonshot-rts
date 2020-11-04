package systems

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/engine"
)

type DrawingEntity interface {
	engine.Entity
	components.WorldSpaceOwner
	components.DrawableOwner
}

type DrawingSystem struct {
	base BaseSystem

	entities []EntityList

	canvas engine.Sprite
	layers [][]spriteAtPosition
}

type spriteAtPosition struct {
	Sprite   engine.Sprite
	Position engine.Vector
}

func NewDrawingSystem(config Config, eventBus *engine.EventBus, spawner spawner) *DrawingSystem {
	canvas := engine.NewSpriteFromImage(ebiten.NewImage(config.WindowSize()))

	layers := make([][]spriteAtPosition, components.AllLayers)

	return &DrawingSystem{
		base:     NewBaseSystem(config, eventBus, spawner),
		entities: make([]EntityList, components.AllLayers),
		canvas:   canvas,
		layers:   layers,
	}
}

func (d *DrawingSystem) Add(e DrawingEntity) {
	d.entities[e.GetDrawable().Layer].Add(e)
}

func (d *DrawingSystem) Remove(e engine.Entity) {
	for i := range d.entities {
		d.entities[i].Remove(e)
	}
}

func (d *DrawingSystem) Start() {
}

func (d *DrawingSystem) Update(_ float64) {}

func (d DrawingSystem) Draw(screen engine.Sprite) {
	for i := range d.layers {
		// TODO this is probably not really performant
		d.layers[i] = nil
	}

	for _, l := range d.entities {
		for _, e := range l.All() {
			de := e.(DrawingEntity)
			d.drawEntity(de)
		}
	}

	for _, l := range d.layers {
		// Sort the sprites in a layer, so the objects below cover the objects "behind" them
		sort.Slice(l, func(i, j int) bool {
			return l[i].Position.Y < l[j].Position.Y
		})
		for _, s := range l {
			d.canvas.DrawAtVector(s.Sprite, s.Position)
		}
	}

	screen.Draw(d.canvas)
}

func (d DrawingSystem) drawEntity(e DrawingEntity) {
	if e.GetDrawable().Disabled {
		return
	}

	drawable := e.GetDrawable()
	d.layers[drawable.Layer] = append(d.layers[drawable.Layer], spriteAtPosition{
		Sprite:   drawable.Sprite,
		Position: e.GetWorldSpace().WorldPosition(),
	})
}
