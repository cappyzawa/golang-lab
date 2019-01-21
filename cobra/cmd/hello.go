package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type options struct {
	name string
}

func NewOptions() *options {
	return &options{}
}

func NewCmdHello() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use: "hello",
		Short: "hello command",
		Run: func(cmd *cobra.Command, args []string) {
			o.Say(cmd, args)
		},
	}
	cmd.PersistentFlags()
	cmd.Flags().StringVarP(&o.name, "name", "n", "", "set a name")
	return cmd
}

func (o *options) Say(cmd *cobra.Command, args []string) {
	if o.name != "" {
		fmt.Fprintf(os.Stdout, "hello %s\n", o.name)
		return
	}
	if len(os.Args) != 1 {
		fmt.Fprintf(os.Stderr, "error: one arg is required\n")
		return
	}
	fmt.Fprintf(os.Stdout, "hello %s\n", args[0])
	return
}
