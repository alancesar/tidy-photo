package command

import "testing"

func TestExecutor_Execute(t *testing.T) {
	var called bool
	cmd := func(_, _ string) error {
		called = true
		return nil
	}

	err := NewExecutor("source", "dest").Execute(cmd)
	if err != nil {
		t.Error(err)
	}

	if !called {
		t.Error("command was not called")
	}
}
