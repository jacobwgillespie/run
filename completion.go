package main

import (
	"fmt"
	"os"
)

func completion(shell string) error {
	switch shell {
	case "bash":
		return runCmd.GenBashCompletion(os.Stdout)
	case "zsh":
		return runCmd.GenZshCompletion(os.Stdout)
	case "fish":
		return runCmd.GenFishCompletion(os.Stdout, true)
	case "powershell":
		return runCmd.GenPowerShellCompletionWithDesc(os.Stdout)
	case "unknown":
		return runCmd.Help()
	default:
		return fmt.Errorf("unknown shell: %s", shell)
	}
}
