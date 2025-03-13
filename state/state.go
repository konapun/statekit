package state

// State is a struct that represents the state of the application.
type State struct {
	items map[string]Model
}

// NewState creates a new State with the provided items.
func NewState(items ...Model) *State {
	mappedItems := make(map[string]Model)
	for _, item := range items {
		mappedItems[item.Key()] = item
	}
	return &State{
		items: mappedItems,
	}
}

// Get returns the item with the provided key.
func (s *State) Get(key string) (Model, error) {
	if item, ok := s.items[key]; ok {
		return item, nil
	}
	return nil, ErrStateItemNotFound
}

func AccessorFor[T Model](state *State, key string) (*Accessor[T], error) {
	item, err := state.Get(key)
	if err != nil {
		return nil, err
	}

	typedItem, ok := item.(T)
	if !ok {
		return nil, ErrTypeAssertion
	}

	accessor := NewAccessor(typedItem)
	return accessor, nil
}
