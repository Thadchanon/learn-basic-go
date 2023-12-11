package workshops

import "fmt"

func IfElse() {
	ratings := 10.0

	if ratings < 5.0 && ratings >= 0.0 {
		fmt.Println("Disappointed 😞")
	} else if ratings >= 5.0 && ratings < 7.0 {
		fmt.Println("Normal 😐")
	} else if ratings >= 7.0 && ratings <= 10.0 {
		fmt.Println("Good 🥰")
	} else {
		fmt.Println("🤷🤷🤷🤷")
	}
}
