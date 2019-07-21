package data

import "time"

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func UserByEmail() (user User) {
	return
}

func (user *User) CreateSession() (session Session, err error) {
	return
}
