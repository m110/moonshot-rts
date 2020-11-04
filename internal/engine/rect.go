package engine

type Rect struct {
	Position Vector
	Width    float64
	Height   float64
}

func (r Rect) WithinBounds(v Vector) bool {
	return v.X > r.Position.X && v.Y > r.Position.Y && v.X < r.Position.X+r.Width && v.Y < r.Position.Y+r.Height
}

func (r Rect) Intersects(other Rect) bool {
	return !(r.Position.X > other.Position.X+other.Width ||
		r.Position.X+r.Width < other.Position.X ||
		r.Position.Y > other.Position.Y+other.Height ||
		r.Position.Y+r.Height < other.Position.Y)
}
