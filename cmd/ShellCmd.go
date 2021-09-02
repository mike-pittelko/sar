package cmd

import (
	"strings"

	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "This is a shell version of this tool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		runShell()
	},
}
var lsCmd = &cobra.Command{
	Use:   "shell",
	Short: "This is a shell version of this tool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
        runLs()
	},
}

func init() {
	RootCmd.AddCommand(shellCmd)
    shellCmd.AddCommand(lsCmd)
}

func pcFromCommands(parent readline.PrefixCompleterInterface, c *cobra.Command) {
	pc := readline.PcItem(c.Use)
	parent.SetChildren(append(parent.GetChildren(), pc))
	for _, child := range c.Commands() {
		pcFromCommands(pc, child)
	}
}

func runLs() {
    for _, child := range RootCmd.Commands() {
            fmt.print(child)
        }   
}
func runShell() {
	completer := readline.NewPrefixCompleter()
	for _, child := range RootCmd.Commands() {
		pcFromCommands(completer, child)
	}

	shell, err := readline.NewEx(&readline.Config{
		Prompt:       "> ",
		AutoComplete: completer,
		EOFPrompt:    "exit",
	})
	if err != nil {
		panic(err)
	}
	defer shell.Close()

shell_loop:
	for {
		l, err := shell.Readline()
		if err != nil {
			break shell_loop
		}
		cmd, flags, err := RootCmd.Find(strings.Fields(l))
		if err != nil {
			shell.Terminal.Write([]byte(err.Error()))
		}
		cmd.ParseFlags(flags)
		cmd.Run(cmd, flags)
	}

}
