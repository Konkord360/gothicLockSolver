package main

import (
	"fmt"
	"os"
)

const (
	opposite int = -1
	same int = 1
)

type Lock struct {
	latches[] *Latch
}

type Latch struct{
	position int // latch is in a position 1 -7 with 4 being the middle
	effects[] Effect // latch hes effects - can move other latches
}

type Effect struct {
	targets[] *Latch // effect affects latches
	moves[] int // moves to execute on latches - 1 to 1 mapping moves to targets
}

func (l *Latch) move(m int) {
	l.position += m

	for _, e := range l.effects {
		e.apply(m)
	}
}

func (l *Latch) moveWithoutEffects(m int) {
	l.position += m
}

func (e Effect) apply(originalMove int) {
	moves := e.moves	
	targets := e.targets

	for i, target := range targets {
		target.moveWithoutEffects(moves[i]*originalMove)
	}
}

func (l Lock) printLock() {
	fmt.Println("***********")
	for _, e := range l.latches {
		fmt.Println(e.position)

	}
	fmt.Println("***********")
}


func main() {
	cmdArgs := os.Args[:]

	if len(cmdArgs) == 1 {
		fmt.Println("No cmd args provided")
	} else {
		fmt.Printf("Provided argument: %s\n", cmdArgs[1])
	} 

	if cmdArgs[1] == "simulation" {
		fmt.Println("Running lock pick simulation")
		return
	}

	latch1 := Latch{position: 3, effects: []Effect{}} 
	latch2 := Latch{position: 3, effects: []Effect{}} 
	latch3 := Latch{position: 4, effects: []Effect{}} 

	latch1.effects = []Effect{{
		targets: []*Latch{&latch2, &latch3},
		moves: []int{same, opposite},
	}}

	latch2.effects = []Effect{{
		targets: []*Latch{&latch1, &latch3},
		moves: []int{opposite, same},
	}}
	
	lock := Lock{[]*Latch{&latch1, &latch2, &latch3}}

	lock.printLock()

	lock.latches[0].move(-1)

	lock.printLock()

	lock.latches[1].move(-1)

	lock.printLock()

}
