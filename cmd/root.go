package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dotfiles-cli",
		Short: "Basic CLI tool for dotfiles administration",
		Long: `This application is a tool to generate, configure
and restore dotfiles for a unix system.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
