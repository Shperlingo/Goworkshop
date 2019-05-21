package main

import "testing"
import "fmt"
import "strconv"

// name the file whatece_test.go
// import testing
// write the tests
// test function name must start with Test and after it capital letter anything
// go help testfunc
// go test
// go test -v
// go test -run '.*01'
// go test -v -run '.*0[1-2]'
func TestA01(t *testing.T) {
	result := Hello("kuku")
	fmt.Println("1")
	t.Log("2")
	if result != "Hello kuku" {
		t.Error("Expected Hello kuku got: " + result)
	}
}

func TestA02(t *testing.T) {
	result := DoSum(2, 5)
	if result != 7 {
		t.Error("Expected 2+5 == 7 got: " + strconv.Itoa(result))
	}
}

func TestA03(t *testing.T) {
	result := DoSum(2, 5)
	if result != 8 {
		t.Error("simulated error Expected 2+5 == 7 got: " + strconv.Itoa(result))
	}
}
func TestA04(t *testing.T) {
	result := aoSum(4, 8)
	if result != 12 {
		t.Error("Expected 4+8 == 7 got: " + strconv.Itoa(result))
	} else {
		t.Fail()
	}
}
