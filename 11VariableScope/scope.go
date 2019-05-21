package main

import "fmt"

var xx1 int = 7
var xx2 int = 8
var xx3 int = 9

func f1() {
	var xx1 int = 100
	fmt.Println("xx1: ", xx1, " xx2: ", xx2, " xx3: ", xx3)
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("xx1: ", xx1, " xx2: ", xx2, " xx3: ", xx3)
	var xx2 int = 200
	fmt.Println("xx1: ", xx1, " xx2: ", xx2, " xx3: ", xx3)
	f1()
	fmt.Println("xx1: ", xx1, " xx2: ", xx2, " xx3: ", xx3)
	{
		fmt.Println("xx1: ", xx1, " xx2: ", xx2, " xx3: ", xx3)
		var xx3 int = 300
		fmt.Println("xx1: ", xx1, " xx2: ", xx2, " xx3: ", xx3)
	}
	fmt.Println("xx1: ", xx1, " xx2: ", xx2, " xx3: ", xx3)
	f1 := 7
	fmt.Println(f1)
	// # f1() // you don't want that !!!
}

/*

Func <name>(<name> <type>,…) <return type> {
Return <whatever>
}
The same shorting of x,y type is OK here too
Can have multiple return values
If the return value(s) are named you can assign them along the function and just return “naked” without any value. It is not recommended in long functions because it diminish readability.
The second return value is usually am error object
Go does not have exceptions !!!

Also demonstrate passing parameters by value and ref after functions

Println(x,*x,&x ect.)

*/
