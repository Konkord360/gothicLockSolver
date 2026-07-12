package solver

import (
	"fmt"
	"gothicModSover/lock"
	"strconv"
	"strings"
)

func RunSolver() error {
//	lck, err := lock.CreateLockFromUserInput()
//	if err != nil {
//		return fmt.Errorf("Error occured during lock creation form user input: %w", err)
//	}
	lck := lock.CreateValidationLock6()
	lck.PrintLock()
	lck.PrintEffects()

	solveLock(lck)

	return nil
}


func solveLock(lck *lock.Lock) {
 // lock positions array
	startingPositions := getPositionsArray(lck)
	fmt.Println(startingPositions)

	// array of moves array 
	moves := getMovesArray(lck)
	fmt.Println(moves)
	moves = addNegativeMoves(moves)
	fmt.Println(moves)

	fmt.Printf("Starting position: %d \n", startingPositions)
	currentPosition, err := move(startingPositions, moves[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Lock position after move: %d\n", moves[0])
	fmt.Println(currentPosition)


	solvingMoves, err := solve(startingPositions, moves) 
	if err != nil { return }

	fmt.Println("Lock is solved - required moves: ")
	fmt.Println(solvingMoves)
	fmt.Println("So that would be:" )
	latchCount := len(startingPositions)


	solutionString := ""
	for _, mv := range solvingMoves {
		idOfMove := indexOf(mv, moves)	
		fmt.Println("Id of the move")	
		fmt.Println(idOfMove)

		if idOfMove > latchCount -1 { 
			solutionString += strconv.Itoa(idOfMove - latchCount + 1) + "L"	
		} else {
			solutionString += strconv.Itoa(idOfMove + 1) + "P"	
		}
	}

	fmt.Println("Final solution combination:")
	fmt.Println(solutionString)
	// now we have array representation of a lock with an array of all the possible moves in each positions
	// so now we iterate over all the moves - keeping already done ones 
}

func indexOf(move []int, moves [][]int) int {
	for i, m := range moves {
		areSame := true
		for j := range m {
	//	fmt.Println(move, m)
			if move[j] != m[j] { 
				areSame = false
				continue
			}
		}
		if areSame {
			return i
		}
	}
	return -1
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
	for i, l := range lck.Latches {
		movesSlice := []int{0,0,0,0,0,0}
		for _, e := range l.Effects {
			movesSlice[lck.IndexOf(e.Target)] = e.Move
		}
		movesSlice[i] = 1
		fmt.Println(movesSlice)
		moves = append(moves, movesSlice[:])
	}
	return moves
}

func addNegativeMoves(moves [][]int) [][]int{
	negativeMoves := [][]int{}
	for i, mv := range moves {
		negativeMoves = append(negativeMoves, []int{0,0,0,0,0,0})
		for j, mv2 := range mv {
			negativeMoves[i][j] = -mv2
		}
	}
	moves = append(moves, negativeMoves...)
	return moves
}

func move(position []int, move []int) ([]int, error) {
	localPosittion := make([]int, len(position))
	copy(localPosittion, position)
	for i := range localPosittion {
		localPosittion[i] += move[i]
		if localPosittion[i] > 3 || localPosittion[i] < -3 {
			return localPosittion, fmt.Errorf("Invalid move")
		}
	}
	return localPosittion, nil
}

func solve(startingPosition []int, moves [][]int) ([][]int, error) {
	fmt.Printf("Starting to search for a solution of a lock: %d\n", startingPosition)

	m := make(map[string][][]int)

	fmt.Println(m)
	
	solution := [][]int{}
	moveCount := len(moves)
	fmt.Println(moveCount)

	//for {
	positionsAfterMove := [][]int{}
	positionsAfterMove = append(positionsAfterMove, startingPosition)
	for {
		resultPositions := [][]int{}
		for j := range positionsAfterMove {
			for i := range moves {
				resultingPosition, err := move(positionsAfterMove[j], moves[i])
				m[serializeToString(resultingPosition)] = append(m[serializeToString(positionsAfterMove[j])], moves[i])
				if err != nil { 
					fmt.Printf("Move %d is illegal in the %d position\n", moves[i], positionsAfterMove[j])
					continue
				}
				if isSolved(resultingPosition) {
					fmt.Println("Lock is solved!")
					fmt.Println(resultingPosition)
					solution = m[serializeToString(resultingPosition)]
					return solution, nil
				}

				resultPositions = append(resultPositions, resultingPosition)
			}
			fmt.Println("Legal positions after each first move:")
			fmt.Println(resultPositions)
		}
		positionsAfterMove = resultPositions
	}

	//	currentPosition, err := move(startingPosition, moves[0])
	//	if err != nil { break }
	// solution = append(solution, moves[0])

	//`if isSolved(currentPosition) { return solution, nil }

	//}
	 // return solution, nil
}

func serializeToString(pos []int) string {
	if len(pos) == 0 {
		return ""
	}
	var b strings.Builder
	for i, v := range pos {
		if i > 0 {
			b.WriteByte(',') // separator
		}
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()

}

func isSolved(slice []int) bool {
	for i := range slice {
		if slice[i] != 0 {return false}	
	}

	return true
}

func isOpposite(move1 []int, move2 []int) bool{
	for i := range move1 {
		if move1[i] != -move2[i] {
			return false
		}
	}
	return true
}

func isEmpty(slice []int) bool {
	for _, e := range slice {
		if e != 0 {
			return false
		}
	}
	return true
}

