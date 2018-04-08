package backend

import (
	"fmt"
	"encoding/json"
	"time"
	u "com/abnamro/solo/bitbucket-cli/utils"
	m "com/abnamro/solo/bitbucket-cli/model"
)

func GetInactiveUsers(bitbucketURL string, user m.AuthUser, verbose bool){
	start := 0
	limit := 100
	isLastPage := false
	currentTimestamp := time.Now().UnixNano()/ 1000000
	var (
		bbUserDetails []m.BBUserDetails
		inactiveUsers []string
	)
	for isLastPage == false{
		var bbUsers m.BBUsers
		url := fmt.Sprintf("%s/rest/api/1.0/admin/users", bitbucketURL)
		req := u.CreateBaseRequest("GET", url, nil, user, verbose)
		query := req.URL.Query()
		query.Add("start", fmt.Sprintf("%d", start))
		query.Add("limit", fmt.Sprintf("%d", limit))
		req.URL.RawQuery = query.Encode()
		respBody, _ := u.HTTPRequest(req, verbose)
		json.Unmarshal(respBody, &bbUsers)
		for _,user := range bbUsers.Values {
			userDetails := user
			bbUserDetails = append(bbUserDetails, userDetails)
		}
		start = bbUsers.NextPageStart
		isLastPage = bbUsers.IsLastPage
	}
	for _, user := range bbUserDetails{
		lastLoginMilli := currentTimestamp - user.LastAuthenticationTimestamp
		lastLoginDays := lastLoginMilli / (1000*60*60*24)
		if lastLoginDays >= 90 && lastLoginDays < 17624  && user.Name != "atlbitbucket" {
			inactiveUsers = append(inactiveUsers, user.Name)
			fmt.Println(user.Name, lastLoginDays)
		}
	}
	fmt.Println("Total Number of Users : ", len(bbUserDetails))
	fmt.Println("Number of Inactive Users : ", len(inactiveUsers))
}