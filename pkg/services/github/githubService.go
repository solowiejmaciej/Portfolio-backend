package githubService

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"portfolio-api/models"
)

func GetGithubProjects() []models.GithubProjects {
	log.Println("Fetching projects from Github...")
	var client = http.Client{}
	request, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	if err != nil {
		log.Errorf("Error creating new request: %v", err)
		return nil
	}
	request.Header.Set("Authorization", "Bearer "+TOKEN)

	response, err := client.Do(request)
	if err != nil {
		log.Errorf("Error while fetching projects from Github: %v", err)
		return nil
	}

	if response.StatusCode != 200 {
		log.Errorf("Error while fetching projects from Github: %v", response.Status)
		return nil
	}

	var projects []models.GithubProjects

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
		return nil
	}

	err = json.Unmarshal(body, &projects)
	if err != nil {
		log.Errorf("Error while unmarshalling response body: %v", err)
		return nil
	}

	log.Infof("Fetched %d projects from Github", len(projects))

	return projects
}
