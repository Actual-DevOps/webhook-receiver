package logging

import (
	"log/slog"
	"net/http"
)

func LogRequest(r *http.Request, eventType string) {
	slog.Info("Webhook received",
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("remote_addr", r.RemoteAddr),
		slog.String("user_agent", r.UserAgent()),
		slog.String("event_type", eventType),
	)
}
