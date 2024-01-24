package cmd

import (
	"app"
	"app/repo"
)

func main() {
	userRepo := repo.User{}
	userRepo.Save(app.User{Name: "Bob"})
}
