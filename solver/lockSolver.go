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

}
