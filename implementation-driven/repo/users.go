package repo

import (
	"fmt"
	app "github.com/benjohns1/tdd-and-testing-behavior/implementation-driven"
)

type Users struct {
	users map[string]app.User
}

func (u *Users) Save(user app.User) error {
	if _, exists := u.users[user.Name]; exists {
		return fmt.Errorf("user name %q already exists", user.Name)
	}
	if u.users == nil {
		u.users = make(map[string]app.User, 1)
	}
	u.users[user.Name] = user
	return nil
}
