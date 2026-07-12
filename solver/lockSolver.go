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
	//lck := lock.CreateValidationLock6()
	lck := lock.CreateHardLock6()
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

	solvingMoves, err := solve(startingPositions, moves)
	if err != nil {
		return
	}

	fmt.Println("Lock is solved - required moves: ")
	fmt.Println(solvingMoves)
	fmt.Println("So that would be:")
	latchCount := len(startingPositions)

	builder := strings.Builder{}
	currentPosition := startingPositions
	fmt.Printf("Starting position: %d \n", currentPosition)
	for _, mv := range solvingMoves {
		idOfMove := indexOf(mv, moves)
		currentPosition, err = move(currentPosition, mv)
		if err != nil {
			continue
		}
		if idOfMove > latchCount-1 {
			builder.WriteString(strconv.Itoa(idOfMove-latchCount+1))
			builder.WriteString("P ")
		} else {
			builder.WriteString(strconv.Itoa(idOfMove+1))
			builder.WriteString("L ")
		}
	}
	solutionString := builder.String()

	fmt.Printf("Final solution combination of %d moves:\n", len(solvingMoves))
	fmt.Println(compressSolutionMoves(solutionString))
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
		movesSlice := []int{0, 0, 0, 0, 0, 0}
		for _, e := range l.Effects {
			movesSlice[lck.IndexOf(e.Target)] = e.Move
		}
		movesSlice[i] = 1
		fmt.Println(movesSlice)
		moves = append(moves, movesSlice[:])
	}
	return moves
}

func addNegativeMoves(moves [][]int) [][]int {
	negativeMoves := [][]int{}
	for i, mv := range moves {
		negativeMoves = append(negativeMoves, []int{0, 0, 0, 0, 0, 0})
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
	visited := make(map[string]bool)
	solution := [][]int{}
	positionsAfterMove := [][]int{}
	visited[serializeToString(startingPosition)] = true
	positionsAfterMove = append(positionsAfterMove, startingPosition)
	for {
		resultPositions := [][]int{}
		for j := range positionsAfterMove {
			for i := range moves {
				resultingPosition, err := move(positionsAfterMove[j], moves[i])
				if err != nil {
					continue
				}
				if visited[serializeToString(resultingPosition)] {
					continue
				}

				visited[serializeToString(resultingPosition)] = true

				//m[serializeToString(resultingPosition)] = append(m[serializeToString(positionsAfterMove[j])], moves[i])
				currentKey := serializeToString(positionsAfterMove[j])
				resultingKey := serializeToString(resultingPosition)
				parentPath := m[currentKey]
				path := make([][]int, len(parentPath), len(parentPath)+1)
				copy(path, parentPath)
				path = append(path, moves[i])
				m[resultingKey] = path

				if isSolved(resultingPosition) {
					fmt.Println("Lock is solved!")
					fmt.Println(resultingPosition)
					solution = m[serializeToString(resultingPosition)]
					return solution, nil
				}

				resultPositions = append(resultPositions, resultingPosition)
			}
		}
		positionsAfterMove = resultPositions
	}
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
		if slice[i] != 0 {
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

func compressSolutionMoves(solution string) string {
	moves := strings.Fields(solution)
	if len(moves) == 0 {
		return ""
	}

	compressed := []string{}
	current := moves[0]
	count := 1

	for _, move := range moves[1:] {
		if move == current {
			count++
			continue
		}

		if count == 1 {
			compressed = append(compressed, current)
		} else {
			compressed = append(compressed, fmt.Sprintf("%s %dx", current, count))
		}

		current = move
		count = 1
	}

	if count == 1 {
		compressed = append(compressed, current)
	} else {
		compressed = append(compressed, fmt.Sprintf("%s %dx", current, count))
	}

	return strings.Join(compressed, "\n")
}
