package patterns

import "fmt"

// Command представляет интерфейс команды.
type Command interface {
	Execute()
	Undo()
}

// Receiver представляет объект, над которым выполняется команда.
type Receiver struct {
	IsOn bool
}

func (r *Receiver) TurnOn() {
	r.IsOn = true
	fmt.Println("Receiver is ON")
}

func (r *Receiver) TurnOff() {
	r.IsOn = false
	fmt.Println("Receiver is OFF")
}

// ConcreteCommand представляет конкретную команду для включения ресивера.
type ConcreteCommandOn struct {
	receiver *Receiver
}

func NewConcreteCommandOn(receiver *Receiver) *ConcreteCommandOn {
	return &ConcreteCommandOn{receiver: receiver}
}

func (c *ConcreteCommandOn) Execute() {
	c.receiver.TurnOn()
}

func (c *ConcreteCommandOn) Undo() {
	c.receiver.TurnOff()
}

// ConcreteCommandOff представляет конкретную команду для выключения ресивера.
type ConcreteCommandOff struct {
	receiver *Receiver
}

func NewConcreteCommandOff(receiver *Receiver) *ConcreteCommandOff {
	return &ConcreteCommandOff{receiver: receiver}
}

func (c *ConcreteCommandOff) Execute() {
	c.receiver.TurnOff()
}

func (c *ConcreteCommandOff) Undo() {
	c.receiver.TurnOn()
}

// Invoker представляет объект, который отправляет команду.
type Invoker struct {
	onCommand  Command
	offCommand Command
}

func NewInvoker(onCommand, offCommand Command) *Invoker {
	return &Invoker{
		onCommand:  onCommand,
		offCommand: offCommand,
	}
}

func (i *Invoker) TurnOn() {
	i.onCommand.Execute()
}

func (i *Invoker) TurnOff() {
	i.offCommand.Execute()
}
