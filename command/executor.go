package command

// Executor is responsible for execute Command's.
type Executor struct {
	source      string
	destination string
}

// NewExecutor creates a new *Executor.
func NewExecutor(source, destination string) *Executor {
	return &Executor{
		source:      source,
		destination: destination,
	}
}

// Execute all the provided commands, following the original order.
func (e *Executor) Execute(commands ...Command) error {
	for _, command := range commands {
		if err := command(e.source, e.destination); err != nil {
			return err
		}
	}

	return nil
}
