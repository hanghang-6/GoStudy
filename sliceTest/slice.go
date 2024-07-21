package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func PrintSliceStructural(s *[]int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))
	fmt.Printf("slice struct:%+v,slice is %vn\n", ss, s)
}

func case2(s []int) {
	s = s[1:]
	PrintSliceStructural(&s)
}
func case1(s []int) {
	s = s[1:3]
	PrintSliceStructural(&s)
}
func case3(s []int) {
	s = s[len(s)-1:]
	PrintSliceStructural(&s)
}
func case4(s []int) {
	ss := s[1:]
	PrintSliceStructural(&ss)
}
func main() {

	s := make([]int, 9)
	case2(s)
	case1(s)
	case3(s)
	case4(s)
	PrintSliceStructural(&s)
}
