package state

import (
	"errors"
)

var (
  // ErrStateItemNotFound is returned when an item is not found in the state.
	ErrStateItemNotFound = errors.New("state item not found")
)
