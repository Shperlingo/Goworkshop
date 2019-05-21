package main

import "fmt"

const n1 = iota
const n2 = iota * 10
const (
	_  = iota
	n3 = iota
	n4 = iota
	n5 = iota * 10
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
}

/*


 */
