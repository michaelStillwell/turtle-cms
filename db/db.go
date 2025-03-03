package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/michaelStillwell/turtle-cms/internal/repository"
	_ "github.com/tursodatabase/go-libsql"
)

type Store struct {
	db   *sql.DB
	Repo *repository.Queries
}

func MustNew(url string) *Store {
	db, err := sql.Open("libsql", url)
	if err != nil {
		panic(fmt.Sprintf("failed to open db %s", err))
	}

	queries := repository.New(db)
	return &Store{
		db:   db,
		Repo: queries,
	}
}

func (s *Store) CreateTenant(ctx context.Context, name string) error {

	_, err := s.Repo.CreateTenant(ctx, name)
	if err != nil {
		return err
	}

	// _, err := s.db.Exec(fmt.Sprint(""))

	return nil
}
