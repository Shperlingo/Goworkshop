package main

import "testing"
import "fmt"
import "strconv"

// go test -cover
//# go tool cover -help
//# go test -coverprofile=c.out
//# go tool cover -html=c.out -o coverage.html
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
