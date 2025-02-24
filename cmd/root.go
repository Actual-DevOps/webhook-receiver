package cmd

import (
	"fmt"
	"github.com/Actual-DevOps/webhook-receiver/internal/config"
	"github.com/Actual-DevOps/webhook-receiver/internal/handlers"
	"github.com/spf13/cobra"
	"log/slog"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "webhook-receiver",
	Short: "Start webhook receiver server",
	Run: func(cmd *cobra.Command, _ []string) {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			slog.Error("Failed to parse flag: " + err.Error())
			os.Exit(1)
		}

		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			slog.Error("Failed to load config: " + err.Error())
			os.Exit(1)
		}

		serverPort := fmt.Sprintf(":%s", cfg.ServerPort)

		http.HandleFunc("/webhook/gitea", handlers.HandleGiteaWebhook(cfg))

		slog.Info("Webhook server is running on " + serverPort)

		if err := http.ListenAndServe(serverPort, nil); err != nil {
			slog.Error("Failed to start server: " + err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
	}
}
