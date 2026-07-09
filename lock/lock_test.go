package lock

import (
	"testing"
)

func TestLatchEffectsShouldAplyCorrectly(t *testing.T) {
	lock := createExampleLock()
	lock.latches[0].Move(Left)

	pos1, pos2, pos3 := lock.latches[0].position, lock.latches[1].position, lock.latches[2].position
	if pos1 != 2 || pos2 != 2 || pos3 != 5 {
		t.Fatalf("Lock positions should be 2, 2, 5, but was %d, %d, %d", pos1, pos2, pos3)
	}
	lock.latches[1].Move(Right)

	pos1, pos2, pos3 = lock.latches[0].position, lock.latches[1].position, lock.latches[2].position
	if pos1 != 1 || pos2 != 3 || pos3 != 6 {
		t.Fatalf("Lock positions should be 2, 2, 5, but was %d, %d, %d", pos1, pos2, pos3)
	}
}

func createExampleLock() Lock {
	latch1 := Latch{position: 3, effects: []Effect{}}
	latch2 := Latch{position: 3, effects: []Effect{}}
	latch3 := Latch{position: 4, effects: []Effect{}}

	latch1.effects = []Effect{{
		targets: []*Latch{&latch2, &latch3},
		moves:   []int{same, opposite},
	}}

	latch2.effects = []Effect{{
		targets: []*Latch{&latch1, &latch3},
		moves:   []int{opposite, same},
	}}

	return Lock{[]*Latch{&latch1, &latch2, &latch3}}
}
