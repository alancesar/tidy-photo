package command

type Executor struct {
	source      string
	destination string
}

func NewExecutor(source, destination string) *Executor {
	return &Executor{
		source:      source,
		destination: destination,
	}
}

func (e *Executor) Execute(commands ...Command) error {
	for _, command := range commands {
		if err := command(e.source, e.destination); err != nil {
			return err
		}
	}

	return nil
}
