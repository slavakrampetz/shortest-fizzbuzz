package main

import . "fmt"

func main() {
	p := Print
	for i := 1; i < 101; i++ {
		if i%3 < 1 {
			p("Fizz")
		}
		if i%5 < 1 {
			p("Buzz")
		}
		if i%3*i%5 > 0 {
			p(i)
		}
		p("\n")
	}
}
