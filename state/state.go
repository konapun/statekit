package state

// State is a struct that represents the state of the application.
type State[T Model[T]] struct {
	items map[string]T
}

// NewState creates a new State with the provided items.
func NewState[T Model[T]](items ...T) *State[T] {
	mappedItems := make(map[string]T)
	for _, item := range items {
		mappedItems[item.Key()] = item
	}
	return &State[T]{
		items: mappedItems,
	}
}

// Get returns the item with the provided key.
func (s *State[T]) Get(key string) (T, error) {
	if item, ok := s.items[key]; ok {
		return item, nil
	}
	var zero T
	return zero, ErrStateItemNotFound
}
