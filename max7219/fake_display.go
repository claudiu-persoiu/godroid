package max7219

type FakeDisplay struct {
}

func NewFakeDisplay() *FakeDisplay {
	return &FakeDisplay{}
}

func (f *FakeDisplay) Show(text string) error {
	return nil
}
