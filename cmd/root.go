package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configPath string

func wrapHelpFunctionWithExit(cmd *cobra.Command) {
	helpFunc := cmd.HelpFunc()
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		helpFunc(cmd, args)
		os.Exit(0)
	})
}

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go run main.go",
		Short: "Path to config",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	wrapHelpFunctionWithExit(rootCmd)

	defaultConfigPath := "config/config.yaml"
	rootCmd.Flags().StringVarP(&configPath, "config", "c", defaultConfigPath, "Config path")

	return rootCmd
}

func GetConfigPath() string {
	return configPath
}
