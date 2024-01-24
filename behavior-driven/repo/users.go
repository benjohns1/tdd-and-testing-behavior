package repo

import (
	app "github.com/benjohns1/tdd-and-testing-behavior/behavior-driven"
)

type Users struct {
	user app.User
}

func (u *Users) Save(user app.User) error {
	u.user = user
	return nil
}

func (u *Users) GetAll() ([]app.User, error) {
	return []app.User{u.user}, nil
}
