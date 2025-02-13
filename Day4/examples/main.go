package main

import (
	"fmt"
)

func f1() {
	fmt.Println("this is begining of f1 func")

	fmt.Println("this is end of f1 func")
}
func f2() {
	fmt.Println("this is begining of f2 func")

	fmt.Println("this is end of f2 func")
}
func f3() {
	fmt.Println("this is begining of f3 func")

	fmt.Println("this is end of f3 func")
}
func main() {
	defer fmt.Println("Start of main")
	f1()
	f2()
	f3()
	fmt.Println("end of main")
}
