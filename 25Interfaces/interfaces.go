package main

import "fmt"

type shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}
type Rectangel struct {
	x, y float64
}
type Tryangle struct {
	b, h float64
}

func (c Circle) area() float64 {
	return c.r * c.r * 22 / 7
}
func (r Rectangel) area() float64 {
	return r.x * r.y
}
func (t Tryangle) area() float64 {
	return t.b * t.h / 2
}

// CalcArea haha
func CalcArea(s shape) float64 {
	// calc hahaha
	return s.area()
}

func main() {
	fmt.Println("Hello World")
	r := Rectangel{x: 7, y: 8}
	c := Circle{x: 7, y: 8, r: 7}
	t := Tryangle{b: 6, h: 8}
	fmt.Println(CalcArea(r))
	fmt.Println(CalcArea(c))
	fmt.Println(CalcArea(t))

	fmt.Printf("\n")
	var inn interface{}
	//	inn = 8
	//inn = true
	inn = 6.7
	switch inn.(type) {
	case int:
		fmt.Println("int")
		fmt.Println(inn.(int))
	case float64:
		fmt.Println("f64")
		fmt.Println(inn.(float64))
	case bool:
		fmt.Println("bool")
		fmt.Println(inn.(bool))
	default:
		fmt.Println("stam")
	}
}

/*

Type <name> interface {
     <fname>(<fparams>) fretval
     . . .
}
Interface type

Effective go interfaces

Google interfaces


*/
