package solver

import (
	"fmt"
	"gothicModSover/lock"
)

func RunSolver() error {
	lck, err := lock.CreateLockFromUserInput()
	if err != nil {
		return fmt.Errorf("Error occured during lock creation form user input: %w", err)
	}
	lck.PrintLock()
	lck.PrintEffects()

	solveLock(lck)

	return nil
}


func solveLock(lck *lock.Lock) {
 // lock positions array
	positions := getPositionsArray(lck)
	fmt.Println(positions)

	// array of moves array 
	moves := getMovesArray(lck)
	fmt.Println(moves)
	// great - now for each of them we can do left or right - meaning effects will be negated 0too
	// so now we gonna add or subtract - for the move right or left respectively - keeping the moves executed in memory 

}

func getPositionsArray(lck *lock.Lock) []int {
	positions := []int{}

	for _, l := range lck.Latches {
		positions = append(positions, l.Position)
	}

	return positions
}

func getMovesArray(lck *lock.Lock) [][]int {
	moves := [][]int{}
	for _, l := range lck.Latches {
		movesArray := [5]int{0,0,0,0,0}
		for _, e := range l.Effects {
			movesArray[lck.IndexOf(e.Target)] = e.Move
		}
		fmt.Println(movesArray)
		moves = append(moves, movesArray[:])
	}
	return moves
}

