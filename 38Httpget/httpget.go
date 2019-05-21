package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("*********************************************")
	res, _ := http.Get("http://www.seladeveloperpractice.com//speakers/#Gad_J._Meir")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
	fmt.Println("*********************************************")
	res, err := http.Get("http://www.seladeveloperpractice.com//speakers/#Gad_J._Meir")
	if err != nil {
		fmt.Println("error on get")
		return
	}
	page, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("error on readall")
		return
	}
	fmt.Printf("%s", page)
}
