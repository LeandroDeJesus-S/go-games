package game

type Vector struct {
	x float64
	y float64
}


type Rect struct {
	Position Vector
	Width int
	Height int
}

func NewRect(Pos Vector, W, H int) Rect {
	return Rect{
		Position: Pos,
		Width: W,
		Height: H,
	}
}

func (r *Rect) CollidedWith(other Rect) bool {
	return r.Position.x <= other.MaxX() &&
		   other.Position.x <= r.MaxX() &&
		   r.Position.y <= other.MaxY() &&
		   other.Position.y <= r.MaxY()
}

func (r *Rect) MaxX() float64 {
	return r.Position.x + float64(r.Width)
}

func (r *Rect) MaxY() float64 {
	return r.Position.y + float64(r.Height)
}