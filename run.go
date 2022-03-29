package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var green = lipgloss.NewStyle().Foreground(lipgloss.Color("28"))
var gray = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("8"))
var red = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))

var runCmd = &cobra.Command{
	Use:                "run [script]",
	Short:              "run runs scripts in package.json",
	SilenceUsage:       true,
	Args:               cobra.MinimumNArgs(1),
	DisableFlagParsing: true,

	Long: `To load completions:

Bash:

  $ source <(run --completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ run --completion bash > /etc/bash_completion.d/run
  # macOS:
  $ run --completion bash > /usr/local/etc/bash_completion.d/run

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ run --completion zsh > "${fpath[1]}/_run"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ run --completion fish | source

  # To load completions for each session, execute once:
  $ run --completion fish > ~/.config/fish/completions/run.fish

PowerShell:

  PS> run --completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> run --completion powershell > run.ps1
  # and source this file from your PowerShell profile.
`,

	Run: func(cmd *cobra.Command, args []string) {
		err := runScript(scriptName, scriptArgs)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "%s\n", red.Render(err.Error()))
		}
	},

	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return listScripts(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	runCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	runCmd.InitDefaultHelpFlag()
	_ = runCmd.Flags().MarkHidden("help")
}
