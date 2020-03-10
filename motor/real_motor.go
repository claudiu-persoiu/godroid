package motor

import "github.com/stianeikeland/go-rpio/v4"

type RealMotor struct {
	fwd rpio.Pin
	bw  rpio.Pin
	pwm rpio.Pin
}

func NewRealMotor(fwd, bw, pwm rpio.Pin) *RealMotor {
	fwd.Output()

	bw.Output()

	pwm.Mode(rpio.Pwm)
	pwm.Freq(200)
	pwm.DutyCycle(0, 32)

	return &RealMotor{fwd, bw, pwm}
}

func (m *RealMotor) Forward(spreed int) {
	m.bw.Low()
	m.fwd.High()
	m.pwm.DutyCycle(32, 32)
}

func (m *RealMotor) Backword(spreed int) {
	m.fwd.Low()
	m.bw.High()
	m.pwm.DutyCycle(32, 32)
}

func (m *RealMotor) Stop() {
	m.fwd.Low()
	m.bw.Low()
}
