package main

import (
	"fmt"

	"github.com/konapun/statekit/state"
)

type TestModel struct {
	key  string
	Name string
}

func (t *TestModel) Key() string {
	return t.key
}

func (t *TestModel) Clone() *TestModel {
	return &TestModel{
		key:  t.key,
		Name: t.Name,
	}
}

func (t *TestModel) ToString() string {
	return fmt.Sprintf("Name: %s", t.Name)
}

func main() {
	// Create a new state with a TestModel
	testState := state.NewState(&TestModel{key: "test", Name: "name"})
	manager := state.NewManager(testState)

	// Access the TestModel through the manager
	accessor, err := manager.AccessorFor("test")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Register an observer
	observer := state.NewRuntimeObserver(func(new *TestModel, old *TestModel) {
		fmt.Printf("Observer called: %s -> %s\n", old.ToString(), new.ToString())
	})
	accessor.RegisterObserver(observer)

	// Query the model. Query returns a clone of the model so that the model can't be modified outside of an update.
	model := accessor.Query()
	fmt.Printf("Model name: %s\n", model.Name)
	// Update the model name to see it doesn't persist to state
	model.Name = "updated"
	model = accessor.Query()
	fmt.Printf("Persisted model name: %s\n", model.Name)

	// Update the model
	err = accessor.Update(func(m *TestModel) error {
		m.Name = "updated"
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Query the updated model
	model = accessor.Query()
	fmt.Printf("Updated model name: %s\n", model.Name)
}
