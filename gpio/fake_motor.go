package gpio

import "fmt"

type FakeMotor struct {
	name string
}

func NewFakeMotor(name string) *FakeMotor {
	return &FakeMotor{name}
}

func (m *FakeMotor) Forward(spreed int) {
	fmt.Println("Fw: ", m.name, " spreed: ", spreed)
}

func (m *FakeMotor) Backword(spreed int) {
	fmt.Println("Bw: ", m.name, " spreed: ", spreed)
}

func (m *FakeMotor) Stop() {
	fmt.Println("Stop: ", m.name)
}
