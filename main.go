package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/elbars/webhook_receiver/internal/config"
	"github.com/elbars/webhook_receiver/internal/handlers"
	"github.com/spf13/cobra"
)

func wrapHelpFunctionWithExit(cmd *cobra.Command) {
	helpFunc := cmd.HelpFunc()
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		helpFunc(cmd, args)
		os.Exit(0)
	})
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "go run main.go",
		Short: "Path to config",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	wrapHelpFunctionWithExit(rootCmd)

	defaultConfigPath := "config/config.yaml"

	var configPath string

	rootCmd.Flags().StringVarP(&configPath, "config", "c", defaultConfigPath, "Сonfig path")

	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
	}

	cfg, err := config.LoadConfig(configPath)

	if err != nil {
		slog.Error("Failed to load config: " + err.Error())
	}

	serverPort := fmt.Sprintf(":%s", cfg.ServerPort)

	http.HandleFunc("/webhook/gitea", handlers.HandleGiteaWebhook(cfg))

	slog.Info("Webhook server is running on " + serverPort)

	if err := http.ListenAndServe(serverPort, nil); err != nil {
		slog.Error("Failed to start server: " + err.Error())
	}
}
