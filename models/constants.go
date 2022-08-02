package models

import "errors"

var (
	ErrIdNotFound = errors.New("id not found")
	ErrCategoryIDNotFound = errors.New("cetgory id not found")
	ErrAuthorIDNotFound = errors.New("author id not found")
)
