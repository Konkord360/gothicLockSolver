package main

import (
	"fmt"
	"gothicModSover/simulator"
	"gothicModSover/solver"
	"os"
)

func main() {
	switch parseInputArgs() {
	case "simulation":
		simulator.RunSimulator()
	case "":
		error := solver.RunSolver()
		if error != nil {
			fmt.Printf("Received error from the solver: %s", error)
		}
	}
}

func parseInputArgs() string {
	cmdArgs := os.Args[:]

	if len(cmdArgs) == 1 {
		fmt.Println("No cmd args provided")
	} else {
		fmt.Printf("Provided argument: %s\n", cmdArgs[1])
	}

	if len(cmdArgs) > 1 && cmdArgs[1] == "simulation" {
		fmt.Println("Running lock pick simulation")
		return cmdArgs[1]
	}

	return ""
}
