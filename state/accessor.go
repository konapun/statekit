package state

// Accessor is a struct that provides access to the state.
type Accessor[T Model[T]] struct {
	model     T
	observers []Observer[T]
}

// NewAccessor creates a new Accessor for the provided model.
func NewAccessor[T Model[T]](model T) *Accessor[T] {
	return &Accessor[T]{
		model:     model,
		observers: make([]Observer[T], 0),
	}
}

// RegisterObserver adds an observer to the list of observers that will be notified when the model changes.
func (a *Accessor[T]) RegisterObserver(observer Observer[T]) {
	a.observers = append(a.observers, observer)
}

// Query returns a copy of the model, so that all changes to the model are isolated and forced through the Update method in order to notify observers.
func (a *Accessor[T]) Query() T {
	return a.model.Clone()
}

// Update modifies the model using the provided modifier function and notifies all observers of the change.
func (a *Accessor[T]) Update(modifier func(T) error) error {
	before := a.model.Clone()
	if err := modifier(a.model); err != nil {
		return err
	}
	after := a.model.Clone()

	a.notifyObservers(after, before)
	return nil
}

func (a *Accessor[T]) notifyObservers(new, old T) {
	for _, observer := range a.observers {
		observer.Update(new, old)
	}
}
