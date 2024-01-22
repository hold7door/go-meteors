package game

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewReact(x, y, Width, Height float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  Width,
		Height: Height,
	}
}

func (r *Rect) MaxX() float64 {
	return r.X + r.Width
}

func (r *Rect) MaxY() float64 {
	return r.Y + r.Height
}

func (r Rect) Intersects(other Rect) bool {
	// TODO: Understand this
	return r.X <= other.MaxX() &&
		other.X <= r.MaxX() &&
		r.Y <= other.MaxY() &&
		other.Y <= r.MaxY()
}
