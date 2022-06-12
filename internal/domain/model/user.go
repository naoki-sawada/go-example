package model

import (
	"fmt"
	"time"
)

type User struct {
	ID        string    `db:"id" json:"id" validate:"required,uuid4"`
	FirstName string    `db:"first_name" json:"firstName" validate:"required,min=0,max=30"`
	LastName  string    `db:"last_name" json:"lastName" validate:"required,min=0,max=30"`
	Email     string    `db:"email" json:"email" validate:"required,email"`
	Birthdate time.Time `db:"birthdate" json:"birthdate" validate:"required"`
}

func (u User) Age(t time.Time) uint {
	return Age(u.Birthdate, t)
}

func (u User) IsSenior(t time.Time) bool {
	age := u.Age(t)
	return IsSenior(age)
}

func (u User) IsAdult(t time.Time) bool {
	age := u.Age(t)
	return IsAdult(age)
}

func (u User) IsChild(t time.Time) bool {
	age := u.Age(t)
	return IsChild(age)
}

type UserList []User

func (ul UserList) HasChild(t time.Time) bool {
	for _, u := range ul {
		if u.IsChild(t) {
			return true
		}
	}
	return false
}

func Age(birthdate, today time.Time) uint {
	ty, tm, td := today.Date()
	t := time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	b := time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if t.Before(b) {
		return 0
	}
	age := ty - by
	anniversary := b.AddDate(age, 0, 0)
	if anniversary.After(t) {
		age--
	}
	return uint(age)
}

func IsSenior(age uint) bool {
	return 65 <= age
}

func IsAdult(age uint) bool {
	return 12 <= age && age <= 64
}

func IsChild(age uint) bool {
	return 3 <= age && age <= 11
}

func (u User) FullName() string {
	return fmt.Sprintf("%s %s", u.LastName, u.FirstName)
}
