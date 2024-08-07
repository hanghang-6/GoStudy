package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func test(s []int) {
	PrintSliceStructural(&s)
}

func PrintSliceStructural(s *[]int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))
	fmt.Printf("slice struct:%+v,slice is %vn\n", ss, s)
}

func case1(c []int) {
	c[1] = 1
	PrintSliceStructural(&c)
}

func case2(c []int) {
	c = append(c, 5)
	c[1] = 1
	PrintSliceStructural(&c)
}
func main() {
	
	s := make([]int, 5, 10)
	case2(s)
	PrintSliceStructural(&s)
}
