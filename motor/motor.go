package motor

type Motor interface {
	Forward(int)
	Backword(int)
	Stop()
}
