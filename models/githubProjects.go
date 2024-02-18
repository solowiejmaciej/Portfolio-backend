package models

type GithubProjects struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	FullName    string   `json:"full_name"`
	Private     bool     `json:"private"`
	HtmlUrl     string   `json:"html_url"`
	Description string   `json:"description"`
	Fork        bool     `json:"fork"`
	Url         string   `json:"url"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Topics      []string `json:"topics"`
	Visibility  string   `json:"visibility"`
	Homepage    string   `json:"homepage"`
}
