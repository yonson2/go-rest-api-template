package models

type User struct {
	Base
	AvatarUrl string
	Email     string
	Username  string
}
