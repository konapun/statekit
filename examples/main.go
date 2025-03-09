package main

import (
	"fmt"

	"github.com/konapun/statekit/state"
)

type TestModel struct {
	key string
}

func (t *TestModel) Key() string {
	return t.key
}

func (t *TestModel) Clone() *TestModel {
	return &TestModel{
		key: t.key,
	}
}

func main() {
	// Create a new state with a TestModel
	testState := state.NewState(&TestModel{key: "test"})
	manager := state.NewManager(testState)

	// Access the TestModel through the manager
	accessor, err := manager.AccessorFor("test")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Register an observer
	observer := state.NewRuntimeObserver(func(new *TestModel, old *TestModel) {
		fmt.Printf("Observer called: %s -> %s\n", old.Key(), new.Key())
	})
	accessor.RegisterObserver(observer)

	// Query the model
	model := accessor.Query()
	fmt.Printf("Model key: %s\n", model.Key())

	// Update the model
	err = accessor.Update(func(m *TestModel) error {
		m.key = "updated"
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Query the updated model
	model = accessor.Query()
	fmt.Printf("Updated model key: %s\n", model.Key())
}
