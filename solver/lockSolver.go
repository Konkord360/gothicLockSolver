package solver

import (
	"fmt"
	"gothicModSover/lock"
)

func RunSolver() error{
	lockCraeted, err := lock.CreateLockFromUserInput()
	if (err != nil) {
		return fmt.Errorf("Error occured during lock creation form user input: %w", err)
	}
	fmt.Println(lockCraeted)
	return nil
}

