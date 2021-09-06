package main

import (
	"fmt"
	"sync"
)

func main() {

	/*
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

		fmt.Println(reflect.TypeOf(slice))
		fmt.Println(reflect.TypeOf(array))

		slicePointer(slice) // output: [1 2 6]
		fmt.Println(slice)  // output: [1 2 6]

		inter0 := interImpl0{}
		inter := interImpl{}
		// if inter0 is value implement then you can deliver pointer or value
		// if inter0 is pointer implement then you must deliver pointer
		inter.inter(&inter0)

	*/
	/*
		fmt.Printf("%b \n", -2)
		fmt.Printf("%d \n", -2 << 22)
		// 逻辑右移
		fmt.Printf("%b \n", 2 >> 3)
		fmt.Printf("%b \n", -1)
		fmt.Printf("%b \n", -2 >> 1)
		fmt.Printf("%d \n", -2 >> 4)
		// 算数右移->负数符号位为1 发生右移后最高位肯定为1 符号位补1 （打印的第一位： 1： -  0： +）
		fmt.Printf("%b \n", -20 >> 8)

		var i uint8
		fmt.Println(unsafe.Sizeof(i))
		i = 255
		fmt.Printf("%b \n", i >> 2)

		fmt.Printf("%b \n", 9.2233720368548e+18 >> 3)

	*/

	/*

		s := make([]int, 3, 5)
		// append
		s[0] = 1
		s = append(s, 1, 2, 3)
		s = append(s, 4)
		fmt.Println(s)
		fmt.Println(s[0:2])

		var s1 []int
		_ = s1[:]
		// panic
		// fmt.Println(s1[0])

	*/

	/*

		test1()
		// test2()
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("test recover")
				fmt.Println(err)
			}
		}()

	*/

	p := &sync.Pool{
		New: func() interface{} {
			return 2
		},
	}
	a := p.Get().(int)
	p.Put(1)
	p.Put(3)
	// runtime.GC()
	b := p.Get().(int)
	c := p.Get().(int)
	p.Put(c)
	// d := p.Get().(int)
	d := p.Get().(int)
	// runtime.GC()
	fmt.Println(a, b, c, d)
	// runtime.GC()
	fmt.Println(a, b, c, d)

}

func test1() {
	defer func() {
		fmt.Println("test1 defer1")
		if err := recover(); err != nil {
			fmt.Println("test1 recover")
			fmt.Println(err)
		}
	}()
	defer func() {
		fmt.Println("test1 defer2")
		panic("test1 defer panic")
	}()
	fmt.Println("test1")
	panic("test1 panic")
}

func test2() {
	defer func() {
		fmt.Println("test2 defer1")
	}()
	defer func() {
		fmt.Println("test2 defer2")
	}()
	fmt.Println("test2")
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
