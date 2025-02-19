package cmd

import (
	"fmt"
	"github.com/elbars/webhook_receiver/internal/config"
	"github.com/elbars/webhook_receiver/internal/handlers"
	"github.com/spf13/cobra"
	"log/slog"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "webhook_receiver",
	Short: "Start webhook receiver server",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")

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
		fmt.Println(err)
	}
}
