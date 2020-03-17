package models

// Auth -> return auth user and access token
// after a successful login
type Auth struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
