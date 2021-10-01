package model

type Content struct {
	Username  string `json:"username"`
	Followers int    `json:"followers"`
}

type ResFoll struct {
	Followers int `json:"followers"`
}
