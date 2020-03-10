package max7219

type Display interface {
	ShowText(string) error
}
