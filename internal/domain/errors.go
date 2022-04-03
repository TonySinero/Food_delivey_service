package domain

import "errors"

var (
	ErrCustomerNotFound = errors.New("customer doesn't exists")
)
