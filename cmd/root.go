package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// configFile configuration file
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "dotfiles-cli",
		Short: "Basic CLI tool for dotfiles administration",
		Long: `This application is a tool to generate, configure
and restore dotfiles for a unix system.`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dotfiles.yaml)")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// initConfig before command exec, configure with file
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".dotfiles" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".dotfiles")
	}

	viper.AutomaticEnv()

	var (
		err      = viper.ReadInConfig()
		notFound = &viper.ConfigFileNotFoundError{}
	)

	if err != nil {
		switch {
		case errors.As(err, notFound):
			// The config file is optional, we shouldn't exit when the config is not found
			log.Printf("using defualt config file '%s'\n", cfgFile)
		default:
			cobra.CheckErr(err)
		}
	}
}
