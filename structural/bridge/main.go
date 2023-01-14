package bridge

import "fmt"

type Printer interface {
	PrintFile()
}

type Computer interface {
	Print()
	SetPrinter()
}

type Windows struct {
	printer Printer
}
type Mac struct {
	printer Printer
}

//printer

type Hp struct {
}
type Epson struct {
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}
func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func Run() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()

}
