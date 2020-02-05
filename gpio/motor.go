package gpio

type Motor interface {
	Forward(int)
	Backword(int)
	Stop()
}
