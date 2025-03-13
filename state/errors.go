package state

import (
	"errors"
)

var (
	// ErrStateItemNotFound is returned when an item is not found in the state.
	ErrStateItemNotFound = errors.New("state item not found")
	// ErrTypeAssertion is returned when an item is not of the expected type.
	ErrTypeAssertion = errors.New("item is not of type")
)
