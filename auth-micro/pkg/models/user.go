package models

type User struct {
	Id       int    `json:"uid,omitempty" gorm:"primary_key"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
