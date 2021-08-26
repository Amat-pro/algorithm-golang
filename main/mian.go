package main

import (
	"fmt"
)

func main() {

	a := 5.0
	b := 2.0
	fmt.Println(a / b) // output: 2.5
	fmt.Println(5 / 2) // output: 2

	// warn: see the explains in strPointer
	str := "sss"
	strPointer(str)  // output: pppp
	fmt.Println(str) // output: sss

	array := [2]int{3, 4}
	arrayPointer(array) // output: [1 2]
	fmt.Println(array)  // output: [3 4]

	var slice []int
	slice = []int{3, 5, 6}
	/**
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array))
	*/
	slicePointer(slice) // output: [1 2 6]
	fmt.Println(slice)  // output: [1 2 6]

	inter0 := interImpl0{}
	inter := interImpl{}
	// if inter0 is value implement then you can deliver pointer or value
	// if inter0 is pointer implement then you must deliver pointer
	inter.inter(&inter0)

}

func strPointer(str string) {
	// warn: this str is a parameter, not an argument
	// the follow operation make it point another address
	str = "pppp"
	fmt.Println(str)
}

func arrayPointer(array [2]int) {
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
}

func slicePointer(slice []int) {
	slice[0] = 1
	slice[1] = 2
	fmt.Println(slice)
}

type inter0 interface {
	inter0()
}

type inter interface {
	inter(inter0 inter0)
}

type interImpl struct {
}

type interImpl0 struct {
}

func (interImpl *interImpl) inter(inter02 inter0) {

}

func (interImpl0 interImpl0) inter0() {

}

/**
func (interImpl0 *interImpl0) inter() {

}
*/
