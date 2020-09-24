package urlInterfaces

import "errors"

var (
	ErrAlreadyExists = errors.New("new url already exists")
	ErrNotFound = errors.New("old url not found")
)
