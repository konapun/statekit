package state

// Observer is an interface that defines the methods required for an observer.
type Observer[T any] interface {
	Update(old, new T)
}

// RuntimeObserver allows for the creation of an observer at runtime.
type RuntimeObserver[T any] struct {
	update func(new T, old T)
}

// NewRuntimeObserver creates a new RuntimeObserver with the provided update function.
func NewRuntimeObserver[T any](update func(T, T)) *RuntimeObserver[T] {
	return &RuntimeObserver[T]{
		update: update,
	}
}

// Update notifies the observer of a change in the model.
func (o *RuntimeObserver[T]) Update(new T, old T) {
	o.update(new, old)
}
