package output

var globalPrinter Printer = &VoidPrinter{}

func ChangePrinter(printer Printer) { _ = "STUB: not implemented"; return }

func Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

type Printer interface {
	Printf(format string, args ...interface{})
}

type ConsolePrinter struct{}

func NewConsolePrinter() *ConsolePrinter { _ = "STUB: not implemented"; return nil }

func (c *ConsolePrinter) Printf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

type VoidPrinter struct{}

func (v *VoidPrinter) Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }
