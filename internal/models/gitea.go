package models

type WebhookPayload struct {
	Ref        string     `json:"ref"`
	Repository Repository `json:"repository"`
	Action     string     `json:"action"`
	Commits    []Commit   `json:"commits"`
}

type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    Owner  `json:"owner"`
	URL      string `json:"html_url"`
}

type Owner struct {
	Login string `json:"login"`
}

type Commit struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Author  Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}