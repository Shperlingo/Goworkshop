package main

import "fmt"

func main() {
	haim := int(5)
	fmt.Println(haim)
	fmt.Println(moris(5, 5))
	fmt.Println(moris2())
}

func moris(num1, num2 int) int {
	return num1 + num2
}
func moris2() (int, int) {
	return 5, 6
}
