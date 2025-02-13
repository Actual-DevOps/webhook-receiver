package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerPort string `yaml:"server_port"`
	Jenkins    struct {
		URL             string `yaml:"url"`
		User            string `yaml:"user"`
		Pass            string `yaml:"pass"`
		Token           string `yaml:"token"`
		AllowedWebhooks []struct {
			RepoName string `yaml:"repo_name"`
			RunJobs  []struct {
				JobPath          string `yaml:"job_path"`
				ParameterizedJob bool   `yaml:"parameterized_job"`
			} `yaml:"run_jobs"`
		} `yaml:"allowed_webhooks"`
	} `yaml:"jenkins"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
