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
	return fmt.Errorf("%q : %e", ErrQuery, err)
}

// NewErrDatabase . 
func NewErrDatabase(err error) error {
	return fmt.Errorf("%q : %e", ErrDatabase, err)
}
