package repository

import "database/sql"

type Repo struct {
	db *sql.DB
}

type RepoInterface interface {
	BookRepo
	CategoryRepo
	AuthorRepo
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}
