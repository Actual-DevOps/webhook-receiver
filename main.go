package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/elbars/webhook_receiver/internal/config"
	"github.com/elbars/webhook_receiver/internal/handlers"
	"github.com/elbars/webhook_receiver/cmd"
)

func main() {
	rootCmd := cmd.NewRootCommand()

	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	cfg, err := config.LoadConfig(cmd.GetConfigPath())
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
}
