package backend

import (
	m "com/privatesquare/go/bitbucket-cli/model"
	u "com/privatesquare/go/bitbucket-cli/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetProjectList(bitbucketURL string, user m.AuthUser, verbose bool) []string {
	start := 0
	limit := 100
	isLastPage := false
	var (
		projectInfoList []m.ProjectsInfo
		projectList     []string
	)
	for isLastPage == false {
		var projects m.Projects
		url := fmt.Sprintf("%s/rest/api/1.0/projects", bitbucketURL)
		req := u.CreateBaseRequest("GET", url, nil, user, verbose)
		query := req.URL.Query()
		query.Add("start", fmt.Sprintf("%d", start))
		query.Add("limit", fmt.Sprintf("%d", limit))
		req.URL.RawQuery = query.Encode()
		respBody, _ := u.HTTPRequest(req, verbose)
		json.Unmarshal(respBody, &projects)
		for _, project := range projects.Values {
			projectInfo := project
			projectInfoList = append(projectInfoList, projectInfo)
		}
		start = projects.NextPageStart
		isLastPage = projects.IsLastPage
	}
	for _, projectInfo := range projectInfoList {
		projectList = append(projectList, projectInfo.Key)
	}
	return projectList
}

func ProjectExists(bitbucketURL, projectKey string, user m.AuthUser, verbose bool) bool {
	if projectKey == "" {
		log.Fatal("projectKey is a required parameters for checking of a project exists")
	}
	var isExists bool
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s", bitbucketURL, projectKey)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	_, status := u.HTTPRequest(req, verbose)
	switch status {
	case "200 OK":
		log.Printf("Project with key '%s' exists", projectKey)
		isExists = true
	case "404 Not Found":
		log.Printf("Project with key '%s' does not exist", projectKey)
		isExists = false
	case "401 Unauthorized":
		log.Printf("User '%s' is unauthorized or the password is incorrect", user.Username)
		os.Exit(1)
	default:
		log.Fatal("There was a error while checking if the project exists. Status : ", status)
	}
	return isExists
}
