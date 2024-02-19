package main

import (
	"fmt"
	"time"
)

func variables() {
	var a int
	a = 1
	var b int = 3
	var c = 5
	d := 6

	fmt.Printf("%d %d %d %d", a, b, c, d)
}

func constants() {
	const n = 123456
	// A number constant doesn't have a type until you give it one
	// Or until you use it where a type must be inferred
}

func loops() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Bruh!")
		break
	}
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func conditionals() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// No  ternary!! Boo, you suck
}

func switch_function() {
	// Switches are pretty good
	// You can match on equality
	// Or just have case statements that evaluate to a bool

	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		println("Derp")
	}

	// Multiple cases comma deliniated
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Check this out
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}

func arrays() {
	// This array will have all zeros (default value of an int)
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	// Declare like this
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
}

func slices_n_stuf() {
	// Like arrays, just don't set the length
	// Use make keyword
	s := make([]string, 3) // 3 is the capacity
	println(s)

	// Slices are kinda awful to be honest
	s1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	s2 := s1[3:6]
	fmt.Println(s2) //[d e f] - inclusive, exclusive

	// But we can move the end of the range of s2 up, but not the start range down... weird af
	fmt.Println(len(s2)) // 3
	fmt.Println(cap(s2)) // 6  from the staring point of the slice in the underlying array to the end of the array (d to i)
	s2 = s2[:6]
	fmt.Println(s2) // [d e f g h i]
}

func maps_n_stuff() {
	m := make(map[string]int)

	m["four"] = 4
	m["three"] = 3
	fmt.Println(m["three"])
	delete(m, "three")
	fmt.Println("Map: ", m)

	val, ok := m["five"]
	fmt.Println("Found?", ok, "val:", val)
}

func range_n_stuff() {
	// Range iterates over stuff
	nums := []int{1, 2, 3, 4, 56, 7, 8, 9}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, c := range "this is my string!!%$£¬|" {
		fmt.Println(i, c)
	}

}

func multi_return() (string, string) {
	a := "you"
	b := "Gay!!!!"
	return a, b
}

func varadic_func(nums ...int) {
	// nums is iterable
	for n := range nums {
		fmt.Print(n)
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//variables()
	//slices_n_stuff()
	//maps_n_stuff()
	//range_n_stuff()
	//varadic_func(1, 2, 3)
	//x := []int{2, 3, 65, 9, 8}
	//varadic_func(x...) // The dots expand the slice into args
	//myClosure := intSeq()
	//fmt.Println(myClosure())
	//fmt.Println(myClosure())
	//fmt.Println(myClosure())

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(11))

}
