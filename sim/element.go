package sim

type Vec2 struct{ X, Y float64 }

type Element struct {
	Position Vec2
	Velocity Vec2
	Energy   float64
	Type     int
}
