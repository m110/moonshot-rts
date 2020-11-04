package engine

type Point struct {
	X int
	Y int
}

func (p Point) Unpack() (x int, y int) {
	return p.X, p.Y
}

func (p *Point) Set(x int, y int) {
	p.X = x
	p.Y = y
}

func (p *Point) Translate(dx int, dy int) {
	p.X += dx
	p.Y += dy
}
