package main

import (
	"fmt"
	"os"
)

func main() {

	if err := runMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

var scriptName string
var scriptArgs []string

func runMain() error {
	displayHelp := false

Loop:
	for idx, arg := range os.Args {
		if idx == 0 {
			continue
		}

		switch {
		case arg == "__complete" || arg == "__completeNoDesc":
			break Loop

		case arg == "--help" || arg == "-h":
			displayHelp = true
			continue

		case arg == "--completion":
			shell := "unknown"
			if len(os.Args) > idx+1 {
				shell = os.Args[idx+1]
			}
			return completion(shell)

		case arg[0] != '-':
			scriptName = arg
			scriptArgs = os.Args[idx+1:]
			os.Args = os.Args[:idx+1]
			break Loop

		default:
			return fmt.Errorf("unknown flag: %s", arg)
		}
	}

	if displayHelp || len(os.Args) == 1 {
		return runCmd.Help()
	}

	return runCmd.Execute()
}
