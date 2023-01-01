package main

import "fmt"

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}
func (p Pentagon) CalcAera() {
	fmt.Println("Pentagon area =", (5^2)+1)
}

func main() {
	var p Polygons = Pentagon(50)
	p.Perimeter()
	// The below throws CE because Pentagon interface doesn't has only Perimeter method.
	// p.NumberOfSide()
	var o Pentagon = p.(Pentagon)
	o.NumberOfSide()
	o.CalcAera() // this can only by called by Pentagon object.

	var obj Object = Pentagon(50)
	obj.NumberOfSide()
	// Similar reason like above.
	// obj.Perimeter()
	var pent Pentagon = obj.(Pentagon)
	pent.Perimeter()
}
