package repository

import (
	"errors"
	"fmt"
)

var (
	// ErrQuery .
	ErrQuery = errors.New("query error")
	// ErrDatabase .
	ErrDatabase = errors.New("database error")
)

// NewErrQuery . 
func NewErrQuery(err error) error {
	return fmt.Errorf("%w : %q", ErrQuery, err)
}

// NewErrDatabase . 
func NewErrDatabase(err error) error {
	return fmt.Errorf("%w : %q", ErrDatabase, err)
}
