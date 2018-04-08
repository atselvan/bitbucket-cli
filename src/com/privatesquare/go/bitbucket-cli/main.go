package main

import (
	"log"
	"flag"
	m "com/abnamro/solo/bitbucket-cli/model"
	b "com/abnamro/solo/bitbucket-cli/backend"
	"fmt"
)

func main() {

	getInactiveUsers := flag.Bool("getInactiveUsers", false, "Get a list of users who have not logged in since 3 months")
	getProjectsList := flag.Bool("getProjectsList", false, "Get a list of projects available in bitbucket")
	projectExists := flag.Bool("projectExists", false, "Check if a project exists in bitbucket. Required Parameters: projectKey")
	username := flag.String("username", "", "Username for authentication")
	password := flag.String("password", "", 	"Password for authentication")
	bitbucketURL := flag.String("bitbucketURL", "", "Bitbucket server URL")
	projectKey := flag.String("projectKey", "", "The Key of a bitbucket project")
	verbose := flag.Bool("verbose", false, "Set this flag for Debug logs")
	flag.Parse()
	if *username == "" || *password == "" {
		log.Fatal("username and password are required parameters")
	} else if *bitbucketURL == "" {
		log.Fatal("bitbucketURL is a required parameter")
	}
	user := m.AuthUser{Username:*username, Password:*password}

	if *getInactiveUsers == true {
		b.GetInactiveUsers(*bitbucketURL, user, *verbose)
	}else if *getProjectsList == true {
		projectList := b.GetProjectList(*bitbucketURL, user, *verbose)
		fmt.Println(projectList)
	}else if *projectExists == true {
		b.ProjectExists(*bitbucketURL, *projectKey, user, *verbose)
	}else {
		log.Fatal("Please select a valid action flag. Check bitbucket-cli --help for more detials")
	}
}