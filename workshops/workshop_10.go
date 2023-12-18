package workshops

import "fmt"

/* type movie struct {
	title       string
	year        int
	rating      float32
	genres      []string
	isSuperHero bool
} */ // ประกาศไว้ใน workshop_9 แล้ว

func (m movie) info() {
	fmt.Printf("%s (%d) - %.2f\n", m.title, m.year, m.rating)
	fmt.Println("Genres:")
	for _, g := range m.genres {
		fmt.Printf("\t%s\n", g)
	}
}

func Method() {
	ae := movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	ae.info()
}
