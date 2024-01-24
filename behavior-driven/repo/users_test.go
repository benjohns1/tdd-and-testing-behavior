package repo_test

import (
	app "github.com/benjohns1/tdd-and-testing-behavior/behavior-driven"
	"github.com/benjohns1/tdd-and-testing-behavior/behavior-driven/repo"
	"reflect"
	"testing"
)

func TestUser_Save(t *testing.T) {
	type args struct {
		user app.User
	}
	tests := []struct {
		name      string
		repo      repo.Users
		args      args
		wantErr   error
		wantUsers []app.User
	}{
		{
			name: "should save a user to an empty repo",
			args: args{
				user: app.User{Name: "Ender"},
			},
			wantUsers: []app.User{{Name: "Ender"}},
		},
		{
			name: "should save a user to a repo that already has a user in it",
			repo: func() repo.Users {
				r := repo.Users{}
				if err := r.Save(app.User{
					Name: "Ender",
				}); err != nil {
					t.Fatal(err)
				}
				return r
			}(),
			args: args{
				user: app.User{Name: "Valentine"},
			},
			wantUsers: []app.User{
				{Name: "Ender"},
				{Name: "Valentine"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.repo.Save(tt.args.user)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := tt.repo.GetAll()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.wantUsers) {
				t.Errorf("After Save(), GetAll() = %+v, wantUsers %+v", got, tt.wantUsers)
			}
		})
	}
}
