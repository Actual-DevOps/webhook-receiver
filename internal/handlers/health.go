package handlers

import (
	"log/slog"
	"net/http"
	"github.com/Actual-DevOps/webhook-receiver/internal/logging"
)

func HandleHealthWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logging.LogRequest(r, "health")

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))

		if err != nil {
			slog.Error(err.Error())
		}
	}
}
