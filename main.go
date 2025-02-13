package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/elbars/webhook_receiver/internal/config"
	"github.com/elbars/webhook_receiver/internal/handlers"
)

func main() {

	cfg, err := config.LoadConfig("config.yaml")

	if err != nil {
		slog.Error("Failed to load config: " + err.Error())
	}
	serverPort := fmt.Sprintf(":%s", cfg.ServerPort)

	handlerWithConfig := func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleGiteaWebhook(w, r, cfg)
	}

	http.HandleFunc("/webhook/gitea", handlerWithConfig)

	slog.Info("Webhook server is running on " + serverPort)
	if err := http.ListenAndServe(serverPort, nil); err != nil {
		slog.Error("Failed to start server: " + err.Error())
	}
}
