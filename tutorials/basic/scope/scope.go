/*
	Scope
		- Scope is the accessibility of variables, functions, and objects in some particular part of your code.
		- In other words, scope determines the visibility of variables and other resources in areas of your code.
*/
package main

import "fmt"

func main() {
	closure()
}

/*
	CLOSURE
*/
func newCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func closure() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
}
