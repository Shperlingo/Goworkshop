package main

import "fmt"

func hello(name string) string {
	return "Hello " + name
}

func doSum(num1, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(hello("kuku"))
	fmt.Println(doSum(5, 6))
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
you can give names to the return types and assign to them and just return
But it is not considered a good practice
you can pass multiple params of the same type as f(v ...type) it is received as a fixed array
talk about it after we learn on arrayes
you can pass array as list of elements using arrayname...
if you want to reveive slice f(s []type)

a parameter of a function can be a function

function can be called recursivly



*/
