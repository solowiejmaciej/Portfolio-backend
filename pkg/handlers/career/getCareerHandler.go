package career

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"portfolio-api/models/responses"
)

func GetCareer(w http.ResponseWriter, request *http.Request) {
	var response = []responses.CareerResponse{
		{
			Company:   "Software Studio",
			JobTitle:  "Software Developer",
			StartDate: "2024-02-01",
			Current:   true,
			Technologies: []string{
				"Docker",
				"Vue.js",
				"ASP.NET Core",
				"Entity Framework",
				"SQL Server",
				"Git",
			},
		},
		{
			Company:   "eFitness",
			JobTitle:  "IT Consultant",
			StartDate: "2021-11-01",
			EndDate:   "2024-01-31",
			Current:   false,
			Technologies: []string{
				"SQL Server",
				"Azure",
				"Azure DevOps",
				"Python",
				"IT Consulting",
				"Project Management",
			},
		},
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Errorf("Error while serializing to json: %v", err)
	}
}
