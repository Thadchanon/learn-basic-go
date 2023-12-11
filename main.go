package main

import (
	"fmt"
	"math"
	// workshops "github.com/Thadchanon/learn-basic-go/workshops" // for workshops
)

func main() {
	// msg := "Hello Go!"
	// age := 24
	// price := 22.522
	// check := true

	// var rn rune = 'a'

	// fmt.Printf("type: %T -- msg: %#v\n", msg, msg)       // fmt.Printf("msg: %s\n", msg)
	// fmt.Printf("type: %T -- age: %#v\n", age, age)       // fmt.Printf("age: %d\n", age)
	// fmt.Printf("type: %T -- price: %#v\n", price, price) // fmt.Printf("price: %.2f\n", price)
	// fmt.Printf("type: %T -- check: %#v\n", check, check) // fmt.Printf("check: %t\n", check)
	// fmt.Printf("type: %T -- rune: %#v\n", rn, rn)
	// fmt.Printf("rune: %c\n", rn) // f = format, %c =  print character, \n = new line

	num := 34

	if num == 34 && (num > 30 || num < 39) {
		fmt.Println("Yes!! it's Thirty four.")
	} else if num == 35 {
		fmt.Println("It's Thirty five.")
	} else {
		fmt.Println("No!! it's not Thirty four.")
	}

	if num%2 == 0 {
		fmt.Println("It's even number.")
	} else {
		fmt.Println("It's odd number.")
	}

	lim := 225.0

	if v := math.Pow(30, 2); v < lim { // Pow uses calculations using exponents, v scope in if
		fmt.Println("x power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
	}

	// workshops.PrintfWs() // for workshop 1
	// workshops.PrintlnWs() // for workshop 2
	// workshops.IfElse() // for workshop 3
}
