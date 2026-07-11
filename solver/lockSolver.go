package solver

import (
	"fmt"
	"gothicModSover/lock"
)

func RunSolver() error {
	lockCreated, err := lock.CreateLockFromUserInput()
	if err != nil {
		return fmt.Errorf("Error occured during lock creation form user input: %w", err)
	}
	lockCreated.PrintLock()
	lockCreated.PrintEffects()
	return nil
}
