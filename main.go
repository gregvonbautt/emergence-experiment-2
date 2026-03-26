package main

import (
	"emergence-sim/sim"
	stream "emergence-sim/viz"
	"time"
)

func main() {
	N := 10000
	s := sim.NewSimulation(N, 100, 100)

	// snapshot buffers
	bufA := sim.NewSnapshot(N)
	bufB := sim.NewSnapshot(N)
	current := bufA
	next := bufB

	snapshotChan := make(chan *sim.Snapshot, 1)

	// run streamer
	go stream.Serve(":8080", snapshotChan)

	const delayMs = 10
	const snapshotPeriod = 5

	cnt := 0

	for {
		s.Step(delayMs / 1000.0)

		// write snapshot if needed
		cnt++
		if cnt >= snapshotPeriod {
			cnt = 0
			s.WriteSnapshot(current)
			select {
			case snapshotChan <- current:
				current, next = next, current
			default:
				// drop frame if streamer is slow
			}
		}

		time.Sleep(delayMs * time.Millisecond)
	}
}
