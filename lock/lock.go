package lock

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	opposite int = -1
	same     int = 1

	Left  int = -1
	Right int = 1
)

type Lock struct {
	Latches []*Latch
}

type Latch struct {
	Position int      // latch is in a position -3 to 3 with 0 being the middle
	Effects  []Effect // latch has singular effect
}

type Effect struct {
	Target *Latch // effect affects 0 to n latches
	Move   int    // moves to execute on latches - 1 to 1 mapping moves to targets - same or opposite direction
}

func (l *Latch) Move(m int) error {
	l.Position += m
	if l.Position > 3 || l.Position < -3 {
		return fmt.Errorf("Illegal move - out of bounds of the lock")	
	}

	for _, e := range l.Effects {
		e.apply(m)
	}
	return nil
}

func (e Effect) apply(originalMove int) {
	e.Target.moveWithoutEffects(e.Move * originalMove)
}

func (l *Latch) moveWithoutEffects(m int) error {
	l.Position += m
	if l.Position > 3 || l.Position < -3 {
		return fmt.Errorf("Illegal move - out of bounds of the lock")	
	}
	return nil
}

func (lock Lock) PrintLock() {
	fmt.Println("***********")
	for _, l := range lock.Latches {
		fmt.Println(l.Position)

	}
	fmt.Println("***********")
}

func (lock Lock) PrintEffects() {
	fmt.Println("***********")
	for i, l := range lock.Latches {
		for _, e := range l.Effects {
			fmt.Printf("Latch %d moves latch %d in %d \n", i, lock.indexOf(e.Target), e.Move)
		}
	}
	fmt.Println("***********")
}

func (lock Lock) indexOf(latch *Latch) int {
	for i, lt := range lock.Latches {
		if lt == latch {
			return i
		}
	}
	return 0
}

// dunno about using scanners in function inputs - we'll see how it goes int testing - would be great if it was easily
// changable for an actual tui input later  - probably rewrite to gather user input first - then act on it - do not mix like that
// make it any better plz mate - i cant read it
func CreateLockFromUserInput() (*Lock, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input lock configuration:")
	lock := &Lock{}

	err := getLockLatchConfiguration(lock, scanner)
	if err != nil {
		return nil, fmt.Errorf("Error during latch configuration: %w", err)
	}

	err = assingEffectsToLatches(lock, scanner)
	if err != nil {
		return nil, fmt.Errorf("Error during effect assignment: %w", err)
	}

	err = assignStartingPositionsToLatches(lock, scanner)
	if err != nil {
		return nil, fmt.Errorf("Error starting positions assignment: %w", err)
	}

	return lock, nil
}

func getLockLatchConfiguration(lock *Lock, scanner *bufio.Scanner) error {
	fmt.Print("Number of latches: ")
	scanner.Scan()
	latchesText := scanner.Text()
	latches, err := strconv.Atoi(latchesText)
	if err != nil {
		return fmt.Errorf("Latches must be a number, %w", err)
	}

	if latches <= 2 {
		return fmt.Errorf("There must be more then two latches")
	}

	latchesArray := []*Latch{}
	for range latches {
		latchesArray = append(latchesArray, &Latch{})
	}

	lock.Latches = latchesArray
	return nil
}

func assingEffectsToLatches(lock *Lock, scanner *bufio.Scanner) error {
	fmt.Println("Assing an effect to each of the latches: ")
	for i := range lock.Latches {
		numberOfAffectedLatches, err := readNumberOfEffectsFromUser(i, scanner)
		if err != nil {
			fmt.Println("Error during effects assignment: %w", err)
		}
		for j := range numberOfAffectedLatches {
			targetLatch, moveDir, err := getEffectConfigurationFromUser(j, scanner)
			if err != nil {
				return fmt.Errorf("Error during effect configuration: %w", err)
			}
			lock.Latches[i].Effects = append(lock.Latches[i].Effects, Effect{Target: lock.Latches[targetLatch], Move: moveDir})
		}
	}
	return nil
}

func readNumberOfEffectsFromUser(latch int, scanner *bufio.Scanner) (int, error) {
	fmt.Printf("Latch %d - number of effects: ", latch)
	scanner.Scan()
	effectsText := scanner.Text()
	effects, err := strconv.Atoi(effectsText)
	if err != nil {
		return 0, fmt.Errorf("effects must be a number, %w", err)
	}
	if effects < 0 {
		return 0, fmt.Errorf("effects must be non negative number")
	}
	return effects, nil
}

func getEffectConfigurationFromUser(effectNumber int, scanner *bufio.Scanner) (int, int, error) {
	fmt.Println("Input effects configuration:")
	fmt.Printf("effect %d moves latch x \n", effectNumber)
	scanner.Scan()
	latchMoved := scanner.Text()
	latchNumber, err := strconv.Atoi(latchMoved)
	if err != nil {
		return 0, 0, fmt.Errorf("Latch moved must be a number, %w", err)
	}

	fmt.Printf(" in the x direction (1 means same, -1 means opposite\n")
	scanner.Scan()
	directionText := scanner.Text()
	direction, err := strconv.Atoi(directionText)
	if err != nil {
		return 0, 0, fmt.Errorf("direction must be a number, %w", err)
	}

	return latchNumber, direction, nil
}

func assignStartingPositionsToLatches(lock *Lock, scanner *bufio.Scanner) error {
	fmt.Println("Input latch starting positions")
	for i := range lock.Latches {
		fmt.Printf("Latch %d starting position: ", i)
		scanner.Scan()
		effectsText := scanner.Text()
		pos, err := strconv.Atoi(effectsText)
		if err != nil {
			return fmt.Errorf("Latch starting postion must be a number, %w", err)
		}
		lock.Latches[i].Position = pos
	}
	return nil
}
