package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {}
func (m MultiFunctionPrinter) Fax(d Document)   {}
func (m MultiFunctionPrinter) Scan(d Document)  {}

type OldFashionedPrinter struct{}

func (o OldFashionedPrinter) Print(d Document) {
	// OK
}
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Depreceated...
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// Interface Segregation Principle

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {}

type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	panic("implement me")
}

func (p Photocopier) Print(d Document) {
	panic("implement me")
}

type MultiDevice interface {
	Printer
	Scanner
	// Fax
}

func main() {

}
