package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	st := new(Stack)

	var x int
	for n > 0 {
		fmt.Scan(&x)
		st.Push(x)
		n--
	}
	st.Increment()
	st.Print()
}
