package model

type Person struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Slackid  string `json:"slackid"`
	Name     string `json:"name"`
	Interest string `json:"interest"`
	Image    string `json:"image"`
	Worktime int32  `json:"worktime"`
	Like     int32  `json:"like"`
}
