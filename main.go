package main

import (
	"fmt"
	"os"

	"github.com/blazed/shorten/cmd"
	"github.com/spf13/cobra"
)

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "shorten",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	rootCmd.AddCommand(cmd.Serve())
	return rootCmd
}

func main() {
	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
