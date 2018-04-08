package model

// AuthUser represents the credential for Authentication
type AuthUser struct {
	Username string
	Password string
}

type BBUsers struct {
	Size          int             `json:"size"`
	Limit         int             `json:"limit"`
	IsLastPage    bool            `json:"isLastPage"`
	Values        []BBUserDetails `json:"values"`
	Start         int             `json:"start"`
	NextPageStart int             `json:"nextPageStart"`
}

type BBUserDetails struct {
	Name                        string `json:"name"`
	EmailAddress                string `json:"emailAddress"`
	ID                          int    `json:"id"`
	DisplayName                 string `json:"displayName"`
	Active                      bool   `json:"active"`
	Slug                        string `json:"slug"`
	Type                        string `json:"type"`
	DirectoryName               string `json:"directoryName"`
	Deletable                   bool   `json:"deletable"`
	LastAuthenticationTimestamp int64  `json:"lastAuthenticationTimestamp,omitempty"`
	MutableDetails              bool   `json:"mutableDetails"`
	MutableGroups               bool   `json:"mutableGroups"`
}
