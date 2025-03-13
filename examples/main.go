package main

import (
	"fmt"

	"github.com/konapun/statekit/state"
)

type NamedModel struct {
	Name string
}

func (n *NamedModel) Key() string {
	return "named"
}

func (n *NamedModel) Clone() state.Model {
	return &NamedModel{n.Name}
}

type AgedModel struct {
	Age int
}

func (a *AgedModel) Key() string {
	return "aged"
}

func (a *AgedModel) Clone() state.Model {
	return &AgedModel{a.Age}
}

func main() {
	// Create a new state with a TestModel
	testState := state.NewState(&NamedModel{Name: "name"}, &AgedModel{Age: 10})

	// Access the TestModel through the ma<ager
	accessor, err := state.AccessorFor[*NamedModel](testState, "named")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Register an observer
	observer := state.NewRuntimeObserver(func(new *NamedModel, old *NamedModel) {
		fmt.Printf("Observer called: %s -> %s\n", old.Name, new.Name)
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
	err = accessor.Update(func(m *NamedModel) error {
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
