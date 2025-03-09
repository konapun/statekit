package state

// Model is an interface that defines the methods required for a model, where a model is a struct that represents a piece of state.
type Model[T any] interface {
	Key() string
	Clone() T
}
