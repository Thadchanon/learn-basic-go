package workshops

import "fmt"

func emote(ratings float64) string {
	switch {
	case ratings < 5.0 && ratings >= 0.0:
		return "Disappointed 😞"
	case ratings >= 5.0 && ratings < 7.0:
		return "Normal 😐"
	case ratings >= 7.0 && ratings <= 10.0:
		return "Good 🥰"
	default:
		return "🤷🤷🤷🤷"
	}
}

func Function() {
	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0))
}