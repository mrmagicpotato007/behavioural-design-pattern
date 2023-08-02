package command

type Element interface {
	Press()
}

type button struct {
	command Command
}

func (b *button) Press() {
	b.command.execute()
}

func NewButton(command Command) *button {
	return &button{command: command}
}


func Executor(element Element)  {
	element.Press()
}