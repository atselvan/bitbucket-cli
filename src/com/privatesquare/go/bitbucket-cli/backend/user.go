package backend

import (
	m "com/privatesquare/go/bitbucket-cli/model"
	u "com/privatesquare/go/bitbucket-cli/utils"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func GetInactiveUsers(bitbucketURL string, user m.AuthUser, verbose bool) []string {
	start := 0
	limit := 100
	isLastPage := false
	currentTimestamp := time.Now().UnixNano() / 1000000
	var (
		bbUserDetails []m.BBUserDetails
		inactiveUsers []string
	)
	for isLastPage == false {
		var bbUsers m.BBUsers
		url := fmt.Sprintf("%s/rest/api/1.0/admin/users", bitbucketURL)
		req := u.CreateBaseRequest("GET", url, nil, user, verbose)
		query := req.URL.Query()
		query.Add("start", fmt.Sprintf("%d", start))
		query.Add("limit", fmt.Sprintf("%d", limit))
		req.URL.RawQuery = query.Encode()
		respBody, _ := u.HTTPRequest(req, verbose)
		json.Unmarshal(respBody, &bbUsers)
		for _, user := range bbUsers.Values {
			userDetails := user
			bbUserDetails = append(bbUserDetails, userDetails)
		}
		start = bbUsers.NextPageStart
		isLastPage = bbUsers.IsLastPage
	}
	for _, user := range bbUserDetails {
		lastLoginMilli := currentTimestamp - user.LastAuthenticationTimestamp
		lastLoginDays := lastLoginMilli / (1000 * 60 * 60 * 24)
		if lastLoginDays >= 90 && lastLoginDays < 17624 && user.Name != "atlbitbucket" && user.Name != "abnamro" {
			inactiveUsers = append(inactiveUsers, user.Name)
			fmt.Println(user.Name, lastLoginDays)
		}
	}
	fmt.Println("Total Number of Users : ", len(bbUserDetails))
	fmt.Println("Number of Inactive Users : ", len(inactiveUsers))

	return inactiveUsers
}

func DeactivateInactiveUsers(bitbucketURL string, user m.AuthUser, verbose bool) {
	inactiveUsers := GetInactiveUsers(bitbucketURL, user, verbose)
	var counter int
	for _, userId := range inactiveUsers {
		url := fmt.Sprintf("%s/rest/api/1.0/admin/groups/remove-user", bitbucketURL)
		groupUser := m.GroupUser{Context: "BITBUCKET_USERS", ItemName: userId}
		payload, err := json.Marshal(groupUser)
		u.Error(err, "Error creating the request body")
		req := u.CreateBaseRequest("POST", url, payload, user, verbose)
		_, status := u.HTTPRequest(req, verbose)
		if status == "200 OK" {
			log.Printf("User '%s' has been deactivated\n", userId)
			counter ++
		} else if status == "404 Not Found"{
			log.Printf("The user '%s' is already deactivated", userId)
		} else {
			log.Printf("There was a error while deactivating the user '%s'. Activate verbose logs for more details\n", userId)
		}
	}
	log.Printf("%v users were deactivated\n", counter)
}
