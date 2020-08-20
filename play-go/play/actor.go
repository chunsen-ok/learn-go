package play

// Actor
type Actor interface {
	Move(dx float64, dy float64, dz float64) Pos
	MoveTo(x float64, y float64, z float64) Pos
	Speak() string
}
