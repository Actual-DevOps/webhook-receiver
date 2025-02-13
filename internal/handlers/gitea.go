package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/elbars/webhook_receiver/internal/config"
	"github.com/elbars/webhook_receiver/internal/models"
)

func HandleGiteaWebhook(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	var payload models.WebhookPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	slog.Info("Received webhook from repository: " + payload.Repository.Name)
	slog.Info("Received webhook from repository FullName: " + payload.Repository.FullName)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received"))

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

func sendRequest(url string, user string, pass string) error {

	req, err := http.NewRequest("GET", url, nil)
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
