package example

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type Photocopier interface {
	Photocopy()
}

// //////////////////////////////
type allInOnePrinter struct{}

func (p allInOnePrinter) Print() {
	fmt.Println("Printing...")
}
func (p allInOnePrinter) Scan() {
	fmt.Println("Scanning...")
}
func (p allInOnePrinter) Photocopy() {
	fmt.Println("Photocopying...")
}

// //////////////////////////////
type StandardPrinter struct{}

func (p StandardPrinter) Print() {
	fmt.Println("printing...")
}

// ///////////////////////////////
type SimpleScanner struct{}

func (p SimpleScanner) Scan() {
	fmt.Println("Scannig...")
}
