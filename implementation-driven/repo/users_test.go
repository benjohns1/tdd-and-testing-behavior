package repo

import (
	app "github.com/benjohns1/tdd-and-testing-behavior/implementation-driven"
	"reflect"
	"testing"
)

func TestUser_Save(t *testing.T) {
	type args struct {
		user app.User
	}
	tests := []struct {
		name    string
		repo    Users
		args    args
		want    Users
		wantErr error
	}{
		{
			name: "store a user by name to an empty map",
			args: args{
				user: app.User{Name: "Ender"},
			},
			want: Users{
				map[string]app.User{
					"Ender": {Name: "Ender"},
				},
			},
		},
		{
			name: "store a new user by name to a map that already has a user",
			repo: Users{
				map[string]app.User{
					"Ender": {Name: "Ender"},
				},
			},
			args: args{
				user: app.User{Name: "Valentine"},
			},
			want: Users{
				map[string]app.User{
					"Ender":     {Name: "Ender"},
					"Valentine": {Name: "Valentine"},
				},
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
			if !reflect.DeepEqual(tt.repo, tt.want) {
				t.Errorf("Save() = %+v, want %+v", tt.repo, tt.want)
			}
		})
	}
}
