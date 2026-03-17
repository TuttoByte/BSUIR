package contracts

type TWindow interface {
	Print(msg string)
	PrintWithWait(msg string) string
	Show()
}
