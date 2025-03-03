package main

import (
	"context"
	"fmt"

	"github.com/michaelStillwell/turtle-cms/db"
	"github.com/michaelStillwell/turtle-cms/internal/repository"
)

const password = "$argon2id$v=19$m=65536,t=1,p=8$bcS5PRQuGiZGFDSswumzng$eJjuikQ7e3dORUxIPFTItAs2TCzZgjhYnCBTj6Yc0lQ"

func SeedUsers(ctx context.Context, store *db.Store) {
	for _, u := range users {
		newUser, err := store.Repo.CreateUser(ctx, u)
		if err != nil {
			fmt.Printf("new user err: %v\n", err)
		} else {
			fmt.Printf("new user: %v\n", newUser.UserID)
		}
	}
}

var users = [...]repository.CreateUserParams{
	{
		FirstName: "My",
		LastName:  "Daddy",
		Email:     "asdf@asdf.com",
		Password:  password,
	},
}
