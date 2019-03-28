package model

type Person struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Slackid    string `json:"slackid"`
	Experience string `json:"experience"`
	Interest   string `json:"interest"`
	Worktime   int32  `json:"worktime"`
	Like       int32  `json:"like"`
}
