package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Actual-DevOps/webhook-receiver/internal/config"
	"github.com/Actual-DevOps/webhook-receiver/internal/logging"
	"github.com/Actual-DevOps/webhook-receiver/internal/models"
)

func HandleGiteaWebhook(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.WebhookPayload

		logging.LogRequest(r, "gitea")

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			slog.Error(err.Error())
			return
		}

		slog.Info("Payload data",
			slog.String("repository_name", payload.Repository.Name),
			slog.String("repository_FullName", payload.Repository.FullName),
			slog.String("repository_Owner", payload.Repository.Owner.Login),
		)

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Webhook received"))

		if err != nil {
			slog.Error(err.Error())
		}

		for _, webhook := range cfg.Jenkins.AllowedWebhooks {
			if webhook.RepoName == payload.Repository.FullName {
				for _, job := range webhook.RunJobs {
					slog.Info("Allowed repository: " + payload.Repository.FullName)

					var buildQuery string
					if job.ParameterizedJob {
						buildQuery = "buildWithParameters"
					} else {
						buildQuery = "build"
					}

					url := fmt.Sprintf("%s/%s/%s?token=%s", cfg.Jenkins.URL, job.JobPath, buildQuery, cfg.Jenkins.Token)

					err := sendRequest(url, cfg.Jenkins.User, cfg.Jenkins.Pass)
					if err != nil {
						slog.Error(err.Error())
					}
				}
			}
		}
	}
}

func sendRequest(url string, user string, pass string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	basecAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pass)))
	req.Header.Add("Authorization", basecAuth)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	slog.Info("Response Status: " + resp.Status)

	return nil
}
