package main

import (
	"emergence-sim/sim"
	stream "emergence-sim/viz"
	"time"
)

func main() {
	N := 1000
	s := sim.NewSimulation(N, 100, 100)

	// snapshot buffers
	bufA := sim.NewSnapshot(N)
	bufB := sim.NewSnapshot(N)
	current := bufA
	next := bufB

	snapshotChan := make(chan *sim.Snapshot, 1)

	// run streamer
	go stream.Serve(":8080", snapshotChan)

	ticker := time.NewTicker(time.Millisecond * 16) // ~60 FPS sim
	for range ticker.C {
		s.Step(0.016)

		// every 10 steps, send snapshot
		s.WriteSnapshot(current)
		select {
		case snapshotChan <- current:
			current, next = next, current
		default:
			// drop frame if streamer is slow
		}
	}
}
