package projects

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"portfolio-api/models"
	"portfolio-api/models/responses"
	githubService "portfolio-api/pkg/services/github"
	"time"
)

func GetProjects(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var projects = githubService.GetGithubProjects()
	var filteredProjects = filterProjects(projects)
	var response = mapToResponse(filteredProjects)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Errorf("Error while serializing to json: %v", err)
	}
}

func filterProjects(projects []models.GithubProjects) []models.GithubProjects {
	var filteredProjects []models.GithubProjects

	for _, project := range projects {
		if !project.Private {
			portfolioFound := false
			for _, topic := range project.Topics {
				if topic == "portfolio" {
					portfolioFound = true
					break
				}
			}
			if portfolioFound {
				var newTopics []string
				for _, topic := range project.Topics {
					if topic != "portfolio" {
						newTopics = append(newTopics, topic)
					}
				}
				project.Topics = newTopics
				filteredProjects = append(filteredProjects, project)
			}
		}
	}

	return filteredProjects
}

func mapToResponse(projects []models.GithubProjects) []responses.ProjectsResponse {
	var response []responses.ProjectsResponse
	for _, project := range projects {
		var res responses.ProjectsResponse
		date, _ := time.Parse(time.RFC3339, project.CreatedAt)
		var formattedDate = date.Format("01/02/2006")
		switch project.Id {
		case 720435321:
			res = responses.ProjectsResponse{
				Id:           project.Id,
				Title:        project.Name,
				Description:  "BloodRush is a API for mobile application. Its all about notifying people about blood needs in their area. Purpose of this application is to help people in need of blood and share awareness about blood donation. If you register, you will be notified when there is a need for blood in your area within specified radius. After successful donation you will receive notification with information about your examination results. Resting period will be calculated based on your sex. While resting you will not receive notifications about blood needs. You will be notified when you can donate blood again.",
				GitHubUrl:    project.HtmlUrl,
				CreatedAt:    formattedDate,
				Icon:         "Blood",
				IsLive:       project.Homepage != "",
				LiveUrl:      project.Homepage,
				Technologies: project.Topics,
				Images:       []string{""},
				Tags:         []string{"#backend", "#api"},
			}
		case 706939953:
			res = responses.ProjectsResponse{
				Id:           project.Id,
				Title:        project.Name,
				Description:  "HomeManagementService is a API for home automation. It is a part of my home automation project. It is responsible for managing devices, it integrates with Philips Hue, and my own devices based on Esp32. It uses Azure CosmosDb to store automation rules, automation can be triggered by time, or by location. It runs on Docker, and is deployed on my own server.",
				GitHubUrl:    project.HtmlUrl,
				CreatedAt:    formattedDate,
				Icon:         "Home",
				IsLive:       project.Homepage != "",
				LiveUrl:      project.Homepage,
				Technologies: project.Topics,
				Images:       []string{""},
				Tags:         []string{"#backend", "#api"},
			}
		case 722597011:
			res = responses.ProjectsResponse{
				Id:           project.Id,
				Title:        project.Name,
				Description:  "This is a example system based on events and microservices. It is a simple system for managing users and their data. It is based on the CQRS pattern and the Mediator pattern, it uses RabbitMQ to communicate between services. It also uses Redis to cache data. Users can be created, updated, and deleted. It is a simple system that I created to learn more about microservices and event driven architecture.",
				GitHubUrl:    project.HtmlUrl,
				CreatedAt:    formattedDate,
				Icon:         "Laptop",
				IsLive:       project.Homepage != "",
				LiveUrl:      project.Homepage,
				Technologies: project.Topics,
				Images:       []string{""},
				Tags:         []string{"#backend", "#api", "#microservices"},
			}
		case 744449078:
			res = responses.ProjectsResponse{
				Id:           project.Id,
				Title:        project.Name,
				Description:  "This is my portfolio website. It is a simple website that I created to showcase my projects. It is based on Vue.js, it it uses Vue Router to navigate between pages. You are currently on this website.",
				GitHubUrl:    project.HtmlUrl,
				CreatedAt:    formattedDate,
				Icon:         "Suitcase",
				IsLive:       project.Homepage != "",
				LiveUrl:      project.Homepage,
				Technologies: project.Topics,
				Images:       []string{""},
				Tags:         []string{"#frontend"},
			}

		case 715790788:
			res = responses.ProjectsResponse{
				Id:           project.Id,
				Title:        project.Name,
				Description:  "SmartKnob is a project that I created to learn more about Esp32, and home automation. It is a simple device that can be used to control lights in your home. It is based on Esp32, and uses Mqtt to communicate with other services.",
				GitHubUrl:    project.HtmlUrl,
				CreatedAt:    formattedDate,
				Icon:         "Plug",
				IsLive:       project.Homepage != "",
				LiveUrl:      project.Homepage,
				Technologies: project.Topics,
				Images:       []string{"https://i.imgur.com/U5KOut6.jpg", "https://i.imgur.com/SMTS8Hz.png"},
				Tags:         []string{"#hardware", "#iot"},
			}
		case 722326104:
			res = responses.ProjectsResponse{
				Id:           project.Id,
				Title:        project.Name,
				Description:  "SmartWeatherStation is a project that I created to learn more about Esp32, and home automation. It is a simple device that can be used to measure temperature, humidity, and pressure. It is based on Esp32, and uses Mqtt to communicate with other services.",
				GitHubUrl:    project.HtmlUrl,
				CreatedAt:    formattedDate,
				Icon:         "Temperature",
				IsLive:       project.Homepage != "",
				LiveUrl:      project.Homepage,
				Technologies: project.Topics,
				Images:       []string{"https://i.imgur.com/9JWkqB7.jpeg", "https://i.imgur.com/0jDzREd.png"},
				Tags:         []string{"#hardware", "#iot"},
			}
		case 759441947:
			res = responses.ProjectsResponse{
				Id:           project.Id,
				Title:        project.Name,
				Description:  "This is a simple API that integrates with github, it is a part of my portfolio website. It is responsible for fetching my projects from github, and displaying them on my website. It is based on Go, and uses mux to handle requests.",
				GitHubUrl:    project.HtmlUrl,
				CreatedAt:    formattedDate,
				Icon:         "Default",
				IsLive:       project.Homepage != "",
				LiveUrl:      project.Homepage,
				Technologies: project.Topics,
				Images:       []string{""},
				Tags:         []string{"#backend", "#api"},
			}
		default:
			res = responses.ProjectsResponse{
				Id:           project.Id,
				Title:        project.Name,
				Description:  "Lorem",
				GitHubUrl:    project.HtmlUrl,
				CreatedAt:    formattedDate,
				Icon:         "Default",
				IsLive:       project.Homepage != "",
				LiveUrl:      project.Homepage,
				Technologies: project.Topics,
				Images:       []string{""},
				Tags:         []string{"#new"},
			}
		}

		response = append(response, res)

	}
	return response
}
