package repo

import app "github.com/benjohns1/tdd-and-testing-behavior/implementation-driven"

type Users struct {
	users map[string]app.User
}

func (u *Users) Save(user app.User) error {
	u.users = map[string]app.User{
		user.Name: user,
	}
	return nil
}
