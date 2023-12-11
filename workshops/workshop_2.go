package workshops

import "fmt"

func PrintfWs() {
	title := "Avengers: Endgame"
	year := 2019
	rating := 8.4
	genre := "Sci-Fi"
	isSuperhero := true

	fmt.Printf("เรื่อง: %s\n", title)
	fmt.Printf("ปี: %d\n", year)
	fmt.Printf("เรตติ้ง: %.1f\n", rating)
	fmt.Printf("ประเภท: %s\n", genre)
	fmt.Printf("ซุปเปอร์ฮีโร่: %t\n", isSuperhero)
}
