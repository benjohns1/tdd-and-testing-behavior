package repo

import (
	"fmt"
	app "github.com/benjohns1/tdd-and-testing-behavior/behavior-driven"
)

type Users struct {
	users []app.User
}

func (u *Users) Save(user app.User) error {
	for _, current := range u.users {
		if current.Name == user.Name {
			return fmt.Errorf("user name %q already exists", user.Name)
		}
	}
	u.users = append(u.users, user)
	return nil
}

func (u *Users) GetAll() ([]app.User, error) {
	out := make([]app.User, 0, len(u.users))
	for _, user := range u.users {
		out = append(out, user)
	}
	return out, nil
}
