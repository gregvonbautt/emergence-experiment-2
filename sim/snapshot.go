package sim

type Snapshot struct {
	Positions []Vec2
	Energy    []float64
	Type      []int
}

// allocate buffers
func NewSnapshot(n int) *Snapshot {
	return &Snapshot{
		Positions: make([]Vec2, n),
		Energy:    make([]float64, n),
		Type:      make([]int, n),
	}
}

// copy simulation state into snapshot
func (s *Simulation) WriteSnapshot(buf *Snapshot) {
	for i, e := range s.Elements {
		buf.Positions[i] = e.Position
		buf.Energy[i] = e.Energy
		buf.Type[i] = e.Type
	}
}
