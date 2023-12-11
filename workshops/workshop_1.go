package workshops

import "fmt"

func PrintlnWs() {
	title := "Avengers: Endgame"
	year := 2019
	rating := 8.4
	genre := "Sci-Fi"
	isSuperhero := true

	fmt.Println("เรื่อง:", title)
	fmt.Println("ปี:", year)
	fmt.Println("เรตติ้ง:", rating)
	fmt.Println("ประเภท:", genre)
	fmt.Println("ซุปเปอร์ฮีโร่:", isSuperhero)

	// **raw string**

	/* 	movie := `เรื่อง: Avengers: Endgame
	ปี: 2019
	เรตติ้ง: 8.4
	ประเภท: Sci-Fi
	ซุปเปอร์ฮีโร่: true`
	fmt.Println(movie) */
}
