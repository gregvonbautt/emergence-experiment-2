package sim

type Simulation struct {
	Elements      []Element
	Width, Height float64
}

func NewSimulation(n int, w, h float64) *Simulation {
	elems := make([]Element, n)
	return &Simulation{Elements: elems, Width: w, Height: h}
}

func (s *Simulation) Step(dt float64) {
	for i := range s.Elements {
		e := &s.Elements[i]

		// simple physics / rules
		e.Position.X += e.Velocity.X * dt
		e.Position.Y += e.Velocity.Y * dt

		// wrap around world
		if e.Position.X < 0 {
			e.Position.X += s.Width
		}
		if e.Position.X > s.Width {
			e.Position.X -= s.Width
		}
		if e.Position.Y < 0 {
			e.Position.Y += s.Height
		}
		if e.Position.Y > s.Height {
			e.Position.Y -= s.Height
		}

		// example energy decay
		e.Energy *= 0.99
	}
}
