package engine

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type PivotType int

const (
	PivotTop PivotType = iota
	PivotCenter
	PivotBottom
	PivotTopLeft
	PivotTopRight
	PivotBottomLeft
	PivotBottomRight
)

type Sprite struct {
	image *ebiten.Image
	pivot Vector
}

type Sprites []Sprite

func (s Sprites) Random() Sprite {
	return s[rand.Intn(len(s))]
}

func NewBlankSprite(width int, height int) Sprite {
	return Sprite{
		image: ebiten.NewImage(width, height),
	}
}

func NewFilledSprite(width int, height int, c color.Color) Sprite {
	image := ebiten.NewImage(width, height)
	image.Fill(c)

	return Sprite{
		image: image,
	}
}

func NewBlankSpriteWithPivot(width int, height int, pivot Vector) Sprite {
	return Sprite{
		image: ebiten.NewImage(width, height),
		pivot: pivot,
	}
}

func NewSpriteFromImage(image *ebiten.Image) Sprite {
	return Sprite{
		image: image,
	}
}

func NewSpriteFromImageWithPivotType(image *ebiten.Image, pivotType PivotType) Sprite {
	return Sprite{
		image: image,
		pivot: NewPivotForImage(image, pivotType),
	}
}

func NewSpriteFromImageWithPivot(image *ebiten.Image, pivot Vector) Sprite {
	return Sprite{
		image: image,
		pivot: pivot,
	}
}

func NewSpriteFromSprite(s Sprite) Sprite {
	return Sprite{
		image: ebiten.NewImageFromImage(s.image),
		pivot: s.Pivot(),
	}
}

func (s Sprite) IsZero() bool {
	return s.image == nil
}

func (s Sprite) Image() *ebiten.Image {
	return s.image
}

func (s Sprite) Size() (width int, height int) {
	return s.image.Size()
}

func (s Sprite) Width() int {
	w, _ := s.image.Size()
	return w
}

func (s Sprite) Height() int {
	_, h := s.image.Size()
	return h
}

func (s *Sprite) SetPivot(p Vector) {
	s.pivot = p
}

func (s Sprite) Pivot() Vector {
	return s.pivot
}

func (s *Sprite) Scale(scale Vector) {
	w, h := s.image.Size()
	scaled := ebiten.NewImage(int(float64(w)*scale.X), int(float64(h)*scale.Y))

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale.Unpack())
	scaled.DrawImage(s.image, op)

	s.image = scaled
}

// Draw draws source sprite on the sprite.
func (s Sprite) Draw(source Sprite) {
	s.DrawAtPosition(source, 0, 0)
}

// DrawAtVector draws source sprite on the sprite at given vector.
func (s Sprite) DrawAtVector(source Sprite, v Vector) {
	s.DrawAtPosition(source, int(v.X), int(v.Y))
}

// DrawAtVector draws source sprite on the sprite at given point.
func (s Sprite) DrawAtPoint(source Sprite, p Point) {
	s.DrawAtPosition(source, p.X, p.Y)
}

// DrawAtPosition draws source sprite on the sprite at given position.
func (s Sprite) DrawAtPosition(source Sprite, x int, y int) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(
		float64(x)-source.Pivot().X,
		float64(y)-source.Pivot().Y,
	)

	s.image.DrawImage(source.image, op)
}

func (s Sprite) DrawInSection(source Sprite, rect Rect) {
	op := &ebiten.DrawImageOptions{}

	x := int(rect.Position.X)
	y := int(rect.Position.Y)
	width := x + int(rect.Width)
	height := y + int(rect.Height)

	imageRect := image.Rect(x, y, width, height)

	s.image.SubImage(imageRect).(*ebiten.Image).DrawImage(source.image, op)
}

func NewPivotForSprite(sprite Sprite, pivotType PivotType) Vector {
	return NewPivotForImage(sprite.image, pivotType)
}

func NewPivotForImage(image *ebiten.Image, pivotType PivotType) Vector {
	w, h := image.Size()
	wCenter := float64(w / 2)
	hCenter := float64(h / 2)

	switch pivotType {
	case PivotTop:
		return Vector{X: wCenter, Y: 0}
	case PivotCenter:
		return Vector{X: wCenter, Y: hCenter}
	case PivotBottom:
		return Vector{X: wCenter, Y: float64(h)}
	case PivotTopLeft:
		return Vector{X: 0, Y: 0}
	case PivotTopRight:
		return Vector{X: float64(w), Y: 0}
	case PivotBottomLeft:
		return Vector{X: 0, Y: float64(h)}
	case PivotBottomRight:
		return Vector{X: float64(w), Y: float64(h)}
	default:
		panic(fmt.Sprintf("unknown pivot: %v", pivotType))
	}
}
