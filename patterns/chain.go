package patterns

import "fmt"

// Handler представляет интерфейс обработчика.
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request int)
}

// ConcreteHandler представляет конкретный обработчик.
type ConcreteHandler struct {
	nextHandler Handler
}

func (c *ConcreteHandler) SetNext(handler Handler) {
	c.nextHandler = handler
}

// HandleRequest обрабатывает запрос или передает его следующему обработчику.
func (c *ConcreteHandler) HandleRequest(request int) {
	fmt.Printf("ConcreteHandler handles request: %d\n", request)
	if c.nextHandler != nil {
		c.nextHandler.HandleRequest(request)
	}
}
