package main

import (
	"fmt"
)

type sliceHeader struct {
	Length        int
	ZerothElement *byte
}

func main() {

	test5()
}

func test5() {
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = append(slice, 5, 6, 7, 8)
	fmt.Println(slice)

	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{11, 22, 33, 44, 55, 66}

	fmt.Println("before slice1 is ", slice1)
	slice1 = append(slice1, slice2...)
	fmt.Println("after slice1 is ", slice1)
}
func test4() {
	slice := []byte{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("before slice is ", slice)
	PtrSubtractOneFromLength(&slice)
	fmt.Println("after slice is ", slice)
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}
func test3() {
	slice := []byte{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Before: len(slice) =", len(slice))
	fmt.Println(slice)

	newSlice := SubtractOneFromLength(slice)

	fmt.Println(slice)
	fmt.Println(newSlice)

	fmt.Println("After:len(slice) = ", len(slice))
	fmt.Println("After:len(newSlice = ", len(newSlice))
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}
func test2() {
	var buffer [255]byte

	slice := buffer[10:30]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println("before", slice)

	AddOneToEachElement(slice)

	fmt.Println("after", slice)

	fmt.Println(buffer)
}
func test1() {
	var buffer [255]byte
	buffer[0] = 255

	fmt.Println(buffer)

	var slice []byte = buffer[100:150]

	aa := buffer[10:20]
	fmt.Println(aa)

	slice1 := sliceHeader{
		Length:        50,
		ZerothElement: &buffer[100],
	}

	fmt.Println(slice1)

	slice2 := sliceHeader{
		Length:        10,
		ZerothElement: &buffer[101],
	}

	fmt.Println(slice2)

	slice = slice[10 : len(slice)-1]
	fmt.Println(slice)

	fmt.Println(len(buffer))
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
