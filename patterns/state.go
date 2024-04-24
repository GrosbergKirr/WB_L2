package patterns

// Состояние вентилятора
type FanState interface {
	TurnOn() string
	TurnOff() string
}

// Вентилятор
type Fan struct {
	state FanState
}

func (f *Fan) SetState(state FanState) {
	f.state = state
}

func (f *Fan) TurnOn() string {
	return f.state.TurnOn()
}

func (f *Fan) TurnOff() string {
	return f.state.TurnOff()
}

// Вентилятор включен на первой скорости
type LowSpeedState struct{}

func (s *LowSpeedState) TurnOn() string {
	return "Вентилятор включен на первой скорости"
}

func (s *LowSpeedState) TurnOff() string {
	return "Вентилятор выключен"
}

// Вентилятор включен на второй скорости
type HighSpeedState struct{}

func (s *HighSpeedState) TurnOn() string {
	return "Вентилятор включен на второй скорости"
}

func (s *HighSpeedState) TurnOff() string {
	return "Вентилятор выключен"
}
