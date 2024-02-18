package responses

type ProjectsResponse struct {
	Id           int
	Title        string
	Description  string
	GitHubUrl    string
	CreatedAt    string
	Icon         string
	IsLive       bool
	LiveUrl      string
	Technologies []string
	Images       []string
	Tags         []string
}
