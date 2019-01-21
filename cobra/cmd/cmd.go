package cmd

import (
	"flag"
	"github.com/spf13/cobra"
)

type CommandGroup struct {
	Message string
	Commands []*cobra.Command
}

func (cg CommandGroup) Add(c *cobra.Command) {
	c.AddCommand(cg.Commands...)
}

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "cobra-lab",
		Short: "cli lab with cobra",
		Run: runHelp,
	}

	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	group := CommandGroup{
		Message: "Basic Commands",
		Commands: []*cobra.Command{
			NewCmdHello(),
		},
	}
	group.Add(cmd)
	return cmd
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
