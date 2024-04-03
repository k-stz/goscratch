package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName, LastName string
	Age                 int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		firstName, lastName, age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	person := MakePerson(firstName, lastName, age)
	return &person
}

func exerciseEscapeTest() {
	// Compile prgoram with -gcflags="-m" to get result of escape analysis
	// for which values will escape to the heap
	steve := MakePerson("Steve", "Stacker", 55)
	// ezekiel's pointed to value will escape the stack, because after
	// the stack frame pops, it needs to be referred to. So the escape analysis
	// places it on the heap.
	ezekiel := MakePersonPointer("Ezekiel", "Escapist", 44)
	// Steve Stacker will also escape the stack! That's solely due to this Println,
	// as it expects an `..any`. The current Go compiler moves any argument that is
	// of an interface type to the heap.
	fmt.Println(steve)
	fmt.Println(ezekiel)
}

func UpdateSlice(sl []string, str string) {
	sl[len(sl)-1] = str
	fmt.Println(sl)
}

// GrowSlice will _not_ extend the slice, as go is call-by-value
// and slices are implemented as structs of a ptr to the backing array and a len/cap int
// With append the len of the copy is increased and the backing arrays position changed
// upon function return the original slice has the old len value and thus can't see
// the appended change!
func GrowSlice(sl []string, str string) {
	sl = append(sl, str)
	fmt.Println("inside GrowSlice():", sl)
}

func exerciseSliceMutability() {
	slice := []string{"one", "two"}
	UpdateSlice(slice, "replaced")
	fmt.Println(slice)
	GrowSlice(slice, "extended")
	fmt.Println("after GrowSlice():", slice)
}

func exercise3cap() {
	persons := make([]Person, 10_000_000)
	fmt.Println("persons allocated.")
	for i := 0; i < 10_000_000; i++ {
		persons[i] = Person{"Robin", "Hood", 100}
	}
}

func exercise3() {
	var persons []Person
	for i := 0; i < 10_000_000; i++ {
		persons = append(persons, Person{"Robin", "Hood", 100})
	}
}

func main() {
	// exerciseEscapeTest()
	// exerciseSliceMutability

	// Set env var GODEBUG=gctrace=1 to see when gc happens
	// Change env var GOGC to control when gc happens. Default is 100
	// which is every time the heap doubles

	// startTime := time.Now()
	// exercise3()
	// endTime := time.Now()
	// fmt.Println("Function executed in:", endTime.Sub(startTime))
	// Result:
	// GOGC=50 (default): gc cycles=28, executation time=3.9 ~ 5.9 sec
	// GOGC=100 (default): gc cycles=16, executation time=2.9 ~ 3.2 sec
	// GOGC=200: gc cycles=9, executation time=2.3, ~ 2.9 sec
	// GOGC=400: gc cycles=5, executation time=2.25, ~ 2.7 sec
	// GOGC=800: gc cycles=3-4, executation time=1.6, ~ 1.8 sec
	// GOGC=3200: gc cycles=1, executation time=1.6, ~ 2.5 sec

	// Conclusion:
	// When increasing the size at which GC gets triggered, naturally
	// lowrs the number of cycles and speeds up the program as more
	// time is spend in execution and less in the GC thread!

	fmt.Println("start execution")
	startTime := time.Now()
	exercise3cap()
	endTime := time.Now()
	fmt.Println("Function executed in:", endTime.Sub(startTime))
	// Result:
	// GOGC=50 (default): gc cycles=1, executation time=0.4 ~ 0.47 sec
	// GOGC=100 (default): gc cycles=1, executation time=0.39 ~0.4 sec
	// GOGC=400: gc cycles=1, executation time=0.4 sec
	// GOGC=3200: gc cycles=1, executation time=0.4

	// Conclusion:
	// preallocating the slice with make, only takes a single GC cycle
	// Execution time is vastly increased
	// 
	// If you know you are going to need a block of memory, it's best to 
	// allocate it at once and use it. If you can re-use it, all the better. 
	// That's the reason for the slice buffer pattern, also.
}
