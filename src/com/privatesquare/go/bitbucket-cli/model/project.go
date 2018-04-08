package model

type Projects struct {
	Size          int            `json:"size"`
	Limit         int            `json:"limit"`
	IsLastPage    bool           `json:"isLastPage"`
	Values        []ProjectsInfo `json:"values"`
	Start         int            `json:"start"`
	NextPageStart int            `json:"nextPageStart"`
}

type ProjectsInfo struct {
	Key         string `json:"key"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Type        string `json:"type"`
}
