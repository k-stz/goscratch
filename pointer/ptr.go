package main

import "fmt"

func f(px *int) {
	// why won't this change the address in the callee?
	// Because call-by-value means a copy of the pointer gets
	// passed. The copy arg contains the same address as the callee
	// but when it gets changed in px = &x2
	// It is not changed in the callee! Only in this local copy!
	x2 := 20
	px = &x2
}

func update(px *int) {
	// only this truely changes the pointed to value. Why?
	// Because we dereference the address and thus point to the
	// same location as the callee's pointer! 
	*px = 20
}

func updateSlice(s []int) {
	// append won't change the original slice! Even when
	// when the capacity > length
	// why? Because due to pass-by-value, 's' gets a copy of the
	// pointer to the slice.
	// Slice is implemented as a struct:
	// slice {
	//  ptr [x]array
	//  len, cap int	
	// }
	// such that changing its len via append() won't change
	// it in the callee! The underlying array will have a
	// value appended, but the callee slice won't see it
	// as it's len won't be touched
	s = append(s, 66)
	// s[0] will change the shared underlying array!
	s[0] = 44
	fmt.Println("slice inside updateSlice() at the end", s)
}

func main() {
	x := 10
	f(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
	s1 := make([]int, 3, 6)
	s1[0], s1[1], s1[2] = 1, 2, 3
	fmt.Println("before updateSlice()", s1)
	updateSlice(s1)
	fmt.Println("after updateSlice()", s1)
	fmt.Println("=> updateSlice hasn't appended to s1!")
}
