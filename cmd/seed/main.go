package main

import (
	"context"
	"fmt"
	"os"

	"github.com/michaelStillwell/turtle-cms/db"
)

func main() {
	var (
		u     = os.Getenv("DATABASE_URL")
		store = db.MustNew(u)
		ctx   = context.Background()
	)

	fmt.Println("seeding users...")
	SeedUsers(ctx, store)
	fmt.Println("finished seeding users...")
}
