package state_test

import "github.com/konapun/statekit/state"

type TestModel struct {
	key    string
	String string
	Int    int
}

func NewTestModel(key string) *TestModel {
	return &TestModel{
		key: key,
	}
}

func (t *TestModel) Key() string {
	return t.key
}

func (t *TestModel) Clone() state.Model {
	return &TestModel{
		key:    t.key,
		String: t.String,
		Int:    t.Int,
	}
}
