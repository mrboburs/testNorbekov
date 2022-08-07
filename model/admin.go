package model

type Admin struct {
	Id       int    `json:"-"`
	UserName string `json:"user_name" default:"justice"`
	Passord  string `json:"password" default:"12345"`
}
