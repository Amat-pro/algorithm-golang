package main

import "fmt"

func main() {

	a := 5.0
	b := 2.0
	fmt.Println(a / b) // output: 2.5
	fmt.Println(5 / 2) // output: 2

	str := "sss"
	strPointer(str)  // output: pppp
	fmt.Println(str) // output: sss

}

func strPointer(str string) {
	str = "pppp"
	fmt.Println(str)
}
