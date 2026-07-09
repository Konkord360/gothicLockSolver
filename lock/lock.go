package lock

import( "fmt"
	"bufio"
	"os"
	"strconv" )

const (
	opposite int = -1
	same     int = 1

	Left     int = -1
	Right    int = 1
)

type Lock struct {
	Latches []*Latch
}

type Latch struct {
	Position int      // latch is in a position 1 -7 with 4 being the middle
	Effects  []Effect // latch hes effects - can move other latches
}

type Effect struct {
	Targets []*Latch // effect affects latches
	Moves   []int    // moves to execute on latches - 1 to 1 mapping moves to targets
}

func (l *Latch) Move(m int) {
	l.Position += m

	for _, e := range l.Effects {
		e.apply(m)
	}
}

func (e Effect) apply(originalMove int) {
	moves := e.Moves
	targets := e.Targets

	for i, target := range targets {
		target.moveWithoutEffects(moves[i] * originalMove)
	}
}

func (l *Latch) moveWithoutEffects(m int) {
	l.Position += m
}

func (l Lock) PrintLock() {
	fmt.Println("***********")
	for _, e := range l.Latches {
		fmt.Println(e.Position)

	}
	fmt.Println("***********")
}

//make it any better plz mate - i cant read it 
func CreateLockFromUserInput() (*Lock, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input lock configuration:")
	fmt.Print("Number of latches: ")
	scanner.Scan()
	latchesText := scanner.Text()
	latches, err := strconv.Atoi(latchesText)
	if err != nil {
		return nil, fmt.Errorf("Latches must be a number, %w", err)
	}

	if latches <= 2 {
		return nil, fmt.Errorf("There must be more then two latches")
	}

	latchesArray := []*Latch{}
	for range latches {
		latchesArray = append(latchesArray, &Latch{})
	}

	lockVar := Lock{Latches: latchesArray}

	fmt.Println("Assing an effect to each of the latches: ")

	for i := range latches {
		fmt.Printf("Latch %d - number of effects", i)	
		scanner.Scan()
		effectsText := scanner.Text()
		effects, err := strconv.Atoi(effectsText)

		if err != nil {
			return nil, fmt.Errorf("effects must be a number, %w", err)
		}
		if effects < 0 {
			return nil, fmt.Errorf("effects must be non negative number")
		}

		effectArray := []Effect{}
		latch := Latch{Position: i, Effects: effectArray}
		fmt.Println(latch)
		movesArray := []int{}
		targetArray := []*Latch{}
		for j := range effects {
			fmt.Println("Input effects configuration:")

			fmt.Printf("effect %d moves latch x \n", j)

			scanner.Scan()
			latchMoved := scanner.Text()
			latchNumber, err := strconv.Atoi(latchMoved)
			targetArray = append(targetArray, lockVar.Latches[latchNumber])

			if err != nil {
				return nil, fmt.Errorf("Latch moved must be a number, %w", err)
			}


			fmt.Printf(" in the x direction (1 means same, -1 means opposite\n")
			scanner.Scan()
			directionText := scanner.Text()
			direction, err := strconv.Atoi(directionText)
			movesArray = append(movesArray, direction)

			if err != nil {
				return nil, fmt.Errorf("Latch moved must be a number, %w", err)
			}
			latch.Effects = append(latch.Effects, Effect{Moves: movesArray, Targets: targetArray})
		}
		lockVar.Latches[i] = &latch
	}

	fmt.Println("Input latch starting positions")

	for i := range latchesArray {
		fmt.Printf("Latch %d starting position: ", i)
		scanner.Scan()
		effectsText := scanner.Text()
		pos, err := strconv.Atoi(effectsText)
		if err != nil {
			return nil, fmt.Errorf("Latch starting postion must be a number, %w", err)
		}
		latchesArray[i].Position = pos
	}

		
	fmt.Println("Created lock:")
	lockVar.PrintLock()

	return &lockVar, nil
}
