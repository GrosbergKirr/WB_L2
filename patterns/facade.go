package patterns

import "fmt"

type Engine struct {
}

func (engine *Engine) Start() {
	fmt.Println("start engine")
}

type Transmission struct{}

func (transmission *Transmission) First() {
	fmt.Println("check transmission")
}

type Battery struct {
}

func (battery *Battery) Check() {
	fmt.Println("charge check battery")
}

type Car struct {
	engine  *Engine
	trans   *Transmission
	battery *Battery
}

func NewCar() *Car {
	fmt.Println(" New car created")
	return &Car{&Engine{},
		&Transmission{},
		&Battery{},
	}
}

func (c *Car) LetsGo() {
	c.engine.Start()
	c.battery.Check()
	c.trans.First()
}
