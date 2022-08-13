package models

type Access struct {
	AccessToken string `json:"access_token"`
	Scope       string
}

type User struct {
	AccessToken string
	AvatarUrl   string
	ProfileUrl  string
	FullName    string
	Company     string
	Location    string
}
