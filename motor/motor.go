package motor

import "github.com/stianeikeland/go-rpio/v4"

type Motor struct {
	fwd rpio.Pin
	bw  rpio.Pin
	pwm rpio.Pin
}

func NewMotor(fwd, bw, pwm rpio.Pin) *Motor {
	fwd.Output()

	bw.Output()

	pwm.Mode(rpio.Pwm)
	pwm.Freq(200)
	pwm.DutyCycle(0, 32)

	return &Motor{fwd, bw, pwm}
}

func (m *Motor) Forward(spreed int) {
	m.bw.Low()
	m.fwd.High()
	m.pwm.DutyCycle(32, 32)
}

func (m *Motor) Backword(spreed int) {
	m.fwd.Low()
	m.bw.High()
	m.pwm.DutyCycle(32, 32)
}

func (m *Motor) Stop() {
	m.fwd.Low()
	m.bw.Low()
}
