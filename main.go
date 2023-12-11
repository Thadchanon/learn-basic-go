package main

import (
	"fmt"
	// workshops "github.com/Thadchanon/learn-basic-go/workshops" // for workshops
)

func main() {
	msg := "Hello Go!"
	age := 24
	price := 22.522
	check := true

	var rn rune = 'a'

	fmt.Printf("type: %T -- msg: %#v\n", msg, msg)       // fmt.Printf("msg: %s\n", msg)
	fmt.Printf("type: %T -- age: %#v\n", age, age)       // fmt.Printf("age: %d\n", age)
	fmt.Printf("type: %T -- price: %#v\n", price, price) // fmt.Printf("price: %.2f\n", price)
	fmt.Printf("type: %T -- check: %#v\n", check, check) // fmt.Printf("check: %t\n", check)
	fmt.Printf("type: %T -- rune: %#v\n", rn, rn)
	fmt.Printf("rune: %c\n", rn) // f = format, %c =  print character, \n = new line

	// workshops.PrintfWs() // for workshop 1
	// workshops.PrintlnWs() // for workshop 2
}
